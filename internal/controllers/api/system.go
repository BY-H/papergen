package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/models/system"
	"papergen/internal/models/user"
)

func Notifications(c *gin.Context) {
	var notifications []system.Notification
	global.DB.Find(&notifications)

	c.JSON(http.StatusOK, gin.H{
		"total":         len(notifications),
		"notifications": notifications,
	})
}

func AddNotification(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	err := checkRole(email)
	if err != nil {

	}
	msg := &message.AddNotificationMsg{}
	err = c.BindJSON(&msg)

	n := system.Notification{
		Type:    msg.Type,
		Title:   msg.Title,
		Content: msg.Content,
	}

	global.DB.Save(&n)

	c.JSON(http.StatusOK, gin.H{
		"msg": "add successfully",
	})
}

func checkRole(email string) error {
	var u user.User
	global.DB.Where("email = ?", email).First(&u)
	if u.Email != "" && u.Role == "admin" {
		return nil
	}
	return fmt.Errorf("this user is not admin user")
}
