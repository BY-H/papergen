package router

import (
	"github.com/gin-gonic/gin"
	"papergen/internal/controllers/api"
	"papergen/internal/middleware"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CORS())
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)
	router.GET("/ping_without_login", api.Ping)

	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.GET("/ping", api.Ping)
		// 试卷相关
		papers := authRoutes.Group("/papers")
		{
			papers.GET("/", api.Papers)
			papers.POST("/add", api.CreatePaper)
		}
		// 试题相关
		authRoutes.GET("/questions", api.Questions)
	}

	return router
}
