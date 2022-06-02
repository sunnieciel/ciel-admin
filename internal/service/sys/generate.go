package sys

import (
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/internal/dao"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
)

func Fields(ctx context.Context, tableName string) (map[string]*gdb.TableField, error) {
	fields, err := g.DB().Ctx(ctx).Model(tableName).TableFields(tableName)
	if err != nil {
		return nil, err
	}
	return fields, nil
}
func Tables(ctx context.Context) ([]string, error) {
	return g.DB().Tables(ctx)
}
func genApi(ctx context.Context, category string, name string) error {
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s/del", name), Method: "DELETE", Group: category, Desc: fmt.Sprintf("删除%s", name), Status: 1},
		{Url: fmt.Sprintf("/%s/post", name), Method: "POST", Group: category, Desc: fmt.Sprintf("添加%s", name), Status: 1},
		{Url: fmt.Sprintf("/%s/put", name), Method: "PUT", Group: category, Desc: fmt.Sprintf("修改%s", name), Status: 1},
	}
	for _, i := range array {
		count, err := dao.Api.Ctx(ctx).Count("url", i.Url)
		if err != nil {
			return err
		}
		if count != 0 {
			continue
		}
		if _, err = dao.Api.Ctx(ctx).Insert(i); err != nil {
			return err
		}
	}
	return nil
}
