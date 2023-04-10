package v1

import (
	"freekey-backend/api"
	"freekey-backend/internal/model"
	"freekey-backend/internal/model/do"
	"freekey-backend/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta `tags:"用户" summary:"注册"`
	Uname  string `v:"required#用户名不能为空" dc:"用户名" json:"uname"`
	Pass   string `v:"required|password#密码不能为空｜密码格式不正确" dc:"密码" json:"pass"`
}
type LoginReq struct {
	g.Meta `tags:"用户" summary:"登录"`
	Uname  string `v:"required#用户名不能为空" dc:"用户名" json:"uname"`
	Pass   string `v:"required|password#密码不能为空|密码格式不正确" dc:"密码" json:"pass"`
}
type UserInfoReq struct {
	g.Meta `tags:"用户" summary:"用户登录信息"  `
	api.Authorization
}
type UpdatePassReq struct {
	g.Meta  `tags:"用户" sm:"修改密码"`
	OldPass string `v:"required|password#旧密码不能为空|密码格式不正确" dc:"旧密码" json:"oldPass"`
	NewPass string `v:"required|password#新密码不能为空|密码格式不正确" dc:"新密码" json:"newPass"`
}
type UpdateNicknameReq struct {
	g.Meta   `tags:"用户" sm:"修改昵称"`
	Nickname string `v:"required" dc:"昵称" json:"nickname"`
}

type LoginRes struct {
	Token string `json:"token" dc:"token"`
}
type GetUserInfoReq struct {
	api.Authorization
	g.Meta `tags:"用户" sm:"获取用户信息"`
}
type GetUserInfoRes struct {
	Id       uint64  `json:"id"`
	Uname    string  `json:"uname" dc:"用户名"`
	Nickname string  `json:"nickname" dc:"昵称"`
	Icon     string  `json:"icon" dc:"头像"`
	Summary  string  `json:"summary" dc:"简介"`
	Email    string  `json:"email" dc:"邮箱"`
	Phone    string  `json:"phone" dc:"电话号码"`
	Balance  float64 `json:"balance" dc:"用户金额"`
}

type UpdateUserPassReq struct {
	api.Authorization
	g.Meta  `tags:"用户" sm:"修改用户密码"`
	OldPass string `json:"old_pass" v:"required" sm:"旧密码"`
	Pass    string `json:"pass" v:"required|password" sm:"新密码"`
}

type UpdateIconReq struct {
	api.Authorization
	g.Meta `tags:"用户" sm:"修改头像"`
	Icon   string `json:"icon" v:"required"`
}

//--- TopUp ---------------------------------------------------------

type AddTopUpReq struct {
	g.Meta `tags:"后台" dc:"添加"`
	*do.TopUp
}
type GetTopUpReq struct {
	g.Meta `tags:"后台" dc:"查询一条数据"`
	Id     uint64 `v:"required"`
}
type GetTopUpRes struct {
	Data *entity.TopUp `json:"data"`
}
type ListTopUpReq struct {
	g.Meta `tags:"后台" dc:"查询列表数据"`
	api.PageReq
	Uname      string
	TransId    string
	ChangeType string
	Ip         string
	Desc       string
	Aid        string
	Status     string
}
type ListTopUpRes struct {
	List []*model.TopUp `json:"list"`
	*api.PageRes
}
type DelTopUpReq struct {
	g.Meta `tags:"后台" dc:"删除"`
	Id     uint64 `v:"required"`
}
type UpdateTopUpReq struct {
	g.Meta `tags:"后台" dc:"修改菜单"`
	*do.TopUp
}
type CreateTopUpReq struct {
	g.Meta `tags:"钱包" sm:"创建充值订单"`
	Type   int     `json:"type" sm:"充值类型 1 支付宝, 2 微信, 3 PayPal" v:"required|in:1,2,3"`
	Money  float64 `json:"money" v:"required" sm:"充值金额"`
	Desc   string  `json:"desc" sm:"备注"`
}
type UpdateTopUpByAdminReq struct {
	g.Meta `tags:"后台"`
	Type   int    `json:"type" v:"required|in:1,2"`
	Id     uint64 `json:"id" v:"required"`
}
type ListTopUpForWebReq struct {
	g.Meta `tags:"钱包" sm:"查询充值订单"`
	*api.PageReq
	Status int `json:"status" sm:"状态 1等待, 2成功, 3失败, 不传查询所有"`
}
type ListTopUpForWebRes struct {
	*api.PageRes
	List []*model.TopUpForWeb `json:"list"`
}
type ListWalletChangeLogForWebReq struct {
	api.Authorization
	g.Meta `tags:"钱包" sm:"查询账变记录"`
	*api.PageReq
	Type int `json:"type" dc:"交易类型"`
}
type ListWalletChangeLogForWebRes struct {
	*api.PageRes
	List []*model.WalletChangeLogForWeb `json:"list"`
}

type GetWalletReportReq struct {
	g.Meta `tags:"后台"`
	Uname  string
	Begin  string
	End    string
}
type GetWalletReportRes struct {
	T1  float64 `json:"t1"          description:""`
	T2  float64 `json:"t2"          description:""`
	T3  float64 `json:"t3"          description:""`
	T4  float64 `json:"t4"          description:""`
	T5  float64 `json:"t5"          description:""`
	T6  float64 `json:"t6"          description:""`
	T7  float64 `json:"t7"          description:""`
	T8  float64 `json:"t8"          description:""`
	T9  float64 `json:"t9"          description:""`
	T10 float64 `json:"t10"         description:""`
	T11 float64 `json:"t11"         description:""`
	T12 float64 `json:"t12"         description:""`
	T13 float64 `json:"t13"         description:""`
}
