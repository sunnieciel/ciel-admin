package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/xcmd"
	"ciel-admin/utility/utils/xfile"
	"ciel-admin/utility/utils/xicon"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math"
	"regexp"
	"sort"
	"strings"
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

// genCodeCURD

func GenCodeSetConf(ctx context.Context, d *bo.GenConf, p *gcmd.Parser) int {
	d.T1 = p.GetOpt("t1").String()
	d.T2 = p.GetOpt("t2").String()
	d.T3 = p.GetOpt("t3").String()
	d.T4 = p.GetOpt("t4").String()
	d.T5 = p.GetOpt("t5").String()
	d.T6 = p.GetOpt("t6").String()
	d.QueryField = p.GetOpt("q").String()
	d.OrderBy = p.GetOpt("orderBy").String()
	d.AddBtn = gconv.Int(p.GetOpt("hadd"))
	d.UpdateBtn = gconv.Int(p.GetOpt("hedit"))
	d.DelBtn = gconv.Int(p.GetOpt("hdel"))
	d.ShowStatus = gconv.Int(p.GetOpt("hs"))
	d.PageDesc = p.GetOpt("desc").String()
	d.MenuLogo = p.GetOpt("logo").String()
	d.MenuLevel2 = p.GetOpt("pageName").String()
	d.PageName = p.GetOpt("pageName").String()
	d.HtmlGroup = p.GetOpt("group").String()
	genType := gconv.Int(p.GetOpt("genType"))
	if genType == 0 {
		genType = xcmd.MustScan("生成类型(1 curd 2 静态页面):").Int()
	}
	level1, err := dao.Menu.Ctx(ctx).Array("name", "pid=-1")
	if err != nil {
		panic(err)
	}
	d.MenuLevel1 = xcmd.MustChoose("父级菜单名:", level1).String()
	if d.MenuLevel2 == "" {
		d.MenuLevel2 = xcmd.MustScan("当前页面菜单名:").String()
		d.PageName = d.MenuLevel2 // 设置页面名称
	}
	if d.HtmlGroup == "" {
		d.HtmlGroup = xcmd.MustScan("分组文件夹(项目/resource/template/你输入的分组文件 ):").String()
	}

	d.ApiGroup = p.GetOpt("apiGroup").String()
	if d.ApiGroup == "" {
		d.ApiGroup = xcmd.MustScan("输入API分组:").String()
	}
	return genType
}
func GenCodeGreet(ctx context.Context) {
	g.Log().Notice(ctx, "\n\n███████╗██████╗ ███████╗███████╗    ██╗  ██╗███████╗██╗   ██╗     █████╗ ██████╗ ███╗   ███╗██╗███╗   ██╗\n██╔════╝██╔══██╗██╔════╝██╔════╝    ██║ ██╔╝██╔════╝╚██╗ ██╔╝    ██╔══██╗██╔══██╗████╗ ████║██║████╗  ██║\n█████╗  ██████╔╝█████╗  █████╗      █████╔╝ █████╗   ╚████╔╝     ███████║██║  ██║██╔████╔██║██║██╔██╗ ██║\n██╔══╝  ██╔══██╗██╔══╝  ██╔══╝      ██╔═██╗ ██╔══╝    ╚██╔╝      ██╔══██║██║  ██║██║╚██╔╝██║██║██║╚██╗██║\n██║     ██║  ██║███████╗███████╗    ██║  ██╗███████╗   ██║       ██║  ██║██████╔╝██║ ╚═╝ ██║██║██║ ╚████║\n"+
		"╚═╝     ╚═╝  ╚═╝╚══════╝╚══════╝    ╚═╝  ╚═╝╚══════╝   ╚═╝       ╚═╝  ╚═╝╚═════╝ ╚═╝     ╚═╝╚═╝╚═╝  ╚═══╝\n Welcome to use the build tools of freekey , let's go!                                                                                                        \n")
}
func GenStaticHtmlFile(ctx context.Context, c bo.GenConf) error {
	htmlFilePath := fmt.Sprint("/", c.HtmlGroup, "/", gstr.CaseCamelLower(c.StructName), ".html")
	// gen menu
	if err := genMenu(ctx, c, func(name string) string {
		return fmt.Sprintf("/admin/to/%s", c.StructName)
	}, htmlFilePath); err != nil {
		return err
	}
	temp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.static.html"))
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template", htmlFilePath))
	if err != nil {
		return err
	}
	date := gtime.Now()
	temp = gstr.Replace(temp, "[date]", date.String())
	if _, err = f.WriteString(temp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
func CRUDBefore(ctx context.Context, d *bo.GenConf) error {
	tables, err := g.DB().Tables(ctx)
	if err != nil {
		panic(err)
	}
	if d.T1 == "" {
		d.T1 = xcmd.MustChoose("表名:", gvar.New(tables).Vars()).String()
	}

	if d.T2 != "" { // 当存在关联查询时检查查询字段是否为空，空则要求必须输入
		if d.QueryField == "" {
			d.QueryField = xcmd.MustScan("请输入查询字段：").String()
		}
	}
	glog.Info(ctx, "CRUD base check end.")
	d.StructName = gstr.CaseCamelLower(d.T1[strings.Index(d.T1, "_")+1:])
	return err
}
func CRUDParseFields(ctx context.Context, d *bo.GenConf) {
	fields, err := Fields(ctx, d.T1)
	if err != nil {
		panic(err)
	}
	d.Fields = make([]*bo.GenFiled, 0)

	for _, i := range fields {
		data := &bo.GenFiled{TableField: i}
		if i.Comment != "" {
			findString := makeToJsonStr2(regexp.MustCompile("{.*}").FindString(i.Comment))
			json, err := gjson.DecodeToJson(findString)
			if err != nil {
				panic(fmt.Errorf("解析字段备注时格式不正确!请将格式修改为标准的json字符串。%v %v", findString, err))
			}
			if err = json.Scan(data); err != nil {
				panic(err)
			}
		}
		if data.Label == "" {
			data.Label = i.Name
		}
		if data.FieldType == "" {
			data.FieldType = "text"
		}
		d.Fields = append(d.Fields, data)
	}
	if d.QueryField != "" {
		for {
			yes := xcmd.MustScan("是否添加关联表展示字段？(如果设置了关联查询，请设置相应字段)  1 yes ,0 no:\t")
			if yes.Int() == 0 {
				break
			}
			data := &bo.GenFiled{TableField: &gdb.TableField{Name: xcmd.MustScan("字段名称(eg t2.uname):\t").String()}}
			data.QueryName = xcmd.MustScan("查询字段名称(eg uname):\t").String()
			data.Label = xcmd.MustScan("字段名称:\t").String()
			data.SearchType = xcmd.MustScan("查询类型 0 no,1 =,2 like,3 >, 4 <, 5 >=, 6 <=,7 !=\t").Int()
			data.EditHide = 1 // 编辑时隐藏
			data.AddHide = 1  // 添加时隐藏
			index := xcmd.MustScan("请输入元素位置:\t").Int()
			if index <= 0 || index > len(d.Fields) {
				d.Fields = append(d.Fields, data)
			} else {
				d.Fields = append(d.Fields[:index], append([]*bo.GenFiled{data}, d.Fields[index:]...)...)
			}
		}
	}
}
func GenFile(ctx context.Context, d bo.GenConf) (err error) {
	if d.StructName == "" {
		return fmt.Errorf("结构体名称不能为空")
	}
	if err = genMenu(ctx, d, func(name string) string {
		return fmt.Sprintf("/admin/%s/path", gstr.CaseCamelLower(name))
	}, ""); err != nil {
		return err
	}
	if err = genController(d); err != nil {
		return err
	}
	if err = genRouter(d); err != nil {
		return err
	}
	if err = genHtml(d); err != nil {
		return err
	}
	if err = genApi(ctx, d.HtmlGroup, d.PageName, d.ApiGroup); err != nil {
		return err
	}
	return
}
func genHtml(c bo.GenConf) error {
	if err := genIndex(c); err != nil {
		return err
	}
	if c.AddBtn == 0 {
		if err := genAdd(c); err != nil {
			return err
		}
	}
	if c.UpdateBtn == 0 {
		if err := genEdit(c); err != nil {
			return err
		}
	}
	return nil
}
func genMenu(ctx context.Context, d bo.GenConf, menuPathFuc func(name string) string, filePath string) error {
	menuLevel1 := d.MenuLevel1
	menuLeve2 := d.MenuLevel2
	if menuLevel1 == "" {
		return errors.New("一级菜单不能为空")
	}
	if menuLeve2 == "" {
		return errors.New("菜单不能为空")
	}
	menu1, err := dao.Menu.GetByName(ctx, menuLevel1)
	m1Sort, m2Sort := 0.0, 0.0
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
				Pid:      -1,
				Name:     menuLevel1,
				Type:     2,
				Sort:     m1Sort,
				Status:   1,
				FilePath: filePath,
			})
			if err != nil {
				return err
			}
			g.Log().Debugf(ctx, "新增一级菜单,排序为%v", m1Sort)
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
			m2Sort += m1Sort + 0.1
		} else {
			m2Sort += childrenMaxSort + 0.1
		}
		g.Log().Debugf(ctx, "查询一级菜单，子菜单最大排序为%v", menu1.Sort)
	}
	if menu1.Type != 2 {
		return errors.New("一级菜单必须为分组菜单")
	}
