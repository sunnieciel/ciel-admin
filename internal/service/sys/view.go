package sys

import (
	"github.com/gogf/gf/v2/os/gview"
	"net/url"
)

func QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func BindFuncMap() gview.FuncMap {
	return gview.FuncMap{}
}
