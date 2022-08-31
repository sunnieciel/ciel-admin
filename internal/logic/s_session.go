package logic

import (
	"ciel-admin/internal/model/bo"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	AdminSessionKey = "adminInfo"
	Uid             = "userInfoKey"
)

type session struct {
}

func (s session) SetAdmin(ctx context.Context, b *bo.Admin) error {
	return g.RequestFromCtx(ctx).Session.Set(AdminSessionKey, b)
}

func (session) GetAdmin(r *ghttp.Session) (*bo.Admin, error) {
	get, err := r.Get(AdminSessionKey)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, errors.New("admin info is nil")
	}
	var data *bo.Admin
	err = get.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s session) RemoveAdminFromSession(ctx context.Context) error {
	return g.RequestFromCtx(ctx).Session.Remove(AdminSessionKey)
}

var (
	Session = session{}
)
