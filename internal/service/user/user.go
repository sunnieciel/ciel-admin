package user

import (
	"ciel-admin/apiv1"
	"ciel-admin/internal/logic"
	"context"
)

func UpdateUname(ctx context.Context, uname string, id uint64) error {
	return logic.User.UpdateUname(ctx, uname, id)
}
func UpdatePass(ctx context.Context, pass string, id uint64) error {
	return logic.User.UpdatePass(ctx, pass, id)
}

func Register(ctx context.Context, uname, pass, ip string) (*apiv1.LoginVo, error) {
	return logic.User.Register(ctx, uname, pass, ip)
}
func Login(ctx context.Context, uname, pass, ip string) (*apiv1.LoginVo, error) {
	return logic.User.Login(ctx, uname, pass, ip)
}
