package model

import "github.com/gogf/gf/v2/os/gtime"

type UserLoginLog struct {
	Id        uint64      `json:"id"        description:""`
	Uid       uint64      `json:"uid"       description:""`
	Uname     string      `json:"uname"`
	Ip        string      `json:"ip"        description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" description:""`
}

type Wallet struct {
	Id           uint64      `json:"id"             description:""`
	Uname        string      `json:"uname"`
	Uid          uint64      `json:"uid"            description:""`
	Balance      float64     `json:"balance"        description:""`
	Pass         string      `json:"pass"           description:""`
	PassErrCount uint        `json:"passErrCount"   description:"密码输错次数"`
	Desc         string      `json:"desc"           description:""`
	CreatedAt    *gtime.Time `json:"createdAt"      description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"      description:""`
}

type WalletChangeLog struct {
	Id        uint64      `json:"id"        description:""`
	TransId   string      `json:"transId"   description:""`
	Uname     string      `json:"uname"`
	Amount    float64     `json:"amount"    description:""`
	Balance   float64     `json:"balance"   description:""`
	Type      uint        `json:"type"      description:"1人工充值,2支付宝充值,3微信充值,4paypal充值,5人工扣除"`
	Desc      string      `json:"desc"      description:""`
	CreatedAt *gtime.Time `json:"createdAt" description:""`
}

type WalletChangeLogForWeb struct {
	Id        uint64      `json:"id"        description:"id"`
	TransId   string      `json:"transId"   description:"交易ID"`
	Uid       uint64      `json:"uid" dc:"用户ID"`
	Amount    float64     `json:"amount"    description:"交易金额"`
	Balance   float64     `json:"balance"   description:"余额"`
	Type      uint        `json:"type"      description:"1人工充值,2支付宝充值,3微信充值,4paypal充值,5人工扣除"`
	Desc      string      `json:"desc"      description:"备注"`
	CreatedAt *gtime.Time `json:"createdAt" description:"交易时间"`
}

type WalletStatisticsLog struct {
	Id          uint64      `json:"id"          description:""`
	Uname       string      `json:"uname"`
	T1          float64     `json:"t1"          description:""`
	T2          float64     `json:"t2"          description:""`
	T3          float64     `json:"t3"          description:""`
	T4          float64     `json:"t4"          description:""`
	T5          float64     `json:"t5"          description:""`
	T6          float64     `json:"t6"          description:""`
	T7          float64     `json:"t7"          description:""`
	T8          float64     `json:"t8"          description:""`
	T9          float64     `json:"t9"          description:""`
	T10         float64     `json:"t10"         description:""`
	T11         float64     `json:"t11"         description:""`
	T12         float64     `json:"t12"         description:""`
	T13         float64     `json:"t13"         description:""`
	CreatedDate *gtime.Time `json:"createdDate" description:""`
}

type TopUp struct {
	Id         uint64      `json:"id"         description:""`
	Uname      string      `json:"uname"`
	TransId    string      `json:"transId"    description:"交易id"`
	Money      float64     `json:"money"      description:""`
	ChangeType uint        `json:"changeType" description:"账变类型 最好配置与此表 u_wallet_change_type 中的相对应。 1 支付宝充值, 2 微信充值, 3 Paypal充值"`
	Ip         string      `json:"ip"         description:"用户操作ip"`
	Desc       string      `json:"desc"       description:"备注"`
	Status     uint        `json:"status"     description:"订单状态 1 等待, 2 成功, 3 失败"`
	Aid        uint64      `json:"aid"        description:"管理员id"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改时间"`
}

type TopUpForWeb struct {
	Id         uint64      `json:"id"         description:""`
	Uid        uint64      `json:"uid" dc:"用户ID"`
	TransId    string      `json:"transId"    description:"交易id"`
	Money      float64     `json:"money"      description:""`
	ChangeType uint        `json:"changeType" description:"账变类型 1 支付宝充值, 2 微信充值, 3 Paypal充值"`
	Desc       string      `json:"desc"       description:"备注"`
	Status     uint        `json:"status"     description:"订单状态 1 等待, 2 成功, 3 失败"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:"修改时间"`
}
