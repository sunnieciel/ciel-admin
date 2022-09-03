package cmd

import (
	"ciel-admin/internal/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.RouterGroup) {
	controller.User.RegisterRouter(s)
	controller.UserLoginLog.RegisterRouter(s)
}