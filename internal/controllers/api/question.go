package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/models/question"
)

func Questions(c *gin.Context) {
	email, _ := c.Get("email")
	var msg message.RequestMsg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse)
		return
	}
	var questions []question.Question
	global.DB.Where("email = ?", email).Offset(msg.Page - 1).Limit(msg.PageSize).Find(&questions)

	c.JSON(http.StatusOK, gin.H{
		"page":      msg.Page,
		"page_size": msg.PageSize,
		"list":      questions,
	})

	return
}

func AddQuestion(c *gin.Context) {
	c.Get("email")

}

func DeleteQuestion(c *gin.Context) {}

func EditQuestion(c *gin.Context) {}
