package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"papergen/internal/controllers/message"
	"papergen/internal/global"
	"papergen/internal/middleware"
	"papergen/internal/models/user"
)

func Login(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, message.ErrorResponse(err))
		return
	}

	// 检验用户
	if !checkUser(u) {
		c.JSON(http.StatusForbidden, gin.H{"message": "用户名或密码错误"})
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

func Register(c *gin.Context) {
	var u user.User

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if u.Email == "" || u.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "邮箱或密码不合法"})
		return
	}

	var count int64
	global.DB.Model(&user.User{}).Where("email = ?", u.Email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "账号已存在"})
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

// UsersSummary 用户简要汇总信息
func UsersSummary(c *gin.Context) {
	var count int64
	global.DB.Model(&user.User{}).Count(&count)
	c.JSON(http.StatusOK, gin.H{
		"total": count,
	})
}
