package admin

import (
	"ciel-admin/internal/logic"
	"github.com/gogf/gf/v2/net/ghttp"
)

// AuthMiddleware auth admin
func AuthMiddleware(r *ghttp.Request) {
	logic.Admin.AuthMiddleware(r)
}
func LockMiddleware(r *ghttp.Request) {
	logic.Admin.LockMiddleware(r)
}
func ActionMiddleware(r *ghttp.Request) {
	logic.Admin.ActionMiddleware(r)
}
