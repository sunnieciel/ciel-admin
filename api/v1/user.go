package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta `tags:"用户" summary:"注册"`
	Uname  string `v:"required#用户名不能为空" dc:"用户名"`
	Pass   string `v:"required|password#密码不能为空｜密码格式不正确" dc:"密码"`
}
type LoginReq struct {
	g.Meta `tags:"用户" summary:"登录"`
	Uname  string `v:"required#用户名不能为空" dc:"用户名"`
	Pass   string `v:"required|password#密码不能为空｜密码格式不正确" dc:"密码"`
}
type UserInfoReq struct {
	g.Meta `tags:"用户" summary:"用户登录信息"  `
	Authorization
}
type UpdatePassReq struct {
	g.Meta  `tags:"用户" sm:"修改密码"`
	OldPass string `v:"required|password#旧密码不能为空|密码格式不正确" dc:"旧密码"`
	NewPass string `v:"required|password#新密码不能为空|密码格式不正确" dc:"新密码"`
}
type UpdateNicknameReq struct {
	g.Meta   `tags:"用户" sm:"修改昵称"`
	Nickname string `v:"required" dc:"昵称"`
}
type IconsReq struct {
	g.Meta `tags:"用户" sm:"可选头像"`
}
type UpdateIconReq struct {
	g.Meta `tags:"用户" sm:"修改头像"`
	Icon   string `json:"icon" dc:"图片(不需要前缀,值从可选头像接口获取)"`
}

type LoginRes struct {
	Uname        string `json:"uname" dc:"用户名"`
	Nickname     string `json:"nickname" dc:"昵称"`
	Icon         string `json:"icon" dc:"头像"`
	Summary      string `json:"summary" dc:"简介"`
	Email        string `json:"email" dc:"邮箱"`
	Phone        string `json:"phone" dc:"电话号码"`
	WalletStatus uint   `json:"wallet_status" dc:"钱包状态"`
	Token        string `json:"token" dc:"token"`
}
type IconsRes struct {
	Icons     []string `json:"icons" dc:"图片"`
	ImgPrefix string   `json:"img_prefix" dc:"图片前缀"`
}
