package api

import (
	"fmt"
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

	q := question.Question{
		Question:     msg.Question,
		QuestionType: question.Type(msg.QuestionType),
		Answer:       msg.Answer,
		HardLevel:    msg.HardLevel,
		Score:        msg.Score,
		Tag:          msg.Tag,
		Creator:      email,
	}

	global.DB.Create(&q)
	c.JSON(http.StatusOK, gin.H{
		"msg": "add question successfully",
	})
}

func DeleteQuestion(c *gin.Context) {
	e, _ := c.Get("email")
}

func EditQuestion(c *gin.Context) {}
