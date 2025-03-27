package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"papergen/internal/models/paper"
	"papergen/internal/models/question"
	"papergen/internal/models/system"
	"papergen/internal/models/user"
)

func InitDB(host string, username string, password string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/papergen?charset=utf8&parseTime=True", username, password, host)
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
	// 自动迁移
	err = db.AutoMigrate(
		&user.User{},
		&question.Question{},
		&paper.Paper{},
		&system.Notification{},
		&system.Feedback{},
	)
	fmt.Printf("test db init\n")
	if err != nil {
		return nil, err
	}
	return db, err
}
