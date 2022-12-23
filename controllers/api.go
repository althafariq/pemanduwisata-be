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
	destinationModels		models.DestinationModels,
	reviewModels				models.ReviewModels,
	telpDaruratModels 	models.TelpDaruratModels,
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
		destinationModels: destinationModels,
		reviewModels: reviewModels,
		telpDaruratModels: telpDaruratModels,
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

	router.GET("/api/telp-darurat", api.GetAllTelpDarurat)

	profileRouter := router.Group("/api/profile", AuthMiddleware())
	{
		profileRouter.GET("", api.getProfile)
		profileRouter.PATCH("", api.updateProfile)
		profileRouter.PUT("/avatar", api.changeAvatar)
	}

	router.GET("/api/review", api.GetAllReviews)
	reviewRouterWithAuth := router.Group("/api/review", AuthMiddleware())
	{
		// reviewRouterWithAuth.GET("", api.GetAllReviews)
		reviewRouterWithAuth.POST("", api.CreateReview)
		reviewRouterWithAuth.DELETE("/:id", api.DeleteReview)
	}

	router.GET("/api/destination", api.GetAllDestination)
	
	return api
}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}