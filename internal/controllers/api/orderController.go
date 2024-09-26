package api

import (
	"cyclopropane/internal/models/order"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "订单信息错误"})
		return
	}

	// TODO 将获取用户 email 的方法抽到一个统一的文件中
	creatorId, _ := c.Get("email")
	c.JSON(http.StatusOK, gin.H{
		"creator_id": creatorId,
	})
	return
}

func addOrder(o order.Order) {

}
