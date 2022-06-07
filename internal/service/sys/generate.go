package sys

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/internal/dao"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"sort"
)

func Fields(ctx context.Context, tableName string) ([]*gdb.TableField, error) {
	res := make([]*gdb.TableField, 0)
	fields, err := g.DB().Ctx(ctx).Model(tableName).TableFields(tableName)
	if err != nil {
		return nil, err
	}
	for _, field := range fields {
		res = append(res, field)
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Index < res[j].Index
	})
	return res, nil
}
func Tables(ctx context.Context) ([]string, error) {
	return g.DB().Tables(ctx)
}
func GenFile(ctx context.Context, d *bo.GenConf) (err error) {
	//genController()
	//genRouter()
	//if err = genHtml(ctx, d); err != nil {
	//	return err
	//}
	//if err = genApi(ctx, "", ""); err != nil {
	//	return err
	//}
	return
}

func genHtml(ctx context.Context, c *bo.GenConf) error {
	rootPath := g.Config().MustGet(ctx, "server.rootPath").String()
	htmlGroup := c.HtmlGroup
	if htmlGroup == "" {
		return errors.New("html group can't be empty")
	}
	table := c.Table
	if table == "" {
		return errors.New("table can't ben empty")
	}
	filePath := fmt.Sprint(rootPath, "resource/template/", htmlGroup, "/", table, ".html")
	temp := gfile.GetContents(fmt.Sprintf("%s/resource/gen/templ.html", rootPath))
	// replace name and desc
	if c.PageName == "" {
		return errors.New("name can't be empty")
	}
	if c.PageDesc != "" {
		temp = gstr.Replace(temp, "这里是[name]管理页面,可以对菜单进行添加,修改,删除等操作。", c.PageDesc)
	}
	temp = gstr.Replace(temp, "[name]", c.PageName)
	// replace url
	if c.UrlPrefix == "" {
		return errors.New("urlPrefix can't be empty")
	}
	temp = gstr.Replace(temp, "[urlPrefix]", c.UrlPrefix)
	f, err := gfile.Create(filePath)
	if err != nil {
		return err
	}
	// replace fields
	if len(c.Fields) == 0 {
		return errors.New("fields can't be empty")
	}
	fields := ""
	for index, i := range c.Fields {
		field := "{"
		switch i.Name {
		case "id", "created_at", "updated_at":
			continue
		}
		if i.Label == "" {
			i.Label = i.Name
		}
		// base
		field += fmt.Sprintf(`field: "%s", label: "%s"`, i.Name, i.Label)
		// append search
		if i.SelectType != 0 {
			field += fmt.Sprintf(",search:1")
		}
		// append select options
		if i.FieldType == "select" {
			field += ",type:'select',options:["
			if len(i.Options) == 0 {
				i.Options = make([]*bo.FieldOption, 0)
				i.Options = append(i.Options, &bo.FieldOption{Value: "", Label: "请选择", Type: "info"}, &bo.FieldOption{Value: 1, Label: "正常", Type: "primary"}, &bo.FieldOption{Value: 2, Label: "禁用", Type: "danger"})
			}
			for _, j := range i.Options {
				field += fmt.Sprintf(`{value: "%v", label: "%v", type: "%v"},`, j.Value, j.Label, j.Type)
			}
			field += "]"
		}
		field += "}"
		if index != len(fields) {
			field += ","
		}
		fields += field
	}
	temp = gstr.Replace(temp, "[fields],", fields)
	if _, err = f.WriteString(temp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func genController() {

}
func genRouter() {

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
