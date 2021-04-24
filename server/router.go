package server

import (
	"github.com/gin-gonic/gin"
	"ml-challenge/controllers"
	"ml-challenge/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.Use(middlewares.AuthMiddleware())
	router.GET("/health", health.Status)
	return router

}
