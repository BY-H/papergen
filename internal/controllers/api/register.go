package api

import (
	"cyclopropane/internal/global"
	"cyclopropane/internal/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.DB.First(&models.User{Email: user.Email}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "user existed"})
		return
	}

	addUser(user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "register successfully",
	})
}

// addUser　创建用户
func addUser(user models.User) {
	user.Password, _ = encryptPassword(user.Password)

	result := global.DB.Create(&user)
	if result.Error != nil {
		global.Logger.Error(result.Error.Error())
	}
}

// encryptPassword 生成加密结果
func encryptPassword(pwd string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(encrypt), err
}
