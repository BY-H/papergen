package order

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Platform        string  `gorm:"column:platform;comment:'订单源平台'"`
	AccountID       string  `gorm:"column:account_id;comment:'用户三方平台账号'"`
	AccountPassword string  `gorm:"column:account_password;comment:'用户三方登录密码'"`
	Url             string  `gorm:"column:url;comment:'用户url'"`
	Status          int     `gorm:"column:status;comment:'订单状态';default:0"`
	Amount          int     `gorm:"column:amount;comment:'订单题数'"`
	Accuracy        float64 `gorm:"column:accuracy;comment:'订单正确率'"`
	Money           float64 `gorm:"column:money;comment:'订单金额'"`
	Remark          string  `gorm:"column:remark;comment:'订单备注'"`
	CreatorID       string  `gorm:"column:creator_id;comment:'订单创建人'"`
	SolverID        string  `gorm:"column:solver_id;comment:'订单处理人'"`
}

const (
	STATUS_CLOSE = iota
	STATUS_RECORD
	STATUS_WORKING
	STATUS_FINISHED
	STATUS_CHECK
	STATUS_DOWN
)

var STATUS_MAPPING = map[int]string{
	STATUS_CLOSE:    "STATUS_CLOSE",    // 订单关闭
	STATUS_RECORD:   "STATUS_RECORD",   // 订单已记录
	STATUS_WORKING:  "STATUS_WORKING",  // 订单处理中
	STATUS_FINISHED: "STATUS_FINISHED", // 订单已完成
	STATUS_CHECK:    "STATUS_CHECK",    // 订单已审查
	STATUS_DOWN:     "STATUS_DOWN",     // 订单已分销
}

const (
	PLATFORM_QINGMA = iota
)

var PLATFORM_MAPPING = map[int]string{
	PLATFORM_QINGMA: "PLATFORM_QINGMA", //青马易站渠道
}

func (o Order) CheckOrder() bool {
	switch o.Platform {
	case PLATFORM_MAPPING[PLATFORM_QINGMA]:
		return checkQingmaOrder(o)
	default:
		// 未知订单来源
		return false
	}
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
