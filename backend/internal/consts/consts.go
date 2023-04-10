// Package consts 常量
package consts

var (
	ImgPrefix string
	WhiteIps  string
)

// user
const (
	TokenUserIdKey  = "userInfoKey"
	TokenAdminIdKey = "adminTokenKey"
	TokenAdminUname = "adminUname"
)

// wallet
const (
	ApplicationStatusWait    = 1
	ApplicationStatusSuccess = 2
	ApplicationStatusFail    = 3

	// 账变类型

	ChangeHumanTopUp  = 1 // 人工充值
	ChangeAliPayTopUp = 2 // 支付宝充值
	ChangeWxPayTopUp  = 3 // 微信充值
	ChangePayPalTopUp = 4 // Paypal 充值

)
