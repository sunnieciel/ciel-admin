package view

import (
	"ciel-admin/internal/consts"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func nav(s []interface{}, path interface{}) string {
	res := ""
	nav := "<div id='nav'>"
	subNav := "<div id='sub-nav'>"
	for _, i := range s {
		pClass := ""
		p := gjson.New(i)
		children := p.GetJsons("Children")
		for _, c := range children {
			if path == c.Get("path").String() { // 如果当前的页面地址在此 children  则父nav 加一个 link-2-active,
				pClass = "link-2-active"
			}
		}
		subDisplay := "style='display: none'" // 子链接是否展示
		for _, c := range children {
			if pClass != "" {
				subDisplay = ""
			}
			currentPageClas := "" // 当前页面的class
			if path == c.Get("path").String() {
				currentPageClas = "tag-blue"
			}
			subNav += fmt.Sprintf("<a class='link-3 %s' data='%s' href='%s' %s>%s</a>",
				currentPageClas,
				p.Get("name"),
				c.Get("path"),
				subDisplay,
				c.Get("name"),
			)
		}
		name := p.Get("name").String()
		nav += fmt.Sprintf(`<a class='link-2 %s' data="%s" >%s</a>`, pClass, name, name)
		subDisplay = ""
		pClass = ""
	}
	nav += "</div>"
	res += nav
	subNav += "</div>"
	res += subNav
	res += ""
	return res
}

var (
	chooseEmptyErr  = fmt.Errorf("输入的选择类型为空")
	optionFormatErr = fmt.Errorf("选项类型格式不正确,正确格式为:`value1:label1:class1,value2:label2:class2`,eg: `1:菜单:tag-info,2:分组:tag-warning`")
)

func option(in string, flag interface{}) string {
	if in == "" {
		panic(chooseEmptyErr)
	}
	split := strings.Split(in, ",")
	if len(split) == 0 {
		panic(chooseEmptyErr)
	}
	res := ""
	for _, i := range split {
		d := strings.Split(i, ":")
		if len(d) != 3 {
			panic(optionFormatErr)
		}
		selected := ""
		if d[0] == gconv.String(flag) {
			selected = "selected"
		}
		res += fmt.Sprintf("<option value='%v' class='%v'  %v>%v</option>", d[0], d[2], selected, d[1])
	}
	return res
}

func chooseSpan(in string, flag interface{}) string {
	if in == "" {
		panic(chooseEmptyErr)
	}
	split := strings.Split(in, ",")
	if len(split) == 0 {
		panic(chooseEmptyErr)
	}
	res := ""
	for _, i := range split {
		d := strings.Split(i, ":")
		if len(d) != 3 {
			panic(optionFormatErr)
		}
		if d[0] == gconv.String(flag) {
			res = fmt.Sprintf("<span  class='%v' >%v</option>", d[2], d[1])
			break
		}
	}
	if res == "" {
		res = fmt.Sprintf("<span  class='tag-danger' >ERROR</option>")
	}
	return res
}
func img(in interface{}) string {
	url := gconv.String(in)
	if url == "" {
		return fmt.Sprint("<span class='tag-normal'>暂无图片</span>")
	}
	if !strings.HasPrefix(url, "http") {
		url = consts.ImgPrefix + url
	}
	return fmt.Sprintf("<a href='%s' target='_blank'><img class='s-icon' src='%s' alt='not fond'></a>", url, url)
}
