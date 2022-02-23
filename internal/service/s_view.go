package service

import (
	"ciel-begin/manifest/config"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
	"strings"
)

// ---view-------------------------------------------------------------------------------
type view struct{}

func View() *view { return &view{} }
func (s *view) BgColor(value interface{}, color string) string {
	if color != "" {
		return color
	}

	switch gconv.Int(value) % 19 {
	case 1:
		return "#cfc"
	case 2:
		return "#ffcdd2"
	case 3:
		return "#ccf"
	case 4:
		return "#e1bee7"
	case 5:
		return "#ede7f6"
	case -1:
	case 6:
		return "#c5cae9"
	case 7:
		return "#bbdefb"
	case 8:
		return "#81d4fa"
	case 9:
		return "#80deea"
	case 10:
		return "#b2dfdb"
	case 11:
		return "#c8e6c9"
	case 12:
		return "#dcedc8"
	case 13:
		return "#f0f4c3"
	case 14:
		return "#b9f6ca"
	case 15:
		return "#ccff90"
	case 16:
		return "#f4ff81"
	case 17:
		return "#ffccbc"
	case 18:
		return "#bcaaa4"
	case 19:
		return "#e0e0e0"
	case 20:
	default:
		return "#cfd8dc"
	}
	return ""
}
func (s *view) SearchValue(query g.Map, field string) interface{} {
	return query[field]
}
func (s *view) FormLi(field *config.Field, query ...g.Map) string {
	title := field.Title
	if title == "" {
		title = field.Field
	}
	choseField := s.choseField(field)
	name := choseField
	t := field.Type
	if t == "" {
		t = "text"
	}
	hidden := ""
	if field.EditHidden {
		hidden = "none"
	}
	step := field.Step
	value := interface{}(nil)
	if len(query) > 0 {
		value = query[0][choseField]
	}
	if value == nil {
		value = ""
	}
	switch field.Type {
	case "select":
		options := "<option value=''> Chose</option>"
		for _, item := range field.Items {
			selected := ""
			if fmt.Sprint(value) == fmt.Sprint(item.Value) {
				selected = "selected='selected'"
			}
			options += fmt.Sprintf("<option value='%s' style='background-color: %s' %s>%s</option> ", fmt.Sprint(item.Value), s.BgColor(item.Value, item.BgColor), selected, item.Text)
		}
		return fmt.Sprintf(`<li style="display: %s">%s <select  name="%s" value="%s">%s</select></li>`, hidden, title, name, value, options)
	case "showImg":
		return ""
	case "file":
		return fmt.Sprintf(`
<li>
%s
<input type='text'  name='%s' value='%s'/>
</li>
<li>
<input type='file' id='file' name='file' multiple/> 
<input type='button' value='upload' onclick='handleUploadFile()'/> 
</li>`, title, name, value)
	default:
		return fmt.Sprintf(`<li style="display: %s">%s <input autocomplete="off"  name="%s" type="%s"  step="%f" value="%s"></li>`, hidden, title, name, t, step, value)
	}
}
func (s *view) choseField(field *config.Field) string {
	choseField := field.Field
	if strings.Contains(choseField, ".") {
		choseField = field.QueryFiled
		if choseField == "" {
			panic("，QueryFiled不能为空!!! 当Filed字段包含点号时.说明Field字段不为t1表中的字段，所以此时查询值时应该使用 as 之后的新字段名，我是将其设置在 QueryField 字段中的，所以请您设置 QueryField")
		}
	}
	return choseField
}
func (s *view) Th(field *config.Field) string {
	if field.Title == "" {
		field.Title = field.Field
	}
	if field.Hidden {
		return ""
	}
	return fmt.Sprintf(`<th>%s</th>`, strings.ToUpper(field.Title))
}
func (s *view) Td(field *config.Field, m g.Map) string {
	if field.Hidden {
		return ""
	}
	choseField := s.choseField(field)
	v := fmt.Sprint(m[choseField])
	switch field.Type {
	case "select":
		for _, item := range field.Items {
			if v == fmt.Sprint(item.Value) {
				return fmt.Sprintf("<td><span style='background-color: %s; text-align: center'>%s</span></td>", s.BgColor(v, item.BgColor), item.Text)
			}
		}
	case "showImg":
		fileUrl := fmt.Sprint(field.ShowImg.ImgPrefix, m[field.ShowImg.Field])
		return fmt.Sprintf("<td><a href='%s' target='_blank' ><img src='%s' alt='not fond'></a></td>", fileUrl, fileUrl)
	default:
		return fmt.Sprintf("<td><span>%s</span></td>", v)
	}
	return "<td></td>"
}
func (s *view) QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func (s *view) BindFuncMap() gview.FuncMap {
	return gview.FuncMap{
		"bgColor":     s.BgColor,
		"searchValue": s.SearchValue,
		"formli":      s.FormLi,
		"th":          s.Th,
		"td":          s.Td,
		"queryEscape": s.QueryEscape,
	}
}
