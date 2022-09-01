package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/admin"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.RouterGroup) {
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(admin.AuthMiddleware)
		g.GET("/", controller.User.Index)
		g.GET("/add", controller.User.AddIndex)
		g.GET("/edit/:id", controller.User.EditIndex)
		g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
		g.GET("/del/:id", controller.User.Del)
		g.POST("/post", controller.User.Post)
		g.POST("/put", controller.User.Put)
	})
}
