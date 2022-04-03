package service

import (
	"ciel-admin/internal/model/bo"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"sort"
	"strings"
)

type gen struct{ ModName string }

var sGen = newGen()

func newGen() *gen {
	ctx := context.Background()
	get, err := g.Cfg().Get(ctx, "server.modName")
	if err != nil {
		panic(err)
	}
	return &gen{ModName: get.String()}
}
func Gen() *gen {
	return sGen
}
func (s *gen) Fields(ctx context.Context, tableName string) (map[string]*gdb.TableField, error) {
	fields, err := g.DB().Ctx(ctx).Model(tableName).TableFields(tableName)
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (s gen) Tables(ctx context.Context) ([]string, error) {
	return g.DB().Tables(ctx)
}

func (s *gen) GenCode(ctx context.Context, b *bo.GenCodeInfo) error {
	// gen controller
	template, err := s.FileFactory(b, 0)
	if err != nil {
		return err
	}
	err = s.SaveFile(b.Table, template, "", 0)

	// gen cmd
	template, err = s.FileFactory(b, 1)
	if err != nil {
		return err
	}
	err = s.SaveFile(b.Table, template, "", 1)

	// gen html
	template, err = s.FileFactory(b, 2)
	if err != nil {
		return err
	}
	err = s.SaveFile(b.Table, template, b.Category, 2)
	return nil
}

// FileFactory  t 0:controller 1:cmd 2:html
func (s gen) FileFactory(b *bo.GenCodeInfo, t int) (string, error) {
	switch t {
	case 0: // controller
		return s.MakeControllerStr(b)
	case 1: // cmd
		return s.MakeCmdStr(b)
	case 2: // html
		return s.MakeHtmlStr(b)
	}
	return "", nil

}

func (s *gen) MakeControllerStr(b *bo.GenCodeInfo) (string, error) {
	path := fmt.Sprintf("%s/manifest/gen_code_template/controller.text", gfile.MainPkgPath())
	template := gfile.GetContents(path)
	// set modName
	template = strings.ReplaceAll(template, "ciel-admin", s.ModName)
	// set category
	template = strings.ReplaceAll(template, "category", gstr.CaseCamelLower(b.Category))
	// set html file name
	template = strings.ReplaceAll(template, "fileName", b.Table)
	// set structName
	template = strings.ReplaceAll(template, "menu", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "Menu", gstr.CaseCamel(b.StructName))
	// set search config
	temps := ""
	for _, i := range b.Fields {
		switch i.SearchType {
		case "like":
			temps += fmt.Sprintf(`{Field: "%s", Like: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "=":
			temps += fmt.Sprintf(`{Field: "%s",QueryField: "%s"},`, i.Name, i.QueryField)
		case ">":
			temps += fmt.Sprintf(`{Field: "%s",GT: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "<":
			temps += fmt.Sprintf(`{Field: "%s",LT: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case ">=":
			temps += fmt.Sprintf(`{Field: "%s",GTE: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "<=":
			temps += fmt.Sprintf(`{Field: "%s",LTE: true,QueryField: "%s"},`, i.Name, i.QueryField)
		case "in":
			temps += fmt.Sprintf(`{Field: "%s",In: true,QueryField: "%s"},`, i.Name, i.QueryField)
		}
	}
	if temps != "" {
		template = strings.ReplaceAll(template, `{Field: "id"},`, temps)
	}
	// set t1 t2 t3 t4 t5 t6
	manyTable := fmt.Sprintf(` T1:"%s",`, b.Table)
	if b.T2 != "" {
		manyTable += fmt.Sprintf(` T2:"%s",`, b.T2)
	}
	if b.T3 != "" {
		manyTable += fmt.Sprintf(` T3:"%s",`, b.T3)
	}
	if b.T4 != "" {
		manyTable += fmt.Sprintf(` T4:"%s",`, b.T4)
	}
	if b.T5 != "" {
		manyTable += fmt.Sprintf(` T5:"%s",`, b.T5)
	}
	if b.T6 != "" {
		manyTable += fmt.Sprintf(` T6:"%s",`, b.T6)
	}
	if b.OrderBy != "" {
		manyTable += fmt.Sprintf(`OrderBy: "%s",`, b.OrderBy)
	}
	if b.QueryField != "" {
		manyTable += fmt.Sprintf(`SearchFields: "%s",`, b.QueryField)
	}
	template = strings.ReplaceAll(template, "[Tables]", manyTable)
	return template, nil
}
func (s *gen) MakeCmdStr(b *bo.GenCodeInfo) (string, error) {
	path := fmt.Sprintf("%s/manifest/gen_code_template/cmd.text", gfile.MainPkgPath())
	template := gfile.GetContents(path)
	// set structName
	template = strings.ReplaceAll(template, "menu", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "Menu", gstr.CaseCamel(b.StructName))
	return template, nil
}
func (s *gen) MakeHtmlStr(b *bo.GenCodeInfo) (string, error) {
	path := fmt.Sprintf("%s/manifest/gen_code_template/menu.html", gfile.MainPkgPath())
	template := gfile.GetContents(path)
	// set [desc]
	template = strings.ReplaceAll(template, "[desc]", b.Desc)
	// set [menu]
	template = strings.ReplaceAll(template, "[menu]", gstr.CaseCamelLower(b.StructName))
	template = strings.ReplaceAll(template, "[Menu]", gstr.CaseCamel(b.StructName))

	// set tr and td
	var (
		tr, td string
	)

	sort.Slice(b.Fields, func(i, j int) bool {
		return b.Fields[i].Sort < b.Fields[j].Sort
	})
	for _, i := range b.Fields {
		width := 80
		switch i.Name {
		case "id":
			width = 30
		case "created_at", "updated_at":
			width = 120
		default:

		}

		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}

		tr += fmt.Sprintf("                  <th width='%d'>%s</th>\n", width, title)
		if i.Name == "status" {
			td += fmt.Sprintf("                     <td> <el-tag v-if='item.status==1'>启用</el-tag> <el-tag type='danger' v-if='item.status==2'>禁用</el-tag> </td> \n ")
		} else {

			td += fmt.Sprintf("                    <td v-text='item.%s'>\n", itemField)
		}
	}
	template = strings.ReplaceAll(template, `[<th width="30">ID</th>]`, tr)
	template = strings.ReplaceAll(template, `[<td v-text="item.id">]`, td)

	// set search
	search := ""
	for _, i := range b.Fields {
		if i.SearchType == "" {
			continue
		}
		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}
		if i.Name == "status" {
			search += fmt.Sprintf(`
                        <label class="input">状态
                            <select name="status" v-model="search.status" @change="getList()">
                                <option value="">请选择</option>
                                <option value="1">启用</option>
                                <option value="2">禁用</option>
                            </select>
                        </label>
`)
		} else {
			search += fmt.Sprintf(" <label class='input'>%s<input type='text'  v-model='search.%s' autocomplete='off' @keyup.enter='getList()'></label> \n", title, itemField)
		}
	}
	template = strings.ReplaceAll(template, `[<label class="input">PID<input type="text" name="pid" v-model="search.pid" autocomplete="off" @keyup.enter="getList()"></label>]`, search)

	// set details
	details := ""
	for _, i := range b.Fields {
		disabled := ""
		switch i.DetailsType {
		case "no-show":
			continue
		case "disabled":
			disabled = "disabled"
		}
		itemField := i.QueryField
		if itemField == "" {
			itemField = i.Name
		}
		title := i.Title
		if title == "" {
			title = itemField
		}
		if i.Name == "status" {
			details += fmt.Sprintf(`                <li>
                    <label class="input">状态<select v-model="details.status">
                        <option label="正常" :value="1"></option>
                        <option label="关闭" :value="2"></option>
                    </select></label>
                </li>`)
		} else {
			details += fmt.Sprintf("                <li><label class='input'>%s<input v-model='details.%s' %s></label></li>\n", title, itemField, disabled)
		}
	}
	template = strings.ReplaceAll(template, `<li><label class="input">名称<input v-model="details.name"></label></li>`, details)
	return template, nil
}

func (s *gen) SaveFile(fileName string, fileStr string, category string, t int) error {
	switch t {
	case 0: // controller
		path := fmt.Sprintf("%s/internal/controller/%s.go", gfile.MainPkgPath(), fileName)
		return gfile.PutContents(path, fileStr)
	case 1: // cmd
		path := fmt.Sprintf("%s/internal/cmd/sys_router.go", gfile.MainPkgPath())
		stat, _ := gfile.Stat(path)
		if err := gfile.Truncate(path, (int)(stat.Size()-2)); err != nil {
			return err
		}
		return gfile.PutContentsAppend(path, fileStr)
	case 2: // html
		path := fmt.Sprintf("%s/resource/template/%s/%s.html", gfile.MainPkgPath(), category, fileName)
		return gfile.PutContents(path, fileStr)
	}
	return nil
}
