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

	router.GET("/api/review/:destinationID", api.GetAllReviews)
	router.POST("", api.CreateReview)
	router.DELETE("/:id", api.DeleteReview)

	// reviewRouterWithAuth := router.Group("/api/review", AuthMiddleware())
	// {
	// 	// reviewRouterWithAuth.GET("", api.GetAllReviews)
	// 	reviewRouterWithAuth.POST("", api.CreateReview)
	// 	reviewRouterWithAuth.DELETE("/:id", api.DeleteReview)
	// }

	router.GET("/api/destination", api.GetAllDestination)
	router.GET("/api/destination/:id", api.GetDestination)
	router.POST("", api.CreateDestination)
	router.POST("/images/:id", api.UploadImage)
	router.PUT("", api.UpdateDestination) //edit destination
	router.PUT("/images/:id", api.UploadImage)
	router.DELETE("/:id", api.DeleteDestination)

	// destinationRouterWithAuth := router.Group("/api/destination", AdminMiddleware())
	// {
	// 	destinationRouterWithAuth.POST("", api.CreateDestination)
	// 	destinationRouterWithAuth.POST("/images/:id", api.UploadImage)
	// 	destinationRouterWithAuth.PUT("", api.UpdateDestination)
	// 	destinationRouterWithAuth.PUT("/images/:id", api.UploadImage)
	// 	destinationRouterWithAuth.DELETE("/:id", api.DeleteDestination)
	// }
	
	return api
}

func (api *API) Handler() *gin.Engine {
	return api.router
}

func (api *API) Start() {
	api.Handler().Run()
}