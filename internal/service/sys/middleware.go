package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/utility/utils/xjwt"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
)

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

var (
	Uid = "userInfoKey"
)

func UserAuth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
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
