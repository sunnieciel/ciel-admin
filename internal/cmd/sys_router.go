package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.Server) {
	s.Group("/node", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Node.Path)
		g.GET("/path/add", controller.Node.PathAdd)
		g.GET("/path/edit/:id", controller.Node.PathEdit)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.Node.Del)
		g.POST("/post", controller.Node.Post)
		g.POST("/put", controller.Node.Put)
	})
	s.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.AdminLoginLog.Path)
		g.GET("/path/add", controller.AdminLoginLog.PathAdd)
		g.GET("/path/edit/:id", controller.AdminLoginLog.PathEdit)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.AdminLoginLog.Del)
		g.POST("/post", controller.AdminLoginLog.Post)
		g.POST("/put", controller.AdminLoginLog.Put)
	})
}