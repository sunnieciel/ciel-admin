package role

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

func ClearApi(ctx context.Context, rid interface{}, t int) error {
	return logic.Role.ClearApi(ctx, rid, t)
}
func NoMenu(ctx context.Context, rid interface{}) (gdb.List, error) {
	return dao.RoleMenu.RoleNoMenu(ctx, rid)
}
func AddRoleMenu(ctx context.Context, rid int, mid []int) error {
	return dao.RoleMenu.AddRoleMenu(ctx, rid, mid)
}
func NoApi(ctx context.Context, rid interface{}) (gdb.List, error) {
	return dao.RoleApi.RoleNoApi(ctx, rid)
}
func AddRoleApi(ctx context.Context, rid int, aid []int) error {
	return dao.RoleApi.AddRoleApi(ctx, rid, aid)
}
func Roles(ctx context.Context) (string, error) {
	return logic.Role.Roles(ctx)
}

func GetById(ctx context.Context, id interface{}) (*entity.Role, error) {
	return dao.Role.GetById(ctx, id)
}
func Clear(ctx context.Context, rid interface{}) error {
	return logic.Role.ClearMenu(ctx, rid)
}
