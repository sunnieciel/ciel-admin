package user

import (
	"ciel-admin/internal/consts"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xuser"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func AuthMiddleware(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(xuser.UidKey, userInfo.Uid)
	c.Middleware.Next()
}
