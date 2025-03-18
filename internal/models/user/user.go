package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;comment:'用户名'"`
	Password string `gorm:"column:password;comment:'密码'"`
	Email    string `gorm:"column:email;comment:'email/唯一标识符'"`
	Role     string `gorm:"column:role;comment:'角色，admin与user';default:user"`
}
