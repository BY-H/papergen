package router

import (
	"cyclopropane/internal/controllers/api"
	"cyclopropane/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(
		middleware.AuthService(),
		middleware.CorsService(),
	)
	router.GET("/ping", api.Ping)
	return router
}
