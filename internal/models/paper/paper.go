package paper

import "gorm.io/gorm"

type Paper struct {
	gorm.Model
	Title     string  `gorm:"column:title;comment:'试卷标题'" json:"title"`
	Questions []int32 `gorm:"column:questions;type:json;comment:'试卷包含题目ID'" json:"questions"`
}
