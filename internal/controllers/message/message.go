package message

import "github.com/gin-gonic/gin"

type RequestMsg struct {
	DateStart string `json:"date_start" form:"date_start"`
	DateEnd   string `json:"date_end" form:"date_end"`
	Page      int    `json:"page" form:"page" default:"1"`
	PageSize  int    `json:"page_size" form:"page_size" default:"20"`
}

type GetQuestionMsg struct {
	RequestMsg
	QuestionIds []int `json:"question_ids" form:"question_ids[]"`
}

func ErrorResponse(err error) gin.H {
	return gin.H{
		"status": "error",
		"error":  err.Error(),
	}
}

type AddQuestionMsg struct {
	Question     string `json:"question" form:"question"`
	QuestionType string `json:"question_type" form:"question_type"`
	Options      string `json:"options" form:"options"`
	Answer       string `json:"answer" form:"answer"`
	HardLevel    int    `json:"hard_level" form:"hard_level"`
	Score        int    `json:"score" form:"score"`
	Tag          string `json:"tag" form:"tag"`
}

type DeleteQuestionMsg struct {
	IDs string `json:"question_ids"`
}

type EditQuestionMsg struct {
	ID           int    `json:"ID" form:"ID"`
	Question     string `json:"question" form:"question"`
	QuestionType string `json:"question_type" form:"question_type"`
	Options      string `json:"options" form:"options"`
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

type AddNotificationMsg struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AddFeedbackMsg struct {
	Content string `json:"content"`
}

type AutoCreatePaperMsg struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Tag               string `json:"tag"`
	SingleChoiceCount int    `json:"single_choice_count"`
	MultiChoiceCount  int    `json:"multi_choice_count"`
	TrueFalseCount    int    `json:"true_false_count"`
	FillBlankCount    int    `json:"fill_blank_count"`
	ShortAnswerCount  int    `json:"short_answer_count"`
}

type ManualCreatePaperMsg struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	QuestionIds []int  `json:"question_ids"`
}

type DeletePaperMsg struct {
	IDs string `json:"paper_ids"`
}

type ExportPaperMsg struct {
	ID string `json:"paper_id"`
}
