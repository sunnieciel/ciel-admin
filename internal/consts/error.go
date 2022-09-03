package consts

import "errors"

var (
	ErrAuthNotEnough     = errors.New("暂无当前操作权限")
	ErrAuth              = errors.New("未认证")
	ErrLogin             = errors.New("用户名或密码错误")
	ErrPassEmpty         = errors.New("密码不能为空")
	ErrFormatEmail       = errors.New("邮箱格式不正确")
	ErrUnameAlreadyExist = errors.New("用户名已存在")
	ErrDataNotFound      = errors.New("数据不存在")
	ErrNotAuth           = errors.New("未登录")
	ErrUnameExist        = errors.New("用户名已存在")
)
