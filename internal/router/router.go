package router

import (
	"cyclopropane/internal/controllers/api"
	"cyclopropane/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.GET("/ping_without_login", api.Ping)

	authRoutes := router.Group("/")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.GET("/ping", api.Ping)
	}

	return router
}
