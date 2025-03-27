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
		"total": len(notifications),
		"list":  notifications,
	})
}

func AddNotification(c *gin.Context) {
	e, _ := c.Get("email")
	email := e.(string)
	err := checkRole(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
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

func Feedbacks(c *gin.Context) {
	var feedbacks []system.Feedback
	global.DB.Find(&feedbacks)

	c.JSON(http.StatusOK, gin.H{
		"total": len(feedbacks),
		"list":  feedbacks,
	})
}

func AddFeedback(c *gin.Context) {
	msg := message.AddFeedbackMsg{}
	err := c.BindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	f := system.Feedback{
		Content: msg.Content,
	}

	global.DB.Save(&f)

	c.JSON(http.StatusOK, gin.H{
		"msg": "add successfully",
	})
}
