package question

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// 定义允许的 QuestionType 值
const (
	TypeDefault        string = ""
	TypeSingleChoice   string = "single_choice"
	TypeMultipleChoice string = "multiple_choice"
	TypeTrueFalse      string = "true_false"
	TypeShortAnswer    string = "short_answer"
	TypeFillBlank      string = "fill_blank"
)

// 验证 questionType 是否合法
func isValid(qt string) error {
	switch qt {
	case TypeSingleChoice, TypeMultipleChoice, TypeTrueFalse, TypeShortAnswer, TypeFillBlank, TypeDefault:
		return nil
	default:
		return errors.New("invalid question type" + qt)
	}
}

type Question struct {
	gorm.Model
	Question     string `gorm:"column:question;comment:'题目正文'" json:"question"`
	QuestionType string `gorm:"column:question_type;comment:'题目类型'" json:"question_type"`
	Options      string `gorm:"column:options;comment:'题目选项'" json:"options"`
	Answer       string `gorm:"column:answer;comment:'答案'" json:"answer"`
	HardLevel    int    `gorm:"column:hard_level;comment:'难度'" json:"hard_level"`
	Score        int    `gorm:"column:score;comment:'分值'" json:"score"`
	Tag          string `gorm:"column:tag;comment:'题目标签'" json:"tag"`
	Creator      string `gorm:"column:creator;comment:'题目创建人'" json:"creator"`
}

// BeforeCreate 创建 Question 时验证 QuestionType
func (q *Question) BeforeCreate(tx *gorm.DB) error {
	if err := isValid(q.QuestionType); err != nil {
		return fmt.Errorf("invalid question type: %v", err)
	}
	return nil
}

// BeforeUpdate 更新 Question 时验证 QuestionType
func (q *Question) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println(q.QuestionType)
	if err := isValid(q.QuestionType); err != nil {
		return fmt.Errorf("invalid question type: %v", err)
	}
	return nil
}
