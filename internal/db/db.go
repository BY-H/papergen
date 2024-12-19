package db

import (
	"cyclopropane/internal/models/order"
	"cyclopropane/internal/models/user"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(host string, username string, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/cyclopropane?charset=utf8&parseTime=True", username, password, host)
	return initDB(dsn)
}

func initDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 设置日志级别为 Info
	})
	if err != nil {
		fmt.Printf("%t\n", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&user.User{},
		&order.Order{},
	)
	if err != nil {
		return nil, err
	}
	return db, err
}
