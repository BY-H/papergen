package paper

import "gorm.io/gorm"
import "gorm.io/datatypes"

type Paper struct {
	gorm.Model
	Title       string         `gorm:"column:title;comment:'试卷标题'" json:"title"`
	Description string         `gorm:"column:description;comment:'试卷描述'" json:"description"`
	Questions   datatypes.JSON `gorm:"column:questions;type:json;comment:'试卷包含题目ID'" json:"questions"`
	Creator     string         `gorm:"column:creator;comment:'试卷创建人'" json:"creator"`
}
