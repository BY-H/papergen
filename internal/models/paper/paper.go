package paper

import "gorm.io/gorm"

type Paper struct {
	gorm.Model
	Title     string `gorm:"column:title;comment:'试卷标题'" json:"title"`
	Questions []byte `gorm:"column:questions;type:json;comment:'试卷包含题目ID'" json:"questions"`
	Creator   string `gorm:"column:creator;comment:'试卷创建人'" json:"creator"`
}
