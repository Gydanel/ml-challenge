package server

import (
	"github.com/gin-gonic/gin"
	"ml-challenge/app/middlewares"
	"ml-challenge/app/rest"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(rest.HealthHandler)
	topSecret := new(rest.TopSecretHandler)
	split := new(rest.TopSecretSplitHandler)

	topSecret.Init()
	split.Init()

	router.Use(middlewares.AuthMiddleware())
	api := router.Group("api")
	{
		api.GET("/health", health.Status)
		api.POST("/topsecret", topSecret.DecodeMessage)
		splitGroup := api.Group("topsecret_split")
		{
			splitGroup.POST("/:satellite_name", split.AddSignal)
			splitGroup.GET("", split.DecodeMessage)
		}
	}

	return router

}
