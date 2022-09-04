package user

import (
	"ciel-admin/internal/logic"
	"context"
)

func UpdateUname(ctx context.Context, uname string, id uint64) error {
	return logic.User.UpdateUname(ctx, uname, id)
}
func UpdatePass(ctx context.Context, pass string, id uint64) error {
	return logic.User.UpdatePass(ctx, pass, id)
}
