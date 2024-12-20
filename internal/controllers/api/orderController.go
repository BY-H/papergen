package api

import (
	"cyclopropane/internal/controllers/message"
	"cyclopropane/internal/global"
	"cyclopropane/internal/models/order"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AddOrder 创建订单
func AddOrder(c *gin.Context) {
	var o order.Order

	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检验订单
	if !o.CheckOrder() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "订单信息错误"})
		return
	}

	o.Status = order.STATUS_RECORD
	o.ReportDate = time.Now().Format("2006-01-02")

	err := addOrder(o)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

	return
}

func addOrder(o order.Order) error {
	result := global.DB.Create(&o)
	if result.Error != nil {
		global.Logger.Error(result.Error.Error())
		return result.Error
	}
	return nil
}

// GetOrder 获取订单信息
func GetOrder(c *gin.Context) {
	// TODO 添加筛选条件
	msg := message.OrderMsg{}

	if err := c.ShouldBindQuery(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var builder = global.DB.Model(&order.Order{})

	if msg.DateStart != "" && msg.DateEnd != "" {
		builder = builder.Where("report_date BETWEEN ? AND ?", msg.DateStart, msg.DateEnd)
	}

	// 分页
	page := msg.Page
	if page == 0 {
		page = 1
	}
	pageSize := msg.PageSize
	if pageSize == 0 {
		pageSize = 20
	}

	var total int64
	if err := builder.Count(&total).Error; err != nil {
		global.Logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count records"})
		return
	}

	// 查询分页数据
	var orders []order.Order
	query := builder.Order("id desc").Offset((page - 1) * pageSize).Limit(pageSize)
	if err := query.Find(&orders).Error; err != nil {
		global.Logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total": total,
		"list":  orders,
	})
	return
}

type updateOrderStatusMessage struct {
	Ids        []int  `json:"ids"`
	StatusCode string `json:"status"`
}

func StartOrder(c *gin.Context) {
	msg := updateOrderStatusMessage{}
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := updateOrderStatus(msg); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ids": msg.Ids,
	})
}

func updateOrderStatus(msg updateOrderStatusMessage) error {
	result := global.DB.Table("orders").
		Where("id IN ?", msg.Ids).
		Updates(order.Order{Status: msg.StatusCode})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
