package router

import (
	"github.com/gin-gonic/gin"
	"papergen/internal/controllers/api"
	"papergen/internal/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.GET("/ping_without_login", api.Ping)

	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		// 测试接口
		authRoutes.GET("/ping", api.Ping)
		authRoutes.GET("/papers", api.Papers)
	}

	return router
}
