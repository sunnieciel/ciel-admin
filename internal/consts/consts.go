// Package consts 常量
package consts

var (
	ImgPrefix string
	WhiteIps  string
)

const (
	MsgPrimary = `<div class="msg-primary" onclick="$(this).hide(200)"> <li class="fa fa-exclamation-triangle"></li>%s</div> `
	MsgWarning = `<div class="msg-warning" onclick="$(this).hide(200)"> <li class="fa fa fa-exclamation"></li>%s</div>`
)
const (
	UserStatusOk    = 1 // 用户状态正常
	UserStatusError = 2 // 用户状态禁用
	UserStatusLock  = 3 // 用户被锁定
)
