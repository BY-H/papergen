package router

import (
	"cyclopropane/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(
		middleware.AuthService(),
		middleware.CorsService(),
	)
	return router
}
