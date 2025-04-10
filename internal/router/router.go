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
		// 系统相关
		system := authRoutes.Group("/system")
		{
			notification := system.Group("/notifications")
			{
				notification.GET("/list", api.Notifications)
				notification.POST("/add", api.AddNotification)
			}
			feedback := system.Group("/feedbacks")
			{
				feedback.GET("/list", api.Feedbacks)
				feedback.POST("/add", api.AddFeedback)
			}
		}
		// 试卷相关
		papers := authRoutes.Group("/papers")
		{
			papers.GET("/list", api.Papers)
			papers.GET("/summary", api.PapersSummary)
			papers.POST("/auto_create", api.AutoCreatePaper)
			papers.POST("/manual_create", api.ManualCreatePaper)
			papers.DELETE("/delete", api.DeletePaper)
			papers.POST("/export", api.ExportPaper)
		}
		// 试题相关
		questions := authRoutes.Group("/questions")
		{
			questions.GET("/list", api.Questions)
			questions.GET("/summary", api.QuestionSummary)
			questions.POST("/add", api.AddQuestion)
			questions.PATCH("/edit", api.EditQuestion)
			questions.DELETE("/delete", api.DeleteQuestion)
			questions.GET("/tags", api.QuestionTags)
		}
	}

	return router
}
