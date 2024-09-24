package api

import (
	"cyclopropane/internal/global"
	"cyclopropane/internal/middleware"
	"cyclopropane/internal/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检验用户
	if !checkUser(user) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 返回对应 jwt 密钥
	token, err := middleware.MakeClaimsToken(middleware.JWTClaim{Email: user.Email})
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	return
}

// checkUser 检验用户
func checkUser(user models.User) bool {
	var dbUser models.User
	result := global.DB.Where("email = ?", user.Email).First(&dbUser)

	// 检查是否找到用户
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	// 检查是否有其他错误
	if result.Error != nil {
		fmt.Println(result.Error.Error())
		return false
	}

	// 检查密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
