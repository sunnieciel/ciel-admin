package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/do"
	"ciel-admin/utility/utils/xcaptcha"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Admin = admin{}
)

type admin struct {
}

func (admin) Menus(ctx context.Context, rid int, pid int) ([]*bo.Menu, error) {
	var d = make([]*bo.Menu, 0)
	menus, err := dao.RoleMenu.Menus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}

func (a admin) Login(ctx context.Context, id string, code string, uname string, pwd string, ip string) error {
	if !xcaptcha.Store.Verify(id, code, true) {
		return errors.New("验证码错误")
	}
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
	menus, err := Admin.Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	adminInfo := bo.Admin{Admin: admin, Menus: menus}
	if err = Session.SetAdmin(ctx, &adminInfo); err != nil {
		return err
	}
	if _, err = dao.AdminLoginLog.Ctx(ctx).Insert(do.AdminLoginLog{Uid: admin.Id, Ip: ip}); err != nil {
		return err
	}
	return nil
}

func (a admin) UpdatePwd(ctx context.Context, pwd string, pwd2 string) error {
	adminBo, err := Session.GetAdmin(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return err
	}
	u, err := dao.Admin.GetByUname(ctx, adminBo.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = Session.RemoveAdminFromSession(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}

func (a admin) UpdateUname(ctx context.Context, id interface{}, uname interface{}) error {
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

func (a admin) UpdatePwdWithoutOldPwd(ctx context.Context, id interface{}, pwd interface{}) error {
	_, err := dao.Admin.Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd.(string))}, "id", id)
	if err != nil {
		return err
	}
	return nil
}

func (a admin) ClearAllLoginLog(ctx context.Context) error {
	if _, err := dao.AdminLoginLog.Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}
	return nil
}
