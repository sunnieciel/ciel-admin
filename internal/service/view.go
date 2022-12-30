package service

import (
	"ciel-admin/internal/consts"
	"ciel-admin/utility/utils/xhtml"
	"ciel-admin/utility/utils/xtime"
	"ciel-admin/utility/utils/xtrans"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gomarkdown/markdown"
	"math"
	"strings"
)

type view struct{}

var (
	View            = view{}
	chooseEmptyErr  = fmt.Errorf("输入的选择类型为空")
	optionFormatErr = fmt.Errorf("选项类型格式不正确,正确格式为:`value1:label1:class1,value2:label2:class2`,eg: `1:菜单:Tag-info,2:分组:Tag-warning`")
)

func (s view) BindFuncMap() gview.FuncMap {
	return gview.FuncMap{
		"hasPrefix":   s.HasPrefix,
		"toUrlParams": xurl.ToUrlParams,
		"tag":         s.Tag,
		"nodeLevel":   s.NodeLevel,
		"nodeTime":    s.NodeTime,
		"nodeWeek":    s.NodeWeek,
		"nav":         s.Nav,
		"option":      s.Option,
		"chooseSpan":  s.ChooseSpan,
		"img":         s.Img,
		"md":          s.MD,
		"split":       s.Split,
		"trans":       xtrans.T,
		// 查询 input
		"input":      s.Input,
		"inputDate":  s.InputDate,
		"inputHidde": s.InputHidden,
		"options":    s.Options,
		// table
		"th":             s.Th,
		"td":             s.Td,
		"tdImg":          s.TdImg,
		"tdChoose":       s.TdChoose,
		"searchPageSize": s.SearchPageSize,
		"a":              s.A,
		"aDel":           s.ADel,
		"aFun":           s.AFun, // a with function
		"imgSettings":    s.ImgSettings,
		// 添加或修改 tr
		"editTr":                s.EditTr,
		"editTrHidden":          s.EditTrHidden,
		"editTrInputListDesc":   s.EditTrInputListDesc,
		"editTrDesc":            s.EditTrDesc,
		"editTrPass":            s.EditTrPass,
		"editTrReadonly":        s.EditTrReadonly,
		"editTrTextarea":        s.EditTrTextarea,
		"editTrRequired":        s.EditTrRequired,
		"editTrNumber":          s.EditTrNumber,
		"editTrNumberDesc":      s.EditTrNumberDesc,
		"editTrOptions":         s.EditTrOptions,
		"editTrInputOptions":    s.EditTrInputOptions,
		"editTrOptionsReadonly": s.EditTrOptionsReadonly,
		"editTrSubmit":          s.EditTrSubmit,
		// book
		"balance":   s.Balance,
		"topicTime": xtime.TopicTime,
		"IndexNav":  s.IndexNav,
		// sEnglish
		"levelCircle":  s.LevelCircle,
		"englishLevel": s.EnglishLevel,
	}
}

