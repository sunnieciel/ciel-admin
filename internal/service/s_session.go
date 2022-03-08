package service

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---session ------------------------------------------------------------

const AdminSessionKey = "adminInfo"

type session struct{}

func Session() *session { return &session{} }
func (s session) SetAdmin(ctx context.Context, data *bo.Admin) error {
	return g.RequestFromCtx(ctx).Session.Set(AdminSessionKey, data)
}
func (s session) GetAdmin(r *ghttp.Request) (*bo.Admin, error) {
	get, err := r.Session.Get(AdminSessionKey)
	var data *bo.Admin
	err = get.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s session) RemoveAdmin(ctx context.Context) error {
	return g.RequestFromCtx(ctx).Session.Remove(AdminSessionKey)
}
func (s session) AdminIsLogin(r *ghttp.Request) error {
	user, err := s.GetAdmin(r)
	if err != nil {
		return err
	}
	if user == nil {
		return consts.ErrNotAuth
	}
	return nil
}
