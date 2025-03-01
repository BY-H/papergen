package message

import "github.com/gin-gonic/gin"

type RequestMsg struct {
	DateStart string `json:"date_start" form:"date_start"`
	DateEnd   string `json:"date_end" form:"date_end"`
	Page      int    `json:"page" form:"page" default:"1"`
	PageSize  int    `json:"page_size" form:"page_size" default:"20"`
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"status": "error",
		"error":  err,
	}
}

type AddQuestionMsg struct {
	Question     string `json:"question" form:"question"`
	QuestionType string `json:"question_type" form:"question_type"`
	Answer       string `json:"answer" form:"answer"`
	HardLevel    int    `json:"hard_level" form:"hard_level"`
	Score        int    `json:"score" form:"score"`
	Tag          string `json:"tag" form:"tag"`
}

func (m *AddQuestionMsg) Check() bool {
	if m.Question == "" {
		return false
	}
	if m.QuestionType == "" {
		return false
	}
	if m.Answer == "" {
		return false
	}
	return true
}
