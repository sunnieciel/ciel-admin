package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/view"
	"ciel-admin/utility/utils/xfile"
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
	"strings"
)

var (
	Gen = lGen{}
)

type lGen struct{}

func (l lGen) Tables(ctx context.Context) (string, error) {
	var (
		str []string
	)
	tables, err := g.DB().Tables(ctx)
	for index, i := range tables {
		str = append(str, fmt.Sprintf("%s:%s:%s", i, i, view.SwitchTagClass(index)))
	}
	return strings.Join(str, ","), err
}

// Gen 生成代码
func (l lGen) Gen(ctx context.Context, table string, group string, menu string, prefix string, apiGroup string, htmlGroup string) error {
	// 结构体名称
	structName := gstr.CaseCamelLower(gstr.Replace(table, prefix, ""))
	// 表所有的字段
	fields, err := l.TableFields(ctx, table)
	if err != nil {
		return err
	}
	// 生成菜单
	if err = l.doGenMenu(ctx, group, menu, table, prefix); err != nil {
		return err
	}
	// 生成api
	if err = l.genApi(ctx, structName, menu, apiGroup); err != nil {
		return err
	}
	// 生成控制层
	if err = l.genController(table, htmlGroup, structName); err != nil {
		return err
	}
	// 生成 router
	if err = l.genRouter(structName); err != nil {
		return err
	}
	// 生成 html index
	if err = l.genIndex(htmlGroup, structName, fields); err != nil {
		return err
	}
	// 生成 html add
	if err = l.genAdd(htmlGroup, structName, menu, fields); err != nil {
		return err
	}
	// 生成 html edit
	if err = l.genEdit(htmlGroup, structName, menu, fields); err != nil {
		return err
	}
	return nil
}

