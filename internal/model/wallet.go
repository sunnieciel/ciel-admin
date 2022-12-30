package model

import "github.com/gogf/gf/v2/os/gtime"

type TopUpCategory struct {
	Id    int    `json:"id" dc:"id"`
	Title string `json:"title" dc:"名称"`
	Desc  string `json:"desc" dc:"说明"`
}
type TopUpItem struct {
	Id          uint64     `json:"id" dc:"id"`
	TransId     string     `json:"trans_id" dc:"交易ID"`
	Uid         uint64     `json:"uid" dc:"用户ID"`
	ChangeType  int        `json:"change_type" dc:"账变类型： 对应值请查询账变类型接口"`
	Description string     `json:"description" dc:"说明"`
	Status      int        `json:"status" dc:"状态 1等待 2成功 3失败"`
	CreatedAt   gtime.Time `json:"created_at" dc:"交易时间"`
}
type ChangeLogItem struct {
	Id        uint64     `json:"id" dc:"id"`
	TransId   string     `json:"trans_id" dc:"交易ID"`
	Amount    float64    `json:"amount" dc:"交易金额"`
	Balance   float64    `json:"balance" dc:"余额"`
	Type      int        `json:"type" dc:"账变类型：对应值请查询账变类型接口"`
	Desc      string     `json:"desc" dc:"说明"`
	CreatedAt gtime.Time `json:"created_at" dc:"交易时间"`
}
