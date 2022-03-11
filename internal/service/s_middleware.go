package service

import (
	"ciel-admin/utility/utils/res"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---sMiddleware-----------------------------------------------------------------------
type sMiddleware struct{}

var insMiddleware = new(sMiddleware)

func Middleware() *sMiddleware { return insMiddleware }
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (s *sMiddleware) AuthAdmin(r *ghttp.Request) {
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
