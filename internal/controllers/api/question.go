package api

import "github.com/gin-gonic/gin"

func Questions(c *gin.Context) {
	c.Get("email")

	return
}

func AddQuestion(c *gin.Context) {
	c.Get("email")
}

func DeleteQuestion(c *gin.Context) {}

func EditQuestion(c *gin.Context) {}
