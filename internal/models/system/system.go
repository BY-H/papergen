package system

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	Type    string `gorm:"column:type;comment:'类型：通知(Notify)，警告(Alert)'" json:"type"`
	Title   string `gorm:"column:title;comment:'标题'" json:"title"`
	Content string `gorm:"column:content;comment:'正文'" json:"content"`
}

type Feedback struct {
	gorm.Model
	Content   string    `gorm:"column:content;comment:'反馈内容'" json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
