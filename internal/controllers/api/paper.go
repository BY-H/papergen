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

	c.JSON(http.StatusOK, gin.H{
		"page":      msg.Page,
		"page_size": msg.PageSize,
		"list":      papers,
	})

	return
}

// CreatePaper 创建试卷
func CreatePaper(c *gin.Context) {
	c.Get("email")
}

func EditPaper(c *gin.Context) {

}

// RemovePaper 移除试卷
func RemovePaper(c *gin.Context) {

}
