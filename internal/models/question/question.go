package question

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Type string

// 定义允许的 QuestionType 值
const (
	TypeSingleChoice   Type = "single_choice"
	TypeMultipleChoice Type = "multiple_choice"
	TypeTrueFalse      Type = "true_false"
	TypeShortAnswer    Type = "short_answer"
)

// 验证 questionType 是否合法
func (qt Type) isValid() error {
	switch qt {
	case TypeSingleChoice, TypeMultipleChoice, TypeTrueFalse, TypeShortAnswer:
		return nil
	default:
		return errors.New("invalid question type")
	}
}

type Question struct {
	gorm.Model
	Question     string `gorm:"column:question;comment:'题目正文'" json:"question"`
	QuestionType Type   `gorm:"column:question_type;comment:'题目类型'" json:"question_type"`
	Answer       string `gorm:"column:answer;comment:'答案'" json:"answer"`
	HardLevel    int    `gorm:"column:hard_level;comment:'难度'" json:"hard_level"`
	Score        int    `gorm:"column:score;comment:'分值'" json:"score"`
	Tag          string `gorm:"column:tag;comment:'题目标签'" json:"tag"`
	Creator      string `gorm:"column:creator;comment:'题目创建人'" json:"creator"`
}

// BeforeCreate 创建 Question 时验证 QuestionType
func (q *Question) BeforeCreate(tx *gorm.DB) error {
	if err := q.QuestionType.isValid(); err != nil {
		return fmt.Errorf("invalid question type: %v", err)
	}
	return nil
}

// BeforeUpdate 更新 Question 时验证 QuestionType
func (q *Question) BeforeUpdate(tx *gorm.DB) error {
	if err := q.QuestionType.isValid(); err != nil {
		return fmt.Errorf("invalid question type: %v", err)
	}
	return nil
}
