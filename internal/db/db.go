package db

import (
	"cyclopropane/internal/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(host string, username string, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/cyclopropane?charset=utf8&parseTime=True", username, password, host)
	return initDB(dsn)
}

func initDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Printf("%t\n", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}
	return db, err
}
