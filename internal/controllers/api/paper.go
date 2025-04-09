package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/models/paper"
)

// Papers 返回用户创建过的试卷
func Papers(c *gin.Context) {
	// TODO 添加对应条件的查找功能
	email, _ := c.Get("email")
	var msg message.RequestMsg
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}
	var papers []paper.Paper
	global.DB.Where("creator = ?", email).Offset(msg.Page - 1).Limit(msg.PageSize).Find(&papers)

	var count int64
	global.DB.Where("creator = ?", email).Count(&count)
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
	c.Get("email")
}

// ManualCreatePaper 手动创造试卷
func ManualCreatePaper(c *gin.Context) {
	c.Get("email")
}

// RemovePaper 移除试卷
func RemovePaper(c *gin.Context) {

}
