package v1

import (
	"ciel-admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type WalletSetPassReq struct {
	g.Meta `tags:"钱包" dc:"设置钱包密码"`
	Pass   string `json:"pass" v:"required" dc:"6位数字"`
}
type WalletUpdatePassReq struct {
	g.Meta  `tags:"钱包" dc:"修改钱包密码"`
	OldPass string `json:"old_pass" v:"required" dc:"旧密码"`
	NewPass string `json:"new_pass" v:"required" dc:"新密码"`
}
type TopUpCategoryReq struct {
	g.Meta `tags:"钱包" dc:"充值类型"`
	*Authorization
}
type CreateTopUpReq struct {
	g.Meta       `tags:"钱包" dc:"创建充值订单"`
	ChangeTypeId int     `json:"change_type_id" dc:"充值类型(从支持的充值类型中选择id)" v:"required"`
	Money        float64 `json:"money" dc:"金额" v:"required"`
	*Authorization
}
type ListTopUpReq struct {
	g.Meta `tags:"钱包" dc:"查询充值记录"`
	*Authorization
	*PageReq
	Status int `json:"status" dc:"状态 0全部 1等待 2成功 3失败" d:"0"`
}
type ListChangeTypesReq struct {
	g.Meta `tags:"钱包" dc:"账变类型"`
	*Authorization
}
type ListChangeLogReq struct {
	g.Meta `tags:"钱包" dc:"查询账变记录"`
	*Authorization
	*PageReq
	Type int `json:"type" dc:"账变类型：具体查询账变类型接口"`
}
type WalletInfoReq struct {
	g.Meta `tags:"钱包" dc:"查询钱包信息"`
	*Authorization
}
type WalletInfoRes struct {
	Id           uint64  `json:"id" dc:"id"`
	Uid          uint64  `json:"uid" dc:"用户ID"`
	Balance      float64 `json:"balance" dc:"余额"`
	PassErrCount int     `json:"pass_err_count" dc:"密码错误次数"`
}
type ListChangeLogRes struct {
	*PageRes
	Items []*model.ChangeLogItem `json:"items"`
}

type ListChangeTypesRes struct {
	Id    int    `json:"id" dc:"ID"`
	Title string `json:"title" dc:"账变名称"`
	Class string `json:"class" dc:"样式类名"`
}
type ListTopUpRes struct {
	*PageRes
	Items []*model.TopUpItem `json:"items"`
}

type TopUpCategoryRes struct {
	Id    int    `json:"id" dc:"id"`
	Title string `json:"title" dc:"名称"`
	Desc  string `json:"desc" dc:"说明"`
}
