package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.Server) {
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.User.Path)
		g.GET("/", controller.User.List)
		g.GET("/:id", controller.User.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/:id", controller.User.Del)
		g.POST("/", controller.User.Post)
		g.PUT("/", controller.User.Put)
	})
	s.Group("/loginLog", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.LoginLog.Path)
		g.GET("/", controller.LoginLog.List)
		g.GET("/:id", controller.LoginLog.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/:id", controller.LoginLog.Del)
		g.POST("/", controller.LoginLog.Post)
		g.PUT("/", controller.LoginLog.Put)
	})
	s.Group("/admin", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Admin.Path)
		g.GET("/", controller.Admin.List)
		g.GET("/:id", controller.Admin.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/:id", controller.Admin.Del)
		g.POST("/", controller.Admin.Post)
		g.PUT("/", controller.Admin.Put)
	})
	s.Group("/file", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.File.Path)
		g.GET("/", controller.File.List)
		g.GET("/:id", controller.File.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/:id", controller.File.Del)
		g.POST("/", controller.File.Post)
		g.PUT("/", controller.File.Put)
	})
}
