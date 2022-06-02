package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	AdminSessionKey = "adminInfo"
	Uid             = "userInfoKey"
)

func SetAdmin(ctx context.Context, data *bo.Admin) error {
	return g.RequestFromCtx(ctx).Session.Set(AdminSessionKey, data)
}
func GetAdmin(r *ghttp.Request) (*bo.Admin, error) {
	get, err := r.Session.Get(AdminSessionKey)
	var data *bo.Admin
	err = get.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, err
}
func RemoveAdmin(ctx context.Context) error {
	return g.RequestFromCtx(ctx).Session.Remove(AdminSessionKey)
}
func AdminIsLogin(r *ghttp.Request) error {
	user, err := GetAdmin(r)
	if err != nil {
		return err
	}
	if user == nil {
		return consts.ErrNotAuth
	}
	return nil
}
