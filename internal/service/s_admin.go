package service

import (
	"ciel-begin/internal/consts"
	"ciel-begin/internal/model/bo"
	"ciel-begin/internal/service/internal/dao"
	"ciel-begin/utility/utils/xpwd"
	"context"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
)

type admin struct{}

func Admin() *admin {
	return &admin{}
}
func (s *admin) Login(ctx context.Context, uname string, pwd string) error {
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
	menus, err := Role().Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	if err = Session().SetAdmin(ctx, &bo.Admin{Admin: admin, Menus: menus}); err != nil {
		return err
	}
	return nil
}
func (s *admin) Logout(ctx context.Context) error {
	return Session().RemoveAdmin(ctx)
}
func (s *admin) UpdateAdminPwd(ctx context.Context, pwd string, pwd2 string) error {
	admin, err := Session().GetAdmin(ghttp.RequestFromCtx(ctx))
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
	err = Session().RemoveAdmin(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}
