package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Platform        string  `gorm:"column:platform;comment:'订单源平台'" json:"platform"`
	ReportDate      string  `gorm:"column:report_date;index:IDX_REPORT_DATE;comment:'下单日期'" json:"report_date"`
	AccountID       string  `gorm:"column:account_id;comment:'用户三方平台账号'" json:"account_id"`
	AccountPassword string  `gorm:"column:account_password;comment:'用户三方登录密码'" json:"account_password"`
	Url             string  `gorm:"column:url;comment:'用户url'" json:"url"`
	Status          string  `gorm:"column:status;comment:'订单状态'" json:"status"`
	Amount          int     `gorm:"column:amount;comment:'订单题数'" json:"amount"`
	Accuracy        float64 `gorm:"column:accuracy;comment:'订单正确率'" json:"accuracy"`
	Money           float64 `gorm:"column:money;comment:'订单金额'" json:"money"`
	Remark          string  `gorm:"column:remark;comment:'订单备注'" json:"remark"`
	CreatorID       string  `gorm:"column:creator_id;comment:'订单创建人'" json:"creator_id"`
	SolverID        string  `gorm:"column:solver_id;comment:'订单处理人'" json:"solver_id"`
}

const (
	STATUS_CLOSE    = "CLOSE"
	STATUS_RECORD   = "RECORD"
	STATUS_WORKING  = "WORKING"
	STATUS_FINISHED = "FINISHED"
	STATUS_CHECK    = "CHECK"
	STATUS_DOWN     = "DOWN"
)

const (
	PLATFORM_QINGMA = "PLATFORM_QINGMA"
)

func (o Order) CheckOrder() bool {
	o.formalCheck()
	switch o.Platform {
	case PLATFORM_QINGMA:
		return checkQingmaOrder(o)
	default:
		// 未知订单来源
		return false
	}
}

// 通用检测
func (o Order) formalCheck() bool {
	return true
}

func checkQingmaOrder(o Order) bool {
	// 青马渠道现在以 url 作为刷题依据，所以必须要有 url
	if o.Url == "" {
		return false
	}
	if o.Amount == 0 {
		return false
	}
	if o.Accuracy == 0 {
		return false
	}
	return true
}
