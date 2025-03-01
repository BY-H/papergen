package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
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
	e, _ := c.Get("email")
	email := e.(string)
	msg := &message.AddQuestionMsg{}
	err := c.BindJSON(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}
	if !msg.Check() {
		err = fmt.Errorf("question format error")
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
	}

	question := question.Question{
		Question:     msg.Question,
		QuestionType: (msg.QuestionType),
		Answer:       "",
		HardLevel:    0,
		Score:        0,
		Tag:          "",
		Creator:      email,
	}
}

func DeleteQuestion(c *gin.Context) {}

func EditQuestion(c *gin.Context) {}
