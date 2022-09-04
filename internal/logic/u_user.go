package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/utility/utils/xpwd"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = user{}
)

type user struct {
}

func (u user) UpdateUname(ctx context.Context, uname string, id uint64) error {
	count, err := dao.User.Ctx(ctx).Count("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if err = g.Validator().Rules("password").Data(uname).Run(ctx); err != nil {
		return consts.ErrUnameFormat
	}
	if _, err = dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"uname": uname}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (u user) UpdatePass(ctx context.Context, pass string, id uint64) error {
	if err := g.Validator().Rules("password").Data(pass).Run(ctx); err != nil {
		return consts.ErrPassFormat
	}
	if _, err := dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"pass": xpwd.GenPwd(pass)}); err != nil {
		g.Log().Error(ctx)
		return err
	}
	return nil
}
