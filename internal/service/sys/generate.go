package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
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
func GenFile(ctx context.Context, d bo.GenConf) (err error) {
	if d.StructName == "" {
		return fmt.Errorf("结构体名称不能为空")
	}
	//// gen menu
	if err = genMenu(ctx, d, func(name string) string {
		return fmt.Sprintf("/%s/path", gstr.CaseCamelLower(name))
	}, ""); err != nil {
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
func GenStaticHtmlFile(ctx context.Context, c bo.GenConf) error {
	if c.HtmlGroup == "" {
		return fmt.Errorf("html分组不能为空")
	}
	if c.StructName == "" {
		return fmt.Errorf("结构体名称不能为空")
	}
	htmlFilePath := fmt.Sprint("/", c.HtmlGroup, "/", gstr.CaseCamelLower(c.StructName), ".html")
	// gen menu
	if err := genMenu(ctx, c, func(name string) string {
		return fmt.Sprintf("/to/%s", c.StructName)
	}, htmlFilePath); err != nil {
		return err
	}
	// gen file
	return genStaticHtmlFile(ctx, c, htmlFilePath)
}

func genStaticHtmlFile(ctx context.Context, c bo.GenConf, filePath string) error {
	temp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.static.html"))

	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template", filePath))
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

// filePath: Need to use when generating static pages,not when curd.
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
func genController(ctx context.Context, d bo.GenConf) error {
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
				t += fmt.Sprintf(`, QueryName: "%s_%s"`, gstr.CaseCamelLower(d.StructName), queryName)
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
func genRouter(ctx context.Context, d bo.GenConf) error {
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
func genHtml(ctx context.Context, c bo.GenConf) error {
	if err := genIndex(ctx, c); err != nil {
		return err
	}
	if err := genAdd(ctx, c); err != nil {
		return err
	}
	if err := genEdit(ctx, c); err != nil {
		return err
	}
	return nil
}
func genEdit(ctx context.Context, c bo.GenConf) error {
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
		if i.EditDisabled == 1 {
			readonly = "readonly"
		}
		switch i.FieldType {
		case "select":
			tr += fmt.Sprintf("<tr><td align='right'>%s</td><td>", label)
			temp := fmt.Sprintf("<select name='%s' %s>", i.Name, readonly)
			for _, j := range i.Options {
				temp += fmt.Sprintf("<option value='%s' {{if eq .Session.%s_edit.%s %s}} selected {{end}} class='%s'>%v</option>",
					j.Value, structNameLower, i.Name, j.Value, j.Type, j.Label)
			}
			temp += "</select>"
			tr += temp
			tr += "</td>"
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "</tr>"
		case "textarea":
			tr += fmt.Sprintf("<tr><td width='160' align='right'>%s</td><td width='auto' align='left'><textarea name='%s' %s>{{.Session.%s_edit.%s}}</textarea></td>",
				label, i.Name, structNameLower, readonly, i.Name,
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
	// status
	if c.ShowStatus == 0 {
		editTemp = gstr.Replace(editTemp, "[status]", fmt.Sprintf(`<tr> 
<td align='right'>状态</td>
<td>
<select name='status'> 
<option value='1' {{if eq .Session.%s_edit.status 1}} selected {{end}}  class="tag-info">开启</option> 
<option value='2' {{if eq .Session.%s_edit.status 2}} selected {{end}} class="tag-danger">关闭</option> 
</select>
</td>
</tr>`, structNameLower, structNameLower))
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
func genAdd(ctx context.Context, c bo.GenConf) error {
	addTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.add.html"))
	pageName := c.PageName
	addTemp = gstr.Replace(addTemp, "[pageName]", pageName)
	// menu
	addTemp = gstr.Replace(addTemp, "menu", gstr.CaseCamelLower(c.StructName))
	tr := ""
	for _, i := range c.Fields {
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
			tr += fmt.Sprintf("                          <tr>  <td align='right'>%s</td>\n                            <td>", label)
			temp := fmt.Sprintf("                                <select name='%s' %s>\n", i.Name, required)
			for index, j := range i.Options {
				if index == 0 {
					temp += fmt.Sprintf("                                    <option value='%v' selected class='%s'>%v</option>\n", j.Value, j.Type, j.Label)
				} else {
					temp += fmt.Sprintf("                                    <option value='%v'  class='%s'>%v</option>\n", j.Value, j.Type, j.Label)
				}
			}
			temp += "                                </select>\n"
			tr += temp + "</td>"
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "\n </tr>"
		case "textarea":
			tr += fmt.Sprintf("                        <tr>\n                            <td width='160' align='right'>%s</td>\n                            <td width='auto' align='left'><textarea name='%s' %s></textarea></td>",
				label, i.Name, required,
			)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "\n                        </tr>"
		default:
			tr += fmt.Sprintf("                        <tr>\n                            <td width='160' align='right'>%s</td>\n                            <td width='auto' align='left'><input name='%s'  %s></td>",
				label, i.Name, required,
			)
			if i.Comment != "" {
				tr += fmt.Sprintf("<td><span class='tag-info'>%s</span></td>", i.Comment)
			}
			tr += "\n                        </tr>"
		}
	}
	addTemp = gstr.Replace(addTemp, "[tr]", tr)

	date := gtime.Now()
	addTemp = gstr.Replace(addTemp, "[date]", date.String())

	// status
	if c.ShowStatus == 0 {
		addTemp = gstr.Replace(addTemp, "[status]", `<tr> <td align='right'>状态</td><td><select name='status'> <option value='1' selected class="tag-info">开启</option> <option value='2' class="tag-danger">关闭</option> </select></td> </tr>`)
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
func genIndex(ctx context.Context, c bo.GenConf) error {
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
			search += fmt.Sprintf("<label class='input'>%s <select type='text' name='%s_%s' value='{{.Session.%s_%s}}' onchange='this.form.submit()'>", label, structNameLower, queryName, structNameLower, queryName)
			search += "<option value='' class='tag-info'>请选择</option>"
			for _, j := range i.Options {
				search += fmt.Sprintf("<option value='%s' {{if eq .Query.%s_%s %s}} selected {{end}} class='%s'>%s</option>", j.Value, structNameLower, queryName, j.Value, j.Type, j.Label)
			}
			search += "</select></label>"
		default:
			search += fmt.Sprintf(`<label class="input">%s <input type="text" name="%s_%s" value="{{.Session.%s_%s}}" onkeydown="if(event.keyCode===13)this.form.submit()"></label>`,
				label, structNameLower, queryName, structNameLower, queryName,
			)
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[search]", search)
	// table td
	th := ""
	td := ""
	for _, i := range c.Fields {
		if i.NotShow == 1 {
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
		th += fmt.Sprintf("<th>%s</th>\n                        ", label)
		switch i.FieldType {
		case "select":
			temp := "<td>"
			for index, j := range i.Options {
				if index == 0 {
					temp += fmt.Sprintf(`{{if eq .%s "%s"}}<span class="%s">%s</span>`, i.Name, j.Value, j.Type, j.Label)
				} else {
					temp += fmt.Sprintf(`{{else if eq .%s "%s"}}<span class="%s">%s</span>`, i.Name, j.Value, j.Type, j.Label)
				}
			}
			temp += "{{end}}</td>"
			td += temp
		case "img":
			td += fmt.Sprintf(`<td>{{if hasPrefix .%s "http"}} <a href='{{.%s}}' target='_blank'> <img class='s-icon' src='{{.%s}}' alt='not fond'> </a> {{else if ne .%s ""}} <a href='{{$.Config.server.imgPrefix}}{{.%s}}' target='_blank'> <img class='s-icon' src='{{$.Config.server.imgPrefix}}{{.%s}}' alt='not fond'> </a> {{else}} <span v-else class='tag-normal'>暂无图片</span>{{end}}</td>`,
				name, name, name, name, name, name,
			)
		default:
			td += fmt.Sprintf("<td>{{.%s}}</td>\n                        ", name)
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[th]", th)
	indexTemp = gstr.Replace(indexTemp, "[td]", td)

	date := gtime.Now()
	indexTemp = gstr.Replace(indexTemp, "[date]", date.String())

	// status
	if c.ShowStatus == 0 {
		indexTemp = gstr.Replace(indexTemp, "[status_label]", "<th>状态</th>")
		indexTemp = gstr.Replace(indexTemp, "[status]", "<td>{{if eq .status 1}}<span class=\"tag-info\">正常</span>{{else}}<span class=\"tag-danger\">关闭</span>{{end}}</td>")
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
func genApi(ctx context.Context, category string, name, pageName string) error {
	if pageName == "" {
		pageName = name
	}
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s/path", name), Method: "1", Group: category, Desc: fmt.Sprintf("%s页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/add", name), Method: "1", Group: category, Desc: fmt.Sprintf("%s添加页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/edit/:id", name), Method: "1", Group: category, Desc: fmt.Sprintf("%s修改页面", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s/path/del/:id", name), Method: "4", Group: category, Desc: fmt.Sprintf("%s删除操作", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "2", Group: category, Desc: fmt.Sprintf("添加%s", pageName), Status: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "3", Group: category, Desc: fmt.Sprintf("修改%s", pageName), Status: 1},
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
