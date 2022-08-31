package logic

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/service/view"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"regexp"
	"strings"
)

type role struct {
}

func (r role) ClearApi(ctx context.Context, rid interface{}) error {
	_, err := dao.Role.GetById(ctx, rid)
	if err != nil {
		return err
	}
	_, err = dao.RoleApi.Ctx(ctx).Delete("rid", rid)
	return err
}

func (r role) CheckRoleApi(ctx context.Context, rid int, uri string) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}
	if uri == "/" {
		return true
	}
	s := fmt.Sprint(regexp.MustCompile(".+/del/").FindString(uri), ":id")
	if s != ":id" {
		uri = s
	}
	count, _ := g.DB().Ctx(ctx).Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and  t1.id = ?  ", uri, rid).
		Count()
	if count == 1 {
		return false
	}
	return true
}

func (r role) Roles(ctx context.Context) (string, error) {
	var (
		array = make([]string, 0)
	)
	all, err := dao.Role.Ctx(ctx).All()
	if err != nil {
		return "", err
	}
	for index, m := range all {
		id := m["id"]
		name := m["name"]
		array = append(array, fmt.Sprintf(fmt.Sprintf("%v:%v:%s", id, name, view.SwitchTagClass(index))))
	}
	return strings.Join(array, ","), nil
}

var (
	Role = role{}
)
