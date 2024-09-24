package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
}
