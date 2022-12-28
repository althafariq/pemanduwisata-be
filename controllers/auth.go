package controllers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/althafariq/pemanduwisata-be/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type LoginReqBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterReqBody struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role string `json:"role" validate:"oneof=admin user"`
	ConfirmPassword string `json:"confirm_password" binding:"required" validate:"eqfield=Password"`
}

type ProfilePicReqBody struct {
	ProfilePic *multipart.FileHeader `form:"profile_pic"`
}

type LoginSuccessResponse struct {
	Token string `json:"token"`
	Fullname string `json:"fullname"`
	Message string `json:"message"`
}

var jwtKey = []byte("key")

type Claims struct {
	Id    int
	Email string
	Role  string
	jwt.StandardClaims
}

func (api *API) register(c *gin.Context) {
	var input RegisterReqBody
	err := c.BindJSON(&input)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
			return
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, responseCode, err := api.userModels.Register(input.Firstname, input.Lastname, input.Email, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(responseCode, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(&userId, &input.Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{Token: tokenString,
		Fullname: fmt.Sprintf("%s %s", input.Firstname, input.Lastname),
		Message: fmt.Sprintf("You're now login as %s", input.Role),
	})
}

func (api *API) login(c *gin.Context) {
	var loginReq LoginReqBody
	err := c.BindJSON(&loginReq)
	var ve validator.ValidationErrors

	if err != nil {
		if errors.As(err, &ve) {
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"errors": helper.GetErrorMessage(ve)},
			)
		} else {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	userId, err := api.userModels.Login(loginReq.Email, loginReq.Password)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	firstname, lastname, role, err := api.userModels.GetUserRole(*userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := api.generateJWT(userId, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, LoginSuccessResponse{
		Token: tokenString,
		Fullname: fmt.Sprintf("%s %s", *firstname, *lastname),
		Message: fmt.Sprintf("You're now login as %s", *role),
	})
}

func (api *API) changeAvatar(c *gin.Context) {
	var input ProfilePicReqBody
	maxFileSize := int64(1024 * 1024 * 2)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.ProfilePic.Size > maxFileSize {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "file size too large"})
		return
	}

	if !strings.Contains(input.ProfilePic.Header.Get("Content-Type"), "image") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "please upload an image"})
		return
	}

	userId, err := api.getUserIdFromToken(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userData, err := api.userModels.GetUserData(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	oldFileName := userData.Profile_pic

	folderPath := "media/avatar"
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	splitFilename := strings.Split(input.ProfilePic.Filename, ".")
	fileName := fmt.Sprintf("%s %s_%d.%s", userData.Firstname, userData.Lastname, time.Now().Unix(), splitFilename[len(splitFilename)-1])
	filePath := filepath.Join(folderPath, fileName)
	err = c.SaveUploadedFile(input.ProfilePic, filePath)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if oldFileName != nil {
		os.Remove(*oldFileName)
	}

	err = api.userModels.UpdateAvatar(userId, filePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	scheme := "http"
	if c.Request.TLS != nil {
		scheme = "https"
	}
	imgUrl := fmt.Sprintf("%s://%s/%s/%s", scheme, c.Request.Host, folderPath, url.PathEscape(fileName))
	c.JSON(http.StatusOK, gin.H{"message": "success",
		"data": struct {
			Avatar string `json:"avatar"`
		}{
			Avatar: imgUrl,
		},
	})

}

func (api *API) getUserIdFromToken(c *gin.Context) (int, error) {
	tokenString := c.GetHeader("Authorization")[(len("Bearer ")):]
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return -1, err
	}

	if token.Valid {
		claim := token.Claims.(*Claims)
		return claim.Id, nil

	} else {
		return -1, errors.New("invalid token")
	}
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	claim := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, err
}

func (api API) generateJWT(userId *int, role *string) (string, error) {
	expTime := time.Now().Add(180 * time.Minute)

	claims := &Claims{
		Id:   *userId,
		Role: *role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
