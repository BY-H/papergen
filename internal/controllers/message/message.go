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
