package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"papergen/internal/global"
	"papergen/internal/models/user"
)

func Register(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.DB.First(&user.User{Email: u.Email}).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "u existed"})
		return
	}

	addUser(u)
	c.JSON(http.StatusOK, gin.H{
		"msg": "register successfully",
	})
}

// addUser　创建用户
func addUser(u user.User) {
	u.Password, _ = encryptPassword(u.Password)

	result := global.DB.Create(&u)
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
