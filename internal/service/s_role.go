package service

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/internal/dao"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
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
	var d = make([]*bo.Menu, 0)
	get, err := g.Cfg().Get(ctx, "rss")
	if err != nil {
		return nil, err
	}
	array := get.Array()
	if len(array) > 0 {
		children := make([]*bo.Menu, 0)
		d = append(d, &bo.Menu{
			Name: "首页",
			Children: []*bo.Menu{
				{Name: "V2EX", Path: "/"},
				{Name: "Github", Path: "/sys/path/github"},
				{Name: "豆瓣阅读", Path: "/sys/path/douban"},
				{Name: "开源中国", Path: "/sys/path/oschina"},
			},
		})
		for _, item := range array {
			split := gstr.Split(fmt.Sprint(item), ":")
			children = append(children, &bo.Menu{
				Name: split[0],
				Path: split[1],
			})
		}
	}
	menus, err := dao.RoleMenu.Menus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}
