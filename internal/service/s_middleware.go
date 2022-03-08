package service

import (
	"ciel-admin/utility/utils/res"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---middleware-----------------------------------------------------------------------
type middleware struct{}

func Middleware() *middleware { return &middleware{} }
func (s *middleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (s *middleware) AuthAdmin(r *ghttp.Request) {
	user, err := Session().GetAdmin(r)
	if err != nil || user == nil {
		r.Response.RedirectTo("/login")
		return
	}
	b := Role().CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI, r.Method)
	if !b {
		res.Err(errors.New("没有权限"), r)
	}
	r.Middleware.Next()
}
