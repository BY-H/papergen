package message

type RequestMsg struct {
	DateStart string `json:"date_start" form:"date_start"`
	DateEnd   string `json:"date_end" form:"date_end"`
}

type OrderMsg struct {
	RequestMsg
}
