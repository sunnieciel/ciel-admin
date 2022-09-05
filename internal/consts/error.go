package consts

import "errors"

var (
	ErrAuthNotEnough = errors.New("暂无当前操作权限")
	ErrAuth          = errors.New("未认证")
	ErrDataNotFound  = errors.New("数据不存在")

	//  用户相关

	ErrNotAuth           = errors.New("未登录")
	ErrLogin             = errors.New("用户名或密码错误")
	ErrPassEmpty         = errors.New("密码不能为空")
	ErrFormatEmail       = errors.New("邮箱格式不正确")
	ErrUnameAlreadyExist = errors.New("用户名已存在")
	ErrUnameExist        = errors.New("用户名已存在")
	ErrUnameFormat       = errors.New("用户名格式不正确")
	ErrPassFormat        = errors.New("密码格式不正确")
	ErrPassErrorTooMany  = errors.New("密码错误次数太多")
	ErrTwoPassNotMatch   = errors.New("两次输入的密码不匹配")
	ErrOldPassNotMatch   = errors.New("旧密码不正确")
	ErrNicknameEmpty     = errors.New("昵称不能为空")

	ErrMaxLengthSixTy = errors.New("允许的最长字符为16")

	ErrIconEmpty = errors.New("图片不能为空")
)