func (s view) HasPrefix(str interface{}, suffix string) bool {
	d := gconv.String(str)
	return strings.HasPrefix(d, suffix)
}
func (s view) Tag(i interface{}, content interface{}) string {
	if content == "" {
		return ""
	}
	return fmt.Sprintf("<span class='%s'>%v</span>", xhtml.SwitchTagClass(gconv.Int(i)), content)
}
func (s view) Split(str string, delimiter string) []string {
	return gstr.SplitAndTrim(str, delimiter)
}
func (s view) Option(in string, flag interface{}) string {
	if in == "" {
		return ""
	}
	split := strings.Split(in, ",")
	if len(split) == 0 {
		return ""
	}
	res := ""
	for _, i := range split {
		d := strings.Split(i, ":")
		if len(d) != 3 {
			return ""
		}
		selected := ""
		if d[0] == gconv.String(flag) {
			selected = "selected"
		}
		res += fmt.Sprintf("<option value='%v' class='%v'  %v>%v</option>", d[0], d[2], selected, d[1])
	}
	return res
}
func (s view) ChooseSpan(options string, flagValue interface{}) string {
	if options == "" {
		panic(chooseEmptyErr)
	}
	split := strings.Split(options, ",")
	if len(split) == 0 {
		panic(chooseEmptyErr)
	}
	res := ""
	for _, i := range split {
		d := strings.Split(i, ":")
		if len(d) != 3 {
			panic(optionFormatErr)
		}
		if d[0] == gconv.String(flagValue) {
			res = fmt.Sprintf("<span  class='%v' >%v</Option>", d[2], d[1])
			break
		}
	}
	if res == "" {
		res = fmt.Sprintf("<span  class='Tag-danger' >ERROR</Option>")
	}
	return res
}
func (s view) Img(in interface{}) string {
	url := gconv.String(in)
	if url == "" {
		return fmt.Sprint("<span class='tag-info fs-12'>暂无图片</span>")
	}
	if !strings.HasPrefix(url, "http") {
		url = consts.ImgPrefix + url
	}
	return fmt.Sprintf("<a href='%s' target='_blank'><Img class='s-icon' src='%s' alt='not fond'></a>", url, url)
}
func (s view) Input(name, nameDesc string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<label class="input">%s<input type="text" name="%s" value="%v" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>`, nameDesc, name, v)
}
func (s view) InputDate(name, nameDesc string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<label class="input">%s<input type="date" name="%s" value="%v" onchange="this.form.submit()"> </label>`, nameDesc, name, v)
}
func (s view) InputHidden(name string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<input type="hidden" name="%s" value="%v" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>`, name, v)
}
func (s view) Options(name, nameDesc, options string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<label class="input">%s <select name="%s" onchange="this.form.submit()"><option value=''>请选择</option>%s</select></label>`,
		nameDesc, name, s.Option(options, v))
}
func (s view) Th(str string) string {
	var (
		res string
	)
	for _, i := range gstr.Split(str, ",") {
		if i != "" {
			res += fmt.Sprintf("<th>%s</th>", i)
		}
	}
	return res
}
func (s view) Td(name string, value interface{}) string {
	if value == nil {
		value = "- - -"
	}
	return fmt.Sprintf(`<td data-label="%s">%v</td>`, name, value)
}
func (s view) TdImg(name string, value interface{}) string {
	return s.Td(name, s.Img(value))
}
func (s view) TdChoose(name string, options string, value interface{}) string {
	return s.Td(name, s.ChooseSpan(options, value))
}
func (s view) SearchPageSize(query map[string]interface{}, sizeIn ...int) string {
	size := query["size"]
	if size == nil {
		size = 10
	}
	return fmt.Sprintf(`<input id="page" name="page" value="1" hidden><input name="size" value="%v" hidden>`, size)
}
func (s view) A(className string, href string, name string, query ...map[string]interface{}) string {
	var q string
	if len(query) > 0 {
		params := xurl.ToUrlParams(query[0])
		if query != nil && params != "" {
			q = fmt.Sprint("?", params)
		}
	}
	return fmt.Sprintf(`<a class="%s" href="%s%s">%s</a>`, className, href, q, name)
}
func (s view) AFun(className string, name, f string) string {
	return fmt.Sprintf(`<a class="%s" href="#" onclick="%s">%s</a>`, className, f, name)
}
func (s view) ADel(href string, query map[string]interface{}) string {
	return fmt.Sprintf(`<a class="tag-danger" href="#" onclick="if(confirm('确认删除?')){location.href='%s?%s'}">删除</a>`, href, xurl.ToUrlParams(query))
}
func (s view) ImgSettings() string {
	return `<img src="/resource/image/settings.png" alt="Settings" width="64" height="64">`
}
func (s view) EditTr(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" value="%v"></td></tr> `,
		title, name, value)
}
func (s view) EditTrHidden(name string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<input type="hidden" name="%s" value="%v" >`, name, value)
}
func (s view) EditTrDesc(name string, title string, value interface{}, desc string) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" value="%v"></td><td class="color-desc-02 fs-12">%s</td></tr> `,
		title, name, value, desc)
}
func (s view) EditTrInputListDesc(name, title, options, desc string) string {
	var (
		option string
	)
	for _, i := range strings.Split(options, ",") {
		temp := gstr.TrimAll(i)
		if temp != "" {
			option += fmt.Sprintf(`<option>%s</option>`, temp)
		}
	}
	return fmt.Sprintf(`<tr>
<td align="right">%s</td>
<td>
<input type="text" list="list-%s" name="%s">
<datalist id="list-%s">
%s
</datalist>
</td>
<td class="color-desc-02 fs-12">%s</td>
</tr>`, title, name, name, name, option, desc)
}
func (s view) EditTrPass(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input type="password" required name="%s" value="%v"></td></tr> `,
		title, name, value)
}
func (s view) EditTrReadonly(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" readonly value="%v"></td></tr> `,
		title, name, value)
}
func (s view) EditTrRequired(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" required value="%v"></td></tr> `,
		title, name, value)
}
func (s view) EditTrTextarea(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><textarea name="%s" >%v</textarea></td></tr> `,
		title, name, value)
}
func (s view) EditTrNumber(name string, title string, value interface{}, step float64, min, max float64) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input type="number"  name="%s" value="%v" step='%f' min="%f" max="%f"></td></tr> `,
		title, name, value, step, min, max)
}
func (s view) EditTrNumberDesc(name string, title string, value interface{}, step float64, min, max float64, desc string) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input type="number"  name="%s" value="%v" step='%f' min="%f" max="%f"></td><td class='color-desc-02 fs-12'>%s</td></tr> `,
		title, name, value, step, min, max, desc)
}
func (s view) EditTrOptions(name, title, options string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td align="right">%s</td><td><select name="%s">%s</select></td></tr>`, title, name, s.Option(options, value))
}
func (s view) EditTrInputOptions(name, title, options string, value interface{}, id string) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td align="right">%s</td><td><input name=%s list='%s' value='%s'><datalist id='%s'>%s</datalist></td></tr>`, title, name, id, value, id, s.Option(options, value))
}
func (s view) EditTrOptionsReadonly(name, title, options string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td align="right">%s</td><td><select name="%s" disabled>%s</select></td></tr>`, title, name, s.Option(options, value))
}
func (s view) EditTrSubmit() string {
	return fmt.Sprintf(`<tr><td></td><td><button class="btn-info" type="submit">提交</button></td></tr>`)
}

