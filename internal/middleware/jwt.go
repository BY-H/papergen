package middleware

import (
	"cyclopropane/internal/global"
	"github.com/dgrijalva/jwt-go"
	"log"
)

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func MakeClaimsToken(claims JWTClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.JWTKey)
	return tokenString, err
}

// ParseClaimsToken Token解签
func ParseClaimsToken(tokenStr string) (JWTClaim, error) {
	ParsedClaims := JWTClaim{}
	tokenClaims, err := jwt.ParseWithClaims(tokenStr, &ParsedClaims, func(token *jwt.Token) (interface{}, error) {
		return global.JWTKey, nil
	})
	if err != nil && !tokenClaims.Valid {
		// TODO 更换成全局日志
		log.Printf("Invalid Token:%v\n", err)
	}
	return ParsedClaims, err
}
