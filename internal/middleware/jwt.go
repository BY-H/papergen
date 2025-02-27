package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"papergen/internal/global"
	"strings"
	"time"
)

const Issuer = "abing"
const TokenExpireDuration = time.Hour * 24

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func MakeClaimsToken(claims JWTClaim) (string, error) {
	claims.Issuer = Issuer
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(TokenExpireDuration).Unix()
	claims.Subject = "authorization"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(global.JWTKey))
	return tokenString, err
}

// ParseClaimsToken Token解签
func ParseClaimsToken(tokenStr string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(global.JWTKey), nil
	})

	if err != nil {
		fmt.Printf("%+v\n", token.Claims.(*JWTClaim))
		fmt.Println(err.Error())
		return nil, err
	}

	// 校验token
	if claims, ok := token.Claims.(*JWTClaim); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.Next()
			return
		}
		authHeader := c.Request.Header.Get("authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请求头中auth为空",
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}

		claims, err := ParseClaimsToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "无效的token",
			})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
		// 后续的处理函数可以通过c.Get("email")来获取当前请求的用户邮箱信息
	}
}
