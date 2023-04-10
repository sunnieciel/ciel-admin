package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// errors

var (
	ErrAuthNotEnough    = gerror.NewCode(gcode.New(-1, "暂无当前操作权限", ""))
	ErrAuth             = gerror.NewCode(gcode.New(-2, "未认证", ""))
	ErrDataNotFound     = gerror.NewCode(gcode.New(-3, "数据不存在", ""))
	ErrImgCannotBeEmpty = gerror.NewCode(gcode.New(-4, "图片不能为空", ""))
	ErrDataAlreadyExist = gerror.NewCode(gcode.New(-5, "数据已存在", ""))
	ErrActionFast       = gerror.NewCode(gcode.New(-6, "您操作的太快啦", ""))
	ErrBalance          = gerror.NewCode(gcode.New(-7, "用户余额错误", ""))
	ErrCaptcha          = gerror.NewCode(gcode.New(-8, "验证码错误", ""))

	ErrTopUpStatusIsNotWait = gerror.NewCode(gcode.New(-100, "充值订单状态错误", ""))

	ErrLogin             = gerror.NewCode(gcode.New(-200, "用户名或密码错误", ""))
	ErrPassEmpty         = gerror.NewCode(gcode.New(-201, "密码不能为空", ""))
	ErrFormatEmail       = gerror.NewCode(gcode.New(-202, "邮箱格式不正确", ""))
	ErrUnameExist        = gerror.NewCode(gcode.New(-203, "用户名已存在", ""))
	ErrUnameFormat       = gerror.NewCode(gcode.New(-204, "用户名长度在4到12位之间", ""))
	ErrPassFormat        = gerror.NewCode(gcode.New(-205, "密码格式为任意可见字符，长度在6~18之间", ""))
	ErrOldPass           = gerror.NewCode(gcode.New(-206, "旧密码错误", ""))
	ErrHasOrderNotFinish = gerror.NewCode(gcode.New(-207, "还有未完成的订单", ""))
)
