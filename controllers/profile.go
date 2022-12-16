package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Email     string `json:"email" binding:"required"`
}

type UserResponse struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
	Role       string `json:"role"`
}

type Response struct {
	Message string `json:"message"`
}

func (api *API) getProfile(ctx *gin.Context) {
	userID, err := api.getUserIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, Response{"Unauthorized"})
		return
	}

	user, err := api.userModels.GetUserData(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Message: "Internal Server Error"})
		return
	}

	var (
		userProfilePic string
	)

	if user.Profile_pic != nil {
		userProfilePic = *user.Profile_pic
	} 

	ctx.JSON(http.StatusOK, UserResponse{
		ID:         user.ID,
		Firstname:  user.Firstname,
		Lastname:   user.Lastname,
		Email:      user.Email,
		ProfilePic: userProfilePic,
		Role:       user.Role,
	})
}

func (api *API) updateProfile(ctx *gin.Context) {
	var request UpdateUserRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, Response{Message: "Invalid Request"})
		return
	}

	userID, err := api.getUserIdFromToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, Response{"Unauthorized"})
		return
	}

	if err := api.userModels.UpdateUserData(userID, request.Firstname, request.Lastname, request.Email); err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, Response{Message: "Successfully Updated"})
}