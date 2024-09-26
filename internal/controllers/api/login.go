package api

import (
	"cyclopropane/internal/global"
	"cyclopropane/internal/middleware"
	"cyclopropane/internal/models/user"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func Login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检验用户
	if !checkUser(u) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		return
	}

	// 返回对应 jwt 密钥
	token, err := middleware.MakeClaimsToken(middleware.JWTClaim{Email: u.Email})
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	return
}

// checkUser 检验用户
func checkUser(u user.User) bool {
	var dbUser user.User
	result := global.DB.Where("email = ?", u.Email).First(&dbUser)

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
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(u.Password))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}