here:
	// 新增二级菜单
	menuLogo := d.MenuLogo
	if menuLogo == "" {
		menuLogo = xicon.GenIcon()
	}

	menuPath := menuPathFuc(d.StructName)
	// count path
	g.Log().Debug(ctx, "检查二级菜单是否存在")
	pathCount, err := dao.Menu.Ctx(ctx).Where("path=?", menuPath).Count()
	if err != nil {
		return err
	}
	if pathCount > 0 {
		g.Log().Warning(ctx, "菜单路径已存在,未执行插入菜单操作")
		return nil
	}
	if _, err = dao.Menu.Ctx(ctx).Insert(&entity.Menu{
		Pid:      menu1.Id,
		Icon:     menuLogo,
		BgImg:    menuLogo,
		Path:     menuPath,
		Sort:     m2Sort,
		Name:     menuLeve2,
		Status:   1,
		Type:     1,
		Desc:     d.PageDesc,
		FilePath: filePath,
	}); err != nil {
		return err
	}
	g.Log().Debugf(ctx, "新增二级菜单,排序为%v", m2Sort)
	return nil
}
func genController(d bo.GenConf) error {
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
	temp = gstr.Replace(temp, "menu", gstr.CaseCamelLower(d.StructName))

	// tables
	tables := fmt.Sprintf(`T1:"%s"`, d.T1)
	if d.T2 != "" {
		tables += fmt.Sprintf(`,T2:"%s"`, d.T2)
	}
	if d.T3 != "" {
		tables += fmt.Sprintf(`,T3:"%s"`, d.T3)
	}
	if d.T4 != "" {
		tables += fmt.Sprintf(`,T4:"%s"`, d.T4)
	}
	if d.T5 != "" {
		tables += fmt.Sprintf(`,T5:"%s"`, d.T5)
	}
	if d.T6 != "" {
		tables += fmt.Sprintf(`,T6:"%s"`, d.T6)
	}
	temp = gstr.Replace(temp, "[tables]", tables)

	// t1
	temp = gstr.Replace(temp, "[t1]", d.T1)

	// group
	group := d.HtmlGroup
	temp = gstr.Replace(temp, "[group]", group)

	// orderBy
	orderBy := d.OrderBy
	if d.OrderBy == "" {
		orderBy = "t1.id desc"
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
			if queryName == "" {
				queryName = i.Name
			}
			t += fmt.Sprintf(`, QueryName: "%s_%s"`, gstr.CaseCamelLower(d.StructName), queryName)
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
func genRouter(d bo.GenConf) error {
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
func genApi(ctx context.Context, name, pageName, group string) error {
	if err := checkGroupOrSave(ctx, group); err != nil {
		return err
	}
	if pageName == "" {
		pageName = name
	}
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s/path", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/add", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s添加页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/edit/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s修改页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/del/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s删除操作", pageName), Status: 1},
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

func checkGroupOrSave(ctx context.Context, group string) error {
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

func genEdit(c bo.GenConf) error {
	structNameLower := gstr.CaseCamelLower(c.StructName)
	editTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.edit.html"))
	pageName := c.PageName
	editTemp = gstr.Replace(editTemp, "[pageName]", pageName)
	editTemp = gstr.Replace(editTemp, "menu", structNameLower)
	tr := ""
	for _, i := range c.Fields {
		if i.EditHide == 1 {
			continue
		}
		switch i.Name {
		case "id", "status", "created_at", "updated_at":
			continue
		}
		label := i.Label
		if label == "" {
			label = i.Name
		}
		readonly := ""
		if i.Disabled == 1 {
			readonly = "readonly"
		}
		switch i.FieldType {
		case "select":
			tr += fmt.Sprintf("<tr><td align='right'>%s</td><td>", label)
			temp := fmt.Sprintf("<select name='%s' %s>", i.Name, readonly)
			temp += fmt.Sprintf(`{{option "%s" .Session.%s_edit.%s}}`, i.Options, structNameLower, i.Name)
			temp += "</select>"
			tr += temp
			tr += "</td>"
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		case "textarea":
			tr += fmt.Sprintf("<tr><td width='160' align='right'>%s</td><td width='auto' align='left'><textarea name='%s' %s>{{.Session.%s_edit.%s}}</textarea></td>",
				label, i.Name, readonly, structNameLower, i.Name,
			)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		default:
			tr += fmt.Sprintf("<tr><td width='160' align='right'>%s</td><td width='auto' align='left'><input name='%s' value='{{.Session.%s_edit.%s}}' %s></td>",
				label, i.Name, structNameLower, i.Name, readonly)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		}
	}
	editTemp = gstr.Replace(editTemp, "[tr]", tr)

	date := gtime.Now()
	editTemp = gstr.Replace(editTemp, "[date]", date.String())
	if c.ShowStatus == 0 {
		editTemp = gstr.Replace(editTemp, "[status]", fmt.Sprintf(`<tr><td align='right'>状态</td> <td> <select name='status'> {{option .Config.options.status .Session.%s_edit.status}}</select></td></tr>`, structNameLower))
	} else {
		editTemp = gstr.Replace(editTemp, "[status]", ``)
	}
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", c.HtmlGroup, "/", structNameLower, "/edit.html"))
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
func genAdd(c bo.GenConf) error {
	addTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.add.html"))
	pageName := c.PageName
	addTemp = gstr.Replace(addTemp, "[pageName]", pageName)
	// menu
	addTemp = gstr.Replace(addTemp, "menu", gstr.CaseCamelLower(c.StructName))
	tr := ""
	for _, i := range c.Fields {
		if i.AddHide == 1 {
			continue
		}
		switch i.Name {
		case "id", "status", "created_at", "updated_at":
			continue
		}
		label := i.Label
		if label == "" {
			label = i.Name
		}
		required := ""
		if i.Required == 1 {
			required = "required"
		}
		switch i.FieldType {
		case "select":
			tr += fmt.Sprintf("<tr><td align='right'>%s</td><td>", label)
			temp := fmt.Sprintf("<select name='%s' required><option value='' >请选择</option>", i.Name)
			temp += fmt.Sprintf(`{{option "%s" ""}}`, i.Options)
			temp += "</select>"
			tr += temp + "</td>"
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		case "textarea":
			tr += fmt.Sprintf("<tr><td width='160' align='right'>%s</td><td width='auto' align='left'><textarea name='%s' %s></textarea></td>",
				label, i.Name, required,
			)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		default:
			tr += fmt.Sprintf("<tr><td width='160' align='right'>%s</td><td width='auto' align='left'><input name='%s'  %s></td>",
				label, i.Name, required,
			)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		}
	}
	addTemp = gstr.Replace(addTemp, "[tr]", tr)

	date := gtime.Now()
	addTemp = gstr.Replace(addTemp, "[date]", date.String())

	// status
	if c.ShowStatus == 0 {
		addTemp = gstr.Replace(addTemp, "[status]", `<tr> <td align='right'>状态</td><td><select name='status'> {{option .Config.options.status ""}} </select></td> </tr>`)
	} else {
		addTemp = gstr.Replace(addTemp, "[status]", "")
	}
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", c.HtmlGroup, "/", gstr.CaseCamelLower(c.StructName), "/add.html"))
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
func genIndex(c bo.GenConf) error {
	indexTemp := gfile.GetContents(fmt.Sprintf("%s/resource/gen/temp.index.html", gfile.MainPkgPath()))
	group := c.HtmlGroup
	structNameLower := gstr.CaseCamelLower(c.StructName)
	if c.AddBtn == 1 {
		indexTemp = gstr.Replace(indexTemp, "[add]", ``)
	} else {
		indexTemp = gstr.Replace(indexTemp, "[add]", `<a class="tag-info mr-3" href="{{.node.Path}}/add?{{toUrlParams .Query}}" > <i class="fa fa-plus" aria-hidden="true"></i></a> `)
	}
	if c.DelBtn == 1 {
		indexTemp = gstr.Replace(indexTemp, "[del]", "")
	} else {
		indexTemp = gstr.Replace(indexTemp, "[del]", `<a href="#"  onclick="if(confirm('确认删除?')){location.href='{{$.node.Path}}/del/{{.id}}?{{toUrlParams $.Query}}'}" class="tag-danger"><i class="fa fa-trash"></i></a>
`)
	}
	if c.UpdateBtn == 1 {
		indexTemp = gstr.Replace(indexTemp, "[edit]", "")
	} else {
		indexTemp = gstr.Replace(indexTemp, "[edit]", `<a href="#" onclick=" location.href='{{$.node.Path}}/edit/{{.id}}?{{toUrlParams $.Query}}'" class="tag-info"><i class="fa fa-wrench" aria-hidden="true"></i></a>`)
	}
	// search
	search := ``
	for _, i := range c.Fields {
		if i.SearchType == 0 {
			continue
		}
		label := i.Label
		if label == "" {
			label = i.Name
		}

		queryName := i.QueryName
		if queryName == "" {
			queryName = i.Name
		}
		switch i.FieldType {
		case "select":
			if i.Options == "" {
				return fmt.Errorf("%s选项不能为空:eg '1:正常:tag-info,2:禁用:tag-danger'", i.Name)
			}
			search += fmt.Sprintf("<label class='input'>%s <select type='text' name='%s_%s' value='{{.Session.%s_%s}}' onchange='this.form.submit()'>", label, structNameLower, queryName, structNameLower, queryName)
			search += "<option value='' class='tag-info'>请选择</option>"
			search += fmt.Sprintf(`{{option "%s" .Query.%s_%s}}`, i.Options, structNameLower, queryName)
			search += "</select></label>"
		default:
			search += fmt.Sprintf(`<label class="input">%s <input type="text" name="%s_%s" value="{{.Session.%s_%s}}" onkeydown="if(event.keyCode===13)this.form.submit()"></label>`, label, structNameLower, queryName, structNameLower, queryName)
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[search]", search)
	// table td
	th := ""
	td := ""
	for _, i := range c.Fields {
		if i.Hide == 1 {
			continue
		}
		name := i.QueryName
		if name == "" {
			name = i.Name
		}
		label := i.Label
		if label == "" {
			label = i.Name
		}
		switch name {
		case "id", "created_at", "updated_at", "status":
			continue
		}
		th += fmt.Sprintf("<th>%s</th>", label)
		switch i.FieldType {
		case "select":
			temp := fmt.Sprintf("<td data-label='%s'>", label)
			temp += fmt.Sprintf(`{{chooseSpan "%s" .%s}}`, i.Options, i.Name)
			temp += "</td>"
			td += temp
		case "img":
			td += fmt.Sprintf(`<td data-label='%s'>{{img .%s}}</td>`, label, name)
		default:
			td += fmt.Sprintf("<td data-label='%s'>{{.%s}}</td>", label, name)
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[th]", th)
	indexTemp = gstr.Replace(indexTemp, "[td]", td)

	date := gtime.Now()
	indexTemp = gstr.Replace(indexTemp, "[date]", date.String())

	// status
	if c.ShowStatus == 0 {
		indexTemp = gstr.Replace(indexTemp, "[status_label]", "<th>状态</th>")
		indexTemp = gstr.Replace(indexTemp, "[status]", "<td>{{chooseSpan $.Config.options.status .status}}</td>")
	} else {
		indexTemp = gstr.Replace(indexTemp, "[status_label]", "")
		indexTemp = gstr.Replace(indexTemp, "[status]", "")
	}
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

func makeToJsonStr(str string) string {
	// 替换所有空格
	replace, _ := gregex.Replace(`\s`, []byte(""), []byte(str))
	// 处理key未加""的内容字段
	replace, _ = gregex.Replace(`label|"label"`, []byte(`"label"`), replace)
	replace, _ = gregex.Replace(`fieldType|"fieldType"`, []byte(`"fieldType"`), replace)
	replace, _ = gregex.Replace(`searchType|"searchType"`, []byte(`"searchType"`), replace)
	replace, _ = gregex.Replace(`editHide|"editHide"`, []byte(`"editHide"`), replace)
	replace, _ = gregex.Replace(`addHide|"addHide"`, []byte(`"addHide"`), replace)
	replace, _ = gregex.Replace(`hide|"hide"`, []byte(`"hide"`), replace)
	replace, _ = gregex.Replace(`disabled|"disabled"`, []byte(`"disabled"`), replace)
	replace, _ = gregex.Replace(`required|"required"`, []byte(`"required"`), replace)
	replace, _ = gregex.Replace(`comment|"comment"`, []byte(`"comment"`), replace)
	replace, _ = gregex.Replace(`options|"options"`, []byte(`"options"`), replace)
	// 处理值未加个""的字段
	doAdd := func(temp string) []string {
		defer func() {
			if r := recover(); r != nil {
				panic(r)
			}
		}()
		strs := make([]string, 0)
		for _, i := range gstr.Split(temp, ",") {
			i = gstr.TrimAll(i)
			if i == "" {
				continue
			}
			begin := gstr.Split(i, ":")[0]
			end := gstr.Split(i, ":")[1]
			end = gstr.Replace(end, `"`, "")
			strs = append(strs, fmt.Sprintf(`%s:"%s"`, begin, end))
		}
		return strs
	}
	temp := string(replace)
	temp = temp[1 : len(temp)-1]
	strs := make([]string, 0)
	if !gstr.Contains(temp, `"options":`) {
		strs = append(strs, doAdd(temp)...)
	} else {
		t := gstr.Split(temp, `"options":`)
		strs = append(strs, doAdd(t[0])...)
		t[1] = gstr.Replace(t[1], `"`, "")
		t[1] = fmt.Sprintf(`"%s"`, t[1])
		strs = append(strs, fmt.Sprintf(`"options":%s`, t[1]))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(strs, ","))
}
func makeToJsonStr2(str string) string {
	re := regexp.MustCompile(`\s*"?(options)"?\s*:\s*"?((?:[^,]*?:[^,]*?:[^,]*?,?)*)"?\s*([,}])|\s*"?(\w+)"?\s*:\s*"?(.*?)"?\s*([,}])`)
	return re.ReplaceAllString(str, `"$1$4":"$2$5"$3$6`)
}
