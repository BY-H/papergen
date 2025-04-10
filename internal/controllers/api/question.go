package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/models/question"
	"papergen/pkg/utils"
	"strings"
)

func Questions(c *gin.Context) {
	email, _ := c.Get("email")
	var msg message.GetQuestionMsg
	err := c.BindQuery(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}
	var questions []question.Question
	fmt.Printf("%v\n", msg.QuestionIds)
	if msg.QuestionIds != nil && len(msg.QuestionIds) != 0 {
		global.DB.Where("creator = ? or creator = 'system'", email).Where(&msg.QuestionIds).Find(&questions)
	} else {
		global.DB.Where("creator = ? or creator = 'system'", email).Offset(msg.Page - 1).Limit(msg.PageSize).Find(&questions)
	}
	var count int64
	global.DB.Model(&question.Question{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"page":      msg.Page,
		"page_size": msg.PageSize,
		"list":      questions,
		"total":     count,
	})

	return
}

func AddQuestion(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := &message.AddQuestionMsg{}
	err := c.BindJSON(&msg)
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
		QuestionType: msg.QuestionType,
		Options:      msg.Options,
		Answer:       msg.Answer,
		HardLevel:    msg.HardLevel,
		Score:        msg.Score,
		Tag:          msg.Tag,
		Creator:      email,
	}

	global.DB.Create(&q)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "add question successfully",
	})
}

func DeleteQuestion(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := message.DeleteQuestionMsg{}
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	temp := strings.Split(msg.IDs, ",")
	ids, err := utils.StringArrToIntArr(temp)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	global.DB.Where("creator = ?", email).Delete(&question.Question{}, ids)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "delete question successfully",
	})
}

func EditQuestion(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := message.EditQuestionMsg{}
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	// 先检查问题是否存在且属于该用户
	var existingQuestion question.Question
	result := global.DB.Where("(creator = ? or creator = 'system') AND id = ?", email, msg.ID).First(&existingQuestion)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "question not found or not owned by user"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		}
		return
	}

	// 更新字段
	updates := question.Question{
		Question:     msg.Question,
		QuestionType: msg.QuestionType,
		Options:      msg.Options,
		Answer:       msg.Answer,
		HardLevel:    msg.HardLevel,
		Score:        msg.Score,
		Tag:          msg.Tag,
		// 注意不要更新 Creator，因为创建者不应该改变
	}

	// 执行更新
	result = global.DB.Model(&question.Question{}).
		Where("id = ? AND creator = ?", msg.ID, email).
		Updates(updates)

	if result.Error != nil {
		fmt.Println(result.Error)
		fmt.Println(updates.QuestionType)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "update failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "update question successfully",
	})
}

func QuestionSummary(c *gin.Context) {
	var count int64
	global.DB.Model(&question.Question{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"total": count,
	})
}

func QuestionTags(c *gin.Context) {
	var result []string
	global.DB.Model(&question.Question{}).Distinct("tag").Pluck("tag", &result)
	c.JSON(http.StatusOK, gin.H{
		"total": len(result),
		"tag":   result,
	})
}
