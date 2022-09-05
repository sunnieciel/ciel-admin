package sys

import (
	"ciel-admin/internal/consts"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// WhiteIpMiddleware white ip
func WhiteIpMiddleware(r *ghttp.Request) {
	ips := consts.WhiteIps
	if ips != "" {
		if !gstr.Contains(consts.WhiteIps, r.GetClientIp()) {
			r.Response.WriteStatus(http.StatusForbidden, fmt.Sprintf("%s ip error", r.GetClientIp()))
			r.Exit()
		}
	}
	r.Middleware.Next()
}
