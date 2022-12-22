package controllers

import (
	"reflect"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/go-playground/validator/v10"

	"github.com/althafariq/pemanduwisata-be/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type API struct {
	userModels					models.UserModels
	destinationModels		models.DestinationModels
	reviewModels				models.ReviewModels
	telpDaruratModels 	models.TelpDaruratModels
	router 							*gin.Engine
}

func NewApi(
	userModels models.UserModels,
) API {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	router.Use(cors.New(config))
	router.RedirectTrailingSlash = false

	api := API {
		router: router,
		userModels: userModels,
	}

	// Untuk validasi request dengan mengembalikan nama dari tag json jika ada
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
	router.Static("/media", "./media")

	router.POST("/api/login", api.login)
	router.POST("/api/register", api.register)

	profileRouter := router.Group("/api/profile", AuthMiddleware())
	{
		profileRouter.GET("", api.getProfile)
		profileRouter.PATCH("", api.updateProfile)
		profileRouter.PUT("/avatar", api.changeAvatar)
	}

	

	
	
	return api
}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}