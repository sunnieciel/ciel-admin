package admin

import (
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/bo"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Login(ctx context.Context, id, code, uname string, pwd string, ip string) error {
	return logic.Admin.Login(ctx, id, code, uname, pwd, ip)
}

func GetFromSession(r *ghttp.Session) (*bo.Admin, error) {
	return logic.Session.GetAdmin(r)
}

func Logout(ctx context.Context) error {
	return logic.Session.RemoveAdminFromSession(ctx)
}

func UpdatePwd(ctx context.Context, pwd string, pwd2 string) error {
	return logic.Admin.UpdatePwd(ctx, pwd, pwd2)
}

func UpdateUname(ctx context.Context, id, uname interface{}) error {
	return logic.Admin.UpdateUname(ctx, id, uname)
}

func UpdatePwdWithoutOldPwd(ctx context.Context, id, pwd interface{}) error {
	return logic.Admin.UpdatePwdWithoutOldPwd(ctx, id, pwd)
}

func ClearLoginLog(ctx context.Context) error {
	return logic.Admin.ClearAllLoginLog(ctx)
}
