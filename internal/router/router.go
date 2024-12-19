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

	authRoutes := router.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		// 测试接口
		authRoutes.GET("/ping", api.Ping)

		// 订单相关接口
		orders := authRoutes.Group("/order")
		{
			orders.GET("/list", api.GetOrder)
			orders.POST("/add", api.AddOrder)
			orders.PATCH("/update", api.StartOrder)
		}
	}

	return router
}
