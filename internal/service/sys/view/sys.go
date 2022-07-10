package view

import (
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
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
func option(arr []interface{}, flag interface{}) string {
	res := ""
	for _, i := range arr {
		value := gconv.String(i.(map[string]interface{})["value"])
		class := gconv.String(i.(map[string]interface{})["class"])
		label := gconv.String(i.(map[string]interface{})["label"])
		selected := ""
		if value == gconv.String(flag) {
			selected = "selected"
		}
		res += fmt.Sprintf("<option value='%v' class='%v'  %v>%v</option>", value, class, selected, label)
	}
	return res
}
