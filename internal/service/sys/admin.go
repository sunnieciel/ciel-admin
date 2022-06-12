package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Login(ctx context.Context, uname string, pwd string) error {
	admin, err := dao.Admin.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(admin.Pwd, pwd) {
		return consts.ErrLogin
	}

	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	menus, err := Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	if err = SetAdmin(ctx, &bo.Admin{Admin: admin, Menus: menus}); err != nil {
		return err
	}
	return nil
}
func Logout(ctx context.Context) error {
	return RemoveAdmin(ctx)
}
func UpdateAdminPwd(ctx context.Context, pwd string, pwd2 string) error {
	admin, err := GetAdmin(ghttp.RequestFromCtx(ctx))
	if err != nil {
		return err
	}
	u, err := dao.Admin.GetByUname(ctx, admin.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = RemoveAdmin(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}
func UpdateAdminUname(ctx context.Context, id, uname interface{}) error {
	count, err := dao.Admin.Ctx(ctx).Count("uname", uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if _, err = dao.Admin.Ctx(ctx).Update(g.Map{"uname": uname}, "id", id); err != nil {
		return err
	}
	return nil
}
func UpdateAdminPwdWithoutOldPwd(ctx context.Context, id, pwd interface{}) error {
	_, err := dao.Admin.Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd.(string))}, "id", id)
	if err != nil {
		return err
	}
	return nil
}