func (lGen) genEdit(htmlGroup, structName, menu string, fields []*gdb.TableField) error {
	structNameLower := gstr.CaseCamelLower(structName)
	editTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.edit.html"))
	// pageName
	pageName := menu
	editTemp = gstr.Replace(editTemp, "[pageName]", pageName)
	// menu
	editTemp = gstr.Replace(editTemp, "menu", structNameLower)

	// tr
	tr := ""
	for index, i := range fields {
		if index == 0 {
			switch strings.ToLower(i.Name) {
			case "id":
				tr += fmt.Sprintf(`{{editTrReadonly "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`{{editTr "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			}
		} else {
			switch i.Name {
			case "status":
				tr += fmt.Sprintf(`                        {{editTrOptions "%s" "%s" .Config.options.status .Form.%s}}
`, i.Name, i.Name, i.Name)
			case "updated_at", "created_at":
				tr += fmt.Sprintf(`                        {{editTrReadonly "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`                        {{editTr "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			}
		}
	}
	editTemp = gstr.Replace(editTemp, "[tr]", tr)
	// date
	date := gtime.Now()
	editTemp = gstr.Replace(editTemp, "[date]", date.String())
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", htmlGroup, "/", structNameLower, "/edit.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(editTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (lGen) genAdd(htmlGroup, structName, pageName string, fields []*gdb.TableField) error {
	addTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.add.html"))
	addTemp = gstr.Replace(addTemp, "[pageName]", pageName)
	// menu
	addTemp = gstr.Replace(addTemp, "menu", gstr.CaseCamelLower(structName))

	// tr
	tr := ""
	for index, i := range fields {
		if index == 0 {
			switch strings.ToLower(i.Name) {
			case "id", "created_at", "updated_at":
				continue
			default:
				tr += fmt.Sprintf(`{{editTr "%s" "%s" ""}}
`, i.Name, i.Name)
			}
		} else {
			switch strings.ToLower(i.Name) {
			case "created_at", "updated_at", "id":
				continue
			case "status":
				tr += fmt.Sprintf(`                        {{editTrOptions "%s" "%s" .Config.options.status 1}}
`, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`                        {{editTr "%s" "%s" ""}}
`, i.Name, i.Name)
			}
		}
	}
	addTemp = gstr.Replace(addTemp, "[tr]", tr)
	// date
	date := gtime.Now()
	addTemp = gstr.Replace(addTemp, "[date]", date.String())

	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", htmlGroup, "/", gstr.CaseCamelLower(structName), "/add.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(addTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (lGen) genIndex(htmlGroup, structName string, fields []*gdb.TableField) error {
	indexTemp := gfile.GetContents(fmt.Sprintf("%s/resource/gen/temp.index.html", gfile.MainPkgPath()))
	group := htmlGroup
	structNameLower := gstr.CaseCamelLower(structName)
	// Menu
	caseCamel := gstr.CaseCamel(structName)
	indexTemp = gstr.Replace(indexTemp, "Menu", caseCamel)
	indexTemp = gstr.Replace(indexTemp, "menu", gstr.CaseCamelLower(structName))
	// th
	arr := make([]string, 0)
	for _, i := range fields {
		arr = append(arr, strings.ToUpper(i.Name))
	}
	arr = append(arr, "OPERATION")
	th := strings.Join(arr, ",")
	indexTemp = gstr.Replace(indexTemp, "[th]", th)
	// td
	td := ""
	for index, i := range fields {
		if index == 0 { // 如果是第一个
			td += fmt.Sprintf(`{{td "%s" .%s}}
`, i.Name, i.Name)
		} else {
			switch strings.ToLower(i.Name) {
			case "status":
				td += fmt.Sprintf(`                        {{tdChoose "%s" $.Config.options.status .%s}}
`, i.Name, i.Name)
			default:
				td += fmt.Sprintf(`                        {{td "%s" .%s}}
`, i.Name, i.Name)
			}
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[td]", td)

	// date
	date := gtime.Now()
	indexTemp = gstr.Replace(indexTemp, "[date]", date.String())
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", group, "/", structNameLower, "/index.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(indexTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (lGen) genRouter(name string) error {
	temp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/router.temp"))
	structName := gstr.CaseCamelLower(name)
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

func (l lGen) genController(table string, htmlGroup string, structName string) error {
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
	caseCamel := gstr.CaseCamel(structName)
	temp = gstr.Replace(temp, "Menu", caseCamel)
	temp = gstr.Replace(temp, "menu", gstr.CaseCamelLower(structName))

	// group
	temp = gstr.Replace(temp, "[group]", htmlGroup)

	// table
	temp = gstr.Replace(temp, "[table]", table)

	// htmlGroup
	temp = gstr.Replace(temp, "[htmlGroup]", htmlGroup)

	// date
	date := gtime.Now()
	temp = gstr.Replace(temp, "[date]", date.String())

	// file
	filePath := fmt.Sprint(pwd, "/internal/controller/", table, ".go")
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

func (l lGen) genApi(ctx context.Context, name, pageName, group string) error {
	// 检查在字典表中是否存在
	if err := l.checkGroupOrSave(ctx, group); err != nil {
		return err
	}
	if pageName == "" {
		pageName = name
	}
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/add", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s添加页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/edit/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s修改页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/del/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s删除操作", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "2", Group: group, Desc: fmt.Sprintf("添加%s", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "2", Group: group, Desc: fmt.Sprintf("修改%s", pageName), Status: 1},
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

func (lGen) doGenMenu(ctx context.Context, group, menu, table, prefix string) error {
	var (
		m1Sort, m2Sort = 0.0, 0.0
	)
	menu1, err := dao.Menu.GetByName(ctx, group)
	if err != nil {
		if err == consts.ErrDataNotFound {
			g.Log().Debug(ctx, "一级菜单不存在")
			// 新增一级菜单
			maxSort, err := dao.Menu.Ctx(ctx).Max("sort")
			if err != nil {
				return err
			}
			m1Sort = math.Ceil(maxSort)
			m2Sort = m1Sort + 0.1
			id, err := dao.Menu.Ctx(ctx).InsertAndGetId(&entity.Menu{
				Pid:    -1,
				Name:   group,
				Type:   2,
				Sort:   m1Sort,
				Status: 1,
			})
			if err != nil {
				return err
			}
			g.Log().Infof(ctx, "新增一级菜单,排序为%v", m1Sort)
			menu1 = &entity.Menu{Id: int(id)}
			goto here
		}
		return err
	} else {
		// select max sort from menu1'children
		childrenMaxSort, err := dao.Menu.Ctx(ctx).Where("pid=?", menu1.Id).Max("sort")
		if err != nil {
			return err
		}
		if childrenMaxSort == 0 {
			m2Sort += menu1.Sort + 0.1
		} else {
			m2Sort += childrenMaxSort + 0.1
		}
		g.Log().Infof(ctx, "查询一级菜单，子菜单最大排序为%v", menu1.Sort)
	}
	if menu1.Type != 2 {
		return errors.New("一级菜单必须为分组菜单")
	}
here:
	// 查看菜单是否存在
	count, err := dao.Menu.Ctx(ctx).Count("name", menu)
	if count != 0 {
		g.Log().Warningf(ctx, "%s 菜单已存在，就不创建啦", menu)
		return nil
	}
	// 新增二级菜单
	menuPath := fmt.Sprintf("/admin/%s", gstr.CaseCamelLower(gstr.Replace(table, prefix, "")))
	// count path
	g.Log().Debug(ctx, "检查二级菜单是否存在")
	pathCount, err := dao.Menu.Ctx(ctx).Where("path", menuPath).Count()
	if err != nil {
		return err
	}
	if pathCount > 0 {
		g.Log().Warning(ctx, "菜单路径已存在,未执行插入菜单操作")
		return nil
	}
	//menuLogo := xicon.GenIcon()
	if _, err = dao.Menu.Ctx(ctx).Insert(&entity.Menu{
		Pid: menu1.Id,
		//Icon:   menuLogo,
		//BgImg:  menuLogo,
		Path:   menuPath,
		Sort:   m2Sort,
		Name:   menu,
		Status: 1,
		Type:   1,
	}); err != nil {
		return err
	}
	g.Log().Debugf(ctx, "新增二级菜单,排序为%v", m2Sort)
	return nil
}

func (l lGen) checkGroupOrSave(ctx context.Context, group string) error {
	d, err := dao.Dict.GetByKey(ctx, "api_group")
	if err != nil {
		return err
	}
	for _, i := range gstr.Split(d.V, "\n") {
		i = gstr.TrimAll(i)
		if i == group {
			return nil
		}
	}
	d.V += fmt.Sprint("\n", group)
	if _, err = dao.Dict.Ctx(ctx).Save(d); err != nil {
		return err
	}
	g.Log().Warningf(ctx, "%s 分组在词典表中不存在，已添加.", group)
	return nil
}

func (l lGen) TableFields(ctx context.Context, table string) ([]*gdb.TableField, error) {
	var (
		arr = make([]*gdb.TableField, 0)
	)
	fields, err := g.DB().TableFields(ctx, table)
	if err != nil {
		return nil, err
	}
	for _, v := range fields {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].Index < arr[j].Index })
	return arr, nil
}
