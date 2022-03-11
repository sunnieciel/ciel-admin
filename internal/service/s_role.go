package service

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

type sRole struct{}

var insRole = &sRole{}

func Role() *sRole {
	return insRole
}
func (s *sRole) RoleNoMenu(ctx context.Context, rid interface{}) (interface{}, error) {
	return dao.RoleMenu.RoleNoMenu(ctx, rid)
}
func (s *sRole) AddRoleMenu(ctx context.Context, rid int, mid []int) error {
	return dao.RoleMenu.AddRoleMenu(ctx, rid, mid)
}
func (s *sRole) RoleNoApi(ctx context.Context, rid interface{}) (gdb.List, error) {
	return dao.RoleApi.RoleNoApi(ctx, rid)
}
func (s *sRole) AddRoleApi(ctx context.Context, rid int, aid []int) error {
	return dao.RoleApi.AddRoleApi(ctx, rid, aid)
}
func (s *sRole) CheckRoleApi(ctx context.Context, rid int, uri string, method string) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}
	if uri == "/" {
		return true
	}
	count, _ := g.DB().Ctx(ctx).Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and t3.method = ? and t1.id = ?  ", uri, method, rid).
		Count()
	if count == 1 {
		return false
	}
	return true
}
func (s *sRole) Menus(ctx context.Context, rid int, pid int) ([]*bo.Menu, error) {
	return dao.RoleMenu.Menus(ctx, rid, pid)
}
