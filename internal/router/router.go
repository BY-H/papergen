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
		// 用户相关
		users := authRoutes.Group("/users")
		{
			users.GET("/summary", api.UsersSummary)
		}
		// 试卷相关
		papers := authRoutes.Group("/papers")
		{
			papers.GET("/", api.Papers)
			papers.GET("/summary", api.PapersSummary)
			papers.POST("/add", api.CreatePaper)
			papers.PATCH("/edit", api.EditPaper)
			papers.DELETE("/delete", api.RemovePaper)
		}
		// 试题相关
		questions := authRoutes.Group("/questions")
		{
			questions.GET("/", api.Questions)
			questions.GET("/summary", api.QuestionSummary)
			questions.POST("/add", api.AddQuestion)
			questions.PATCH("/edit", api.EditQuestion)
			questions.DELETE("/delete", api.DeleteQuestion)
		}
	}

	return router
}
