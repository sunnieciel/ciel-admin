package service

import (
	"github.com/gogf/gf/v2/os/gview"
	"net/url"
)

// ---view-------------------------------------------------------------------------------
type view struct{}

func View() *view { return &view{} }
func (s *view) QueryEscape(str string) string {
	unescape, err := url.QueryUnescape(str)
	if err != nil {
		panic(err)
	}
	return unescape
}
func (s *view) BindFuncMap() gview.FuncMap {
	return gview.FuncMap{}
}
