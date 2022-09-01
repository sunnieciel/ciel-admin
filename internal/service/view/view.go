package view

import (
	"ciel-admin/utility/utils/xtrans"
	"ciel-admin/utility/utils/xurl"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
	"strings"
)

func QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func HasPrefix(str interface{}, suffix string) bool {
	s := gconv.String(str)
	return strings.HasPrefix(s, suffix)
}
func BindFuncMap() gview.FuncMap {
	return gview.FuncMap{
		"hasPrefix":   HasPrefix,
		"toUrlParams": xurl.ToUrlParams,
		"tag":         Tag,
		"nodeLevel":   NodeLevel,
		"nodeTime":    NodeTime,
		"nodeWeek":    NodeWeek,
		"nav":         Nav,
		"option":      Option,
		"chooseSpan":  ChooseSpan,
		"img":         Img,
		"md":          MD,
		"trans":       xtrans.T,
		// 查询 input
		"input":   Input,
		"options": Options,
		// table
		"th":             Th,
		"td":             Td,
		"tdImg":          TdImg,
		"tdChoose":       TdChoose,
		"searchPageSize": SearchPageSize,
		"a":              A,
		"aDel":           ADel,
		"aFun":           AFun, // a with function
		"imgSettings":    ImgSettings,
		// 添加或修改 tr
		"editTr":         EditTr,
		"editTrDesc":     EditTrDesc,
		"editTrPass":     EditTrPass,
		"editTrReadonly": EditTrReadonly,
		"editTrTextarea": EditTrTextarea,
		"editTrRequired": EditTrRequired,
		"editTrNumber":   EditTrNumber,
		"editTrOptions":  EditTrOption,
		"editTrSubmit":   EditTrSubmit,
	}
}
