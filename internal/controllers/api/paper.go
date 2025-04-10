package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/models/paper"
	"papergen/internal/models/question"
	"papergen/pkg/utils"
	"strings"
)

// Papers 返回用户创建过的试卷
func Papers(c *gin.Context) {
	// TODO 添加对应条件的查找功能
	email, _ := c.Get("email")
	var msg message.RequestMsg
	err := c.BindQuery(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}
	var papers []paper.Paper
	global.DB.Where("creator = ?", email).Offset(msg.Page - 1).Limit(msg.PageSize).Find(&papers)

	var count int64
	global.DB.Model(&paper.Paper{}).Where("creator = ?", email).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"page":      msg.Page,
		"page_size": msg.PageSize,
		"total":     count,
		"list":      papers,
	})

	return
}

// PapersSummary 试卷简要汇总信息
func PapersSummary(c *gin.Context) {
	var count int64
	global.DB.Model(&paper.Paper{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"total": count,
	})
}

// AutoCreatePaper 自动创建试卷
func AutoCreatePaper(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := &message.AutoCreatePaperMsg{}
	err := c.BindJSON(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	// 1. 验证题目数量是否合理
	total := msg.SingleChoiceCount + msg.MultiChoiceCount + msg.TrueFalseCount +
		msg.FillBlankCount + msg.ShortAnswerCount

	if total <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "题目总数不能为0"})
		return
	}

	// 2. 从题库中随机抽取题目
	questionIDs, err := selectRandomQuestions(email, msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	if len(questionIDs) != total {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("题库中符合条件的题目不足，需要%d题但只找到%d题", total, len(questionIDs)),
		})
		return
	}

	// 3. 创建试卷
	questionsJSON, err := json.Marshal(questionIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	newPaper := paper.Paper{
		Creator:     email,
		Title:       msg.Title,
		Description: msg.Description,
		Questions:   datatypes.JSON(questionsJSON),
	}

	if err := global.DB.Create(&newPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "试卷自动创建成功",
		"data":    newPaper,
	})
}

// 从题库中随机抽取题目
func selectRandomQuestions(email string, msg *message.AutoCreatePaperMsg) ([]int, error) {
	var questionIDs []int

	// 定义题型映射
	type questionType struct {
		count int
		qType string
	}

	types := []questionType{
		{msg.SingleChoiceCount, "single_choice"},
		{msg.MultiChoiceCount, "multi_choice"},
		{msg.TrueFalseCount, "true_false"},
		{msg.FillBlankCount, "fill_blank"},
		{msg.ShortAnswerCount, "short_answer"},
	}

	for _, t := range types {
		if t.count <= 0 {
			continue
		}

		var ids []int
		err := global.DB.Model(&question.Question{}).
			Where("creator = ? or creator = 'system'", email).
			Where("tag = ?", msg.Tag).
			Where("question_type = ?", t.qType).
			Order("RAND()").
			Limit(t.count).
			Pluck("id", &ids).
			Error

		if err != nil {
			return nil, err
		}

		if len(ids) < t.count {
			return nil, fmt.Errorf("题型%s的题目不足，需要%d题但只找到%d题",
				t.qType, t.count, len(ids))
		}

		questionIDs = append(questionIDs, ids...)
	}

	return questionIDs, nil
}

// ManualCreatePaper 手动创造试卷
func ManualCreatePaper(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := &message.ManualCreatePaperMsg{}
	err := c.BindJSON(msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	// 验证题目是否存在且属于当前用户
	var count int64
	global.DB.Model(&question.Question{}).
		Where("id IN ?", msg.QuestionIds).
		Where("creator = ? or creator = 'system'", email).
		Count(&count)

	if count != int64(len(msg.QuestionIds)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "部分题目不存在或不属于当前用户"})
		return
	}

	// 将 []int 转换为 JSON
	questionsJSON, err := json.Marshal(msg.QuestionIds)
	if err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	// 从数据库中找出对应的试题
	newPaper := paper.Paper{
		Creator:     email,
		Title:       msg.Title,
		Description: msg.Description,
		Questions:   datatypes.JSON(questionsJSON), // 转换为 JSON 格式
	}

	// 保存到数据库
	if err := global.DB.Create(&newPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "试卷创建成功",
		"paper":   newPaper,
	})
}

// DeletePaper 删除试卷
func DeletePaper(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	msg := message.DeletePaperMsg{}
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
	global.DB.Where("creator = ?", email).Delete(&paper.Paper{}, ids)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"msg":    "delete paper successfully",
	})
}

// ExportPaper 导出试卷
func ExportPaper(c *gin.Context) {
	msg := &message.ExportPaperMsg{}
	if err := c.BindJSON(msg); err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	var exportPaper paper.Paper
	if err := global.DB.Where("id = ?", msg.ID).First(&exportPaper).Error; err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	var questionsIds []int
	if err := json.Unmarshal(exportPaper.Questions, &questionsIds); err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	var questions []question.Question
	if err := global.DB.Where("id IN ?", questionsIds).Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, message.ErrorResponse(err))
		return
	}

	if err := paper.GenerateDocxPaper(exportPaper.Title, questions); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename=\""+exportPaper.Title+".docx\"")
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.wordprocessingml.document")
	c.File("tmp/" + exportPaper.Title + ".docx")
}
