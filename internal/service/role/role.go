package role

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

func ClearApi(ctx context.Context, rid interface{}) error {
	return logic.Role.ClearApi(ctx, rid)
}
func NoMenu(ctx context.Context, rid interface{}) (interface{}, error) {
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
func CheckRoleApi(ctx context.Context, rid int, uri string) bool {
	return logic.Role.CheckRoleApi(ctx, rid, uri)
}
func Roles(ctx context.Context) (string, error) {
	return logic.Role.Roles(ctx)
}
