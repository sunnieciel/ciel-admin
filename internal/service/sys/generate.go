package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/utility/utils/xfile"
	"ciel-admin/utility/utils/xicon"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"math"
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
	// gen menu
	if err = genMenu(ctx, d); err != nil {
		return err
	}
	//gen controller
	if err = genController(ctx, d); err != nil {
		return err
	}
	// gen router
	if err = genRouter(ctx, d); err != nil {
		return err
	}
	// gen html
	if err = genHtml(ctx, d); err != nil {
		return err
	}
	// gen api
	if err = genApi(ctx, d.HtmlGroup, d.StructName, d.PageName); err != nil {
		return err
	}
	return
}

func genMenu(ctx context.Context, d *bo.GenConf) error {
	menuLevel1 := d.MenuLevel1
	menuLeve2 := d.MenuLevel2
	if menuLevel1 == "" {
		return errors.New("一级菜单不能为空")
	}
	if menuLeve2 == "" {
		return errors.New("菜单不能为空")
	}
	menuLogo := d.MenuLogo
	if menuLogo == "" {
		menuLogo = xicon.GenIcon()
	}
	menu1, err := dao.Menu.GetByName(ctx, menuLevel1)
	if err != nil {
		if err == consts.ErrDataNotFound {
			// 新增一级菜单
			maxSort, err := dao.Menu.Ctx(ctx).Max("sort")
			if err != nil {
				return err
			}
			id, err := dao.Menu.Ctx(ctx).InsertAndGetId(&entity.Menu{
				Pid:    -1,
				Name:   menuLevel1,
				Type:   2,
				Sort:   math.Ceil(maxSort),
				Status: 1,
			})
			if err != nil {
				return err
			}
			menu1 = &entity.Menu{Id: int(id)}
			goto here
		}
		return err
	}
	if menu1.Type != 2 {
		return errors.New("一级菜单必须为分组菜单")
	}
here:
	// 新增二级菜单

	return nil
}
func genController(ctx context.Context, d *bo.GenConf) error {
	pwd := gfile.MainPkgPath()
	line, err := xfile.ReadLine(fmt.Sprint(pwd, "/go.mod"), 1)
	if err != nil {
		return err
	}
	// mod
	mod := gstr.SplitAndTrim(line, " ")[1]
	temp := gfile.GetContents(fmt.Sprint(pwd, "/resource/gen/controller.temp"))
	temp = gstr.Replace(temp, "[mod]", mod)

	// Menu
	caseCamel := gstr.CaseCamel(d.StructName)
	temp = gstr.Replace(temp, "Menu", caseCamel)

	// t1
	t1 := d.T1
	temp = gstr.Replace(temp, "[t1]", t1)

	// group
	group := d.HtmlGroup
	temp = gstr.Replace(temp, "[group]", group)

	// orderBy
	orderBy := d.OrderBy
	if d.OrderBy == "" {
		orderBy = "t1.sort desc,t1.id desc"
	}
	temp = gstr.Replace(temp, "[orderBy]", orderBy)

	// searchFields
	temp = gstr.Replace(temp, "[searchFields]", d.QueryField)

	// fields
	fields := ""
	if len(d.Fields) != 0 {
		for _, i := range d.Fields {
			if i.SearchType == 0 {
				continue
			}
			name := i.Name
			searchType := i.SearchType
			queryName := i.QueryName
			t := "{"
			t += fmt.Sprintf(`Name: "%s"`, name)
			if searchType != 0 {
				t += fmt.Sprintf(", SearchType: %d", searchType)
			}
			if queryName != "" {
				t += fmt.Sprintf(`, QueryName: "%s"`, queryName)
			}
			t += "}, "
			fields += t
		}
	}
	temp = gstr.Replace(temp, "[fields]", fields)
	date := gtime.Now()
	temp = gstr.Replace(temp, "[date]", date.String())

	// file
	filePath := fmt.Sprint(pwd, "/internal/controller/", d.T1, ".go")
	f, err := gfile.Create(filePath)
	if err != nil {
		return err
	}
	if _, err = f.WriteString(temp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func genRouter(ctx context.Context, d *bo.GenConf) error {
	temp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/router.temp"))
	structName := gstr.CaseCamelLower(d.StructName)
	caseCamel := gstr.CaseCamel(structName)
	temp = gstr.Replace(temp, "menu", structName)
	temp = gstr.Replace(temp, "Menu", caseCamel)

	// sys_router
	sysRouterPath := fmt.Sprint(gfile.MainPkgPath(), "/internal/cmd/sys_router.go")
	sysRouter := gfile.GetContents(sysRouterPath)
	if gstr.Contains(sysRouter, temp) {
		return nil
	}
	stat, err := gfile.Stat(sysRouterPath)
	if err != nil {
		return err
	}
	if err := gfile.Truncate(sysRouterPath, int(stat.Size()-2)); err != nil {
		return err
	}
	if err := gfile.PutContentsAppend(sysRouterPath, temp); err != nil {
		return err
	}
	return nil
}
func genHtml(ctx context.Context, c *bo.GenConf) error {
	htmlGroup := c.HtmlGroup
	if htmlGroup == "" {
		return errors.New("html group can't be empty")
	}
	table := c.T1
	if table == "" {
		return errors.New("table can't ben empty")
	}
	filePath := fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", htmlGroup, "/", table, ".html")
	temp := gfile.GetContents(fmt.Sprintf("%s/resource/gen/temp.html", gfile.MainPkgPath()))
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
		if i.SearchType != 0 {
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

func genApi(ctx context.Context, category string, name, pageName string) error {
	if pageName == "" {
		pageName = name
	}
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s/path", name), Method: "GET", Group: category, Desc: fmt.Sprintf("%s页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "GET", Group: category, Desc: fmt.Sprintf("查询%s集合", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/:id", name), Method: "GET", Group: category, Desc: fmt.Sprintf("查询%s详情", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/:id", name), Method: "DELETE", Group: category, Desc: fmt.Sprintf("删除%s", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "POST", Group: category, Desc: fmt.Sprintf("添加%s", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "PUT", Group: category, Desc: fmt.Sprintf("修改%s", pageName), Status: 1},
	}
	for _, i := range array {
		count, err := dao.Api.Ctx(ctx).Count("url = ? and method = ?", i.Url, i.Method)
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
