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

	// TODO 将获取用户 email 的方法抽到一个统一的文件中
	id, _ := c.Get("email")
	creatorId, _ := id.(string)
	o.CreatorID = creatorId

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

	var query = global.DB

	if msg.DateStart != "" && msg.DateEnd != "" {
		query = query.Where("report_date BETWEEN ? AND ?", msg.DateStart, msg.DateEnd)
	}

	var orders []order.Order
	result := query.Find(&orders)
	if result.Error != nil {
		global.Logger.Error(result.Error.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total": result.RowsAffected,
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