func (s view) MD(data string) string {
	data = gstr.Replace(data, "\r", "")
	return string(markdown.ToHTML([]byte(data), nil, nil))
}

// Nav web
func (s view) Nav(in []interface{}, path interface{}) string {
	res := ""
	nav := "<div id='nav'>"
	subNav := "<div id='sub-nav'>"
	for _, i := range in {
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
func (s view) Balance(in interface{}) string {
	money := gconv.Float64(in)
	var (
		bronze   = `<img class="mlr-3" height="16" border="0" src="/resource/image/bronze@2x.png" alt="">`
		silver   = `<img class="mlr-3" height="16" border="0" src="/resource/image/silver@2x.png" alt="">`
		gold     = `<img class="mlr-3" height="16" border="0" src="/resource/image/gold@2x.png" alt="">`
		b, si, g string
		m        = fmt.Sprintf("%.f", math.Abs(money))
	)
	if money == 0 {
		return fmt.Sprint(0, bronze)
	}
	switch {
	case len(m) > 4:
		g = m[:len(m)-4]
		si = m[len(m)-4 : len(m)-2]
		b = m[len(m)-2:]
	case len(m) > 2:
		si = m[:len(m)-2]
		b = m[len(m)-2:]
	default:
		b = m
	}
	res := ""
	if g != "" && g != "-" {
		res += fmt.Sprint(g, gold)
	}
	if si != "" && si != "00" && si != "-" {
		if gstr.HasPrefix(si, "0") {
			si = si[1:]
		}
		res += fmt.Sprint(si, silver)
	}
	if b != "" && b != "00" {
		if gstr.HasPrefix(b, "0") {
			b = b[1:]
		}
		res += fmt.Sprint(b, bronze)
	}
	if money < 0 { // if less than 0 add "-"
		res = fmt.Sprint("-", res)
	}
	return res
}

// node

func (s view) NodeWeek(y, m, d interface{}) string {
	if y != nil && m != nil && d != nil {
		t := gtime.NewFromStr(fmt.Sprintf("%v-%v-%v", y, m, d))
		format := t.Format("w")
		week := format
		if week == "0" {
			week = "天"
		}
		switch gconv.Int(t.Format("w")) {
		case 0:
			d = "日"
		case 1:
			d = "一"
		case 2:
			d = "二"
		case 3:
			d = "三"
		case 4:
			d = "四"
		case 5:
			d = "五"
		case 6:
			d = "六"
		default:
		}
		return s.Tag(format, fmt.Sprintf("星期%s", d))
	}
	return ""
}
func (s view) NodeTime(y, m, d interface{}) string {
	res := s.Tag(y, fmt.Sprint(y, "年"))
	if m != nil {
		res += s.Tag(m, fmt.Sprint(m, "月"))
	}
	if d != nil && gconv.Int(d) != 0 {
		res += s.Tag(1, fmt.Sprint(d, "日"))
	}
	return res
}
func (s view) NodeLevel(i interface{}) string {
	content := ""
	i2 := gconv.Int(i)
	switch i2 {
	case 1:
		content = "普通"
	case 2:
		content = "稀有"
	case 3:
		content = "传承"
	case 4:
		content = "唯一"
	case 5:
		content = "史诗"
	case 6:
		content = "传说"
	case 7:
		content = "神话"
	default:
		content = "普通"
	}
	return s.Tag(i, content)
}

func (s view) EnglishLevel(num uint) string {
	class := "tag-info"
	msg := "普通"
	switch num {
	case 1:
		class = "tag-info"
		msg = "普通"
	case 2:
		class = "tag-success"
		msg = "稀有"
	case 3:
		class = "tag-primary"
		msg = "传承"
	case 4:
		class = "tag-warning"
		msg = "唯一"
	case 5:
		class = "tag-brown"
		msg = "史诗"
	case 6:
		class = "tag-purple"
		msg = "传说"
	case 7:
		class = "tag-danger"
		msg = "神话"
	}
	return fmt.Sprintf(`<span class="%s">%s</span>`, class, msg)
}
func (s view) LevelCircle(num uint) string {
	return fmt.Sprint(gconv.Float64(num) * 2.7)
}
func (s view) IndexNav(current string) string {
	var (
		arr = g.Cfg().MustGet(nil, "nav").Vars()
		res string
	)
	for _, i := range arr {
		class := "link-2"
		m := i.Map()
		if m["name"] == current {
			class = "link-2-active"
		}
		res += fmt.Sprintf("<a class='%s' href='%s'>%s</a>", class, m["href"], m["name"])
	}
	return res
}
