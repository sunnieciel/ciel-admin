// Package view html相关的方法/*
package view

import (
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

// Input 拼接input框
func Input(name, nameDesc string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<label class="input">%s<input type="text" name="%s" value="%v" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>`, nameDesc, name, v)
}

func InputHidden(name string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<input type="hidden" name="%s" value="%v" onkeydown="if(event.keyCode===13)this.form.submit()"> </label>`, name, v)
}

func Options(name, nameDesc, options string, query map[string]interface{}) string {
	v := query[name]
	if v == nil {
		v = ""
	}
	return fmt.Sprintf(`<label class="input">%s <select name="%s" onchange="this.form.submit()"><option value=''>请选择</option>%s</select></label>`,
		nameDesc, name, Option(options, v))
}

// Th 拼接 表th
// 输入 id,pid
// 返回 <th>id</th><th>pid</th>
func Th(str string) string {
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

// Td 拼接表 td
// 输入 ID 1
// 输出 <td data-label="ID">1</td>
func Td(name string, value interface{}) string {
	if value == nil {
		value = "- - -"
	}
	return fmt.Sprintf(`<td data-label="%s">%v</td>`, name, value)
}

// TdImg td 拼接图片
func TdImg(name string, value interface{}) string {
	return Td(name, Img(value))
}

// TdChoose td 选择 span展示
// 输入 name 状态
// 输入 options 1:正常:tag-info,2:禁用:tag-danger
// 输入 value 1
// 输出 <td data-label="状态"><span class="tag-info">正常</span></td>
func TdChoose(name string, options string, value interface{}) string {
	return Td(name, ChooseSpan(options, value))
}

func SearchPageSize(query map[string]interface{}) string {
	size := query["size"]
	if size == nil {
		size = 10
	}
	return fmt.Sprintf(`<input id="page" name="page" value="1" hidden><input name="size" value="%v" hidden>`, size)
}
func A(className string, href string, name string, query ...map[string]interface{}) string {
	var q string
	if len(query) > 0 {
		params := xurl.ToUrlParams(query[0])
		if query != nil && params != "" {
			q = fmt.Sprint("?", params)
		}
	}
	return fmt.Sprintf(`<a class="%s" href="%s%s">%s</a>`, className, href, q, name)
}

func AFun(className string, name, f string) string {
	return fmt.Sprintf(`<a class="%s" href="#" onclick="%s">%s</a>`, className, f, name)
}
func ADel(href string, query map[string]interface{}) string {
	return fmt.Sprintf(`<a class="tag-danger" href="#" onclick="if(confirm('确认删除?')){location.href='%s?%s'}">删除</a>`, href, xurl.ToUrlParams(query))
}
func ImgSettings() string {
	return `<img src="/resource/image/settings.png" alt="Settings" width="64" height="64">`
}
func EditTr(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" value="%v"></td></tr> `,
		title, name, value)
}

func EditTrDesc(name string, title string, value interface{}, desc string) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" value="%v"></td><td class="color-desc-02 fs-12">%s</td></tr> `,
		title, name, value, desc)
}

// EditTrInputListDesc
// 输入 name=group
// 输入 一级分类菜单
// 输入 系统:,工具,用户
// 输入 这是一个选项输入框
// 输入 用户
// 输出
func EditTrInputListDesc(name, title, options, desc string) string {
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

func EditTrPass(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input type="password" required name="%s" value="%v"></td></tr> `,
		title, name, value)
}

func EditTrReadonly(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" readonly value="%v"></td></tr> `,
		title, name, value)
}

func EditTrRequired(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input name="%s" required value="%v"></td></tr> `,
		title, name, value)
}

func EditTrTextarea(name string, title string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><textarea name="%s" >%v</textarea></td></tr> `,
		title, name, value)
}

func EditTrNumber(name string, title string, value interface{}, step float64, min, max float64) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td width="160" align="right">%s</td><td align="left"><input type="number"  name="%s" value="%v" step='%f' min="%f" max="%f"></td></tr> `,
		title, name, value, step, min, max)
}

// EditTrOption 编辑或添加 select 类型
// 输入 name type
// 输入 title 类型
// 输入 options
// 输入 value 1
// 输出 <tr><td align="right">类型</td><td><select name="type"><option value="1" class="tag-info">菜单</option><option value="2" class="tag-warning">分组</option></select></td></tr>
func EditTrOption(name, title, options string, value interface{}) string {
	if value == nil {
		value = ""
	}
	return fmt.Sprintf(`<tr><td align="right">%s</td><td><select name="%s">%s</select></td></tr>`, title, name, Option(options, value))
}
func EditTrSubmit() string {
	return fmt.Sprintf(`<tr><td></td><td><button class="btn-info" type="submit">提交</button></td></tr>`)
}
