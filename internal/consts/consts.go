// Package consts 常量
package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	ImgPrefix string
	WhiteIps  string
)

const (
	MsgPrimary = `<div class="msg-primary" onclick="$(this).hide(200)"> <li class="fa fa-exclamation-triangle"></li>%s</div> `
	MsgWarning = `<div class="msg-warning" onclick="$(this).hide(200)"> <li class="fa fa fa-exclamation"></li>%s</div>`
)
const (
	AdminUnreadKey = "unreadNum"
)

// user
const (
	UserStatusOk    = 1 // 用户状态正常
	UserStatusError = 2 // 用户状态禁用
	UserStatusLock  = 3 // 用户被锁定
	UidKey          = "userInfoKey"
)

// wallet

const (
	ApplicationStatusWait    = 1
	ApplicationStatusSuccess = 2
	ApplicationStatusFail    = 3
)

// errors

var (
	// sys

	ErrAuthNotEnough    = gerror.NewCode(gcode.New(-1, "暂无当前操作权限", ""))
	ErrAuth             = gerror.NewCode(gcode.New(-2, "未认证", ""))
	ErrDataNotFound     = gerror.NewCode(gcode.New(-3, "数据不存在", ""))
	ErrImgCannotBeEmpty = gerror.NewCode(gcode.New(-4, "图片不能为空", ""))

	//  用户相关

	ErrLogin                   = gerror.NewCode(gcode.New(-100, "用户名或密码错误", ""))
	ErrPassEmpty               = gerror.NewCode(gcode.New(-101, "密码不能为空", ""))
	ErrFormatEmail             = gerror.NewCode(gcode.New(-102, "邮箱格式不正确", ""))
	ErrUnameExist              = gerror.NewCode(gcode.New(-103, "用户名已存在", ""))
	ErrUnameFormat             = gerror.NewCode(gcode.New(-104, "用户名长度在4到12位之间", ""))
	ErrPassFormat              = gerror.NewCode(gcode.New(-105, "密码格式为任意可见字符，长度在6~18之间", ""))
	ErrPassErrorTooMany        = gerror.NewCode(gcode.New(-106, "密码错误次数太多", ""))
	ErrOldPassNotMatch         = gerror.NewCode(gcode.New(-107, "旧密码不正确", ""))
	ErrNicknameEmpty           = gerror.NewCode(gcode.New(-108, "昵称不能为空", ""))
	ErrMaxLengthSixTy          = gerror.NewCode(gcode.New(-109, "允许的最长字符为16", ""))
	ErrIconEmpty               = gerror.NewCode(gcode.New(-110, "图片不能为空", ""))
	ErrUserDoesNotExist        = gerror.NewCode(gcode.New(-111, "用户不存在", ""))
	ErrUseWalletPassAlreadySet = gerror.NewCode(gcode.New(-112, "用户钱包密码已设置", ""))
	ErrTopUpOrderAlreadyHas    = gerror.NewCode(gcode.New(-113, "已有未完成的充值订单", ""))
	ErrMinTopUpOrderMoney      = gerror.NewCode(gcode.New(-114, "充值订单金额在10~10000之间", ""))
	ErrTopUpType               = gerror.NewCode(gcode.New(-115, "错误的充值类型", ""))
	ErrUserNotFound            = gerror.NewCode(gcode.New(-116, "用户数据不存在", ""))

	// About format

	ErrFormatNotNumber     = gerror.NewCode(gcode.New(-200, "格式不正确，请输入数字", ""))
	ErrFormatKeepLengthSix = gerror.NewCode(gcode.New(-201, "请将长度保持为6位", ""))
)
