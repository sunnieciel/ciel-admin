package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.Server) {
	s.Group("/menu", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Menu.Path)
		g.GET("/path/add", controller.Menu.PathAdd)
		g.GET("/path/edit/:id", controller.Menu.PathEdit)
		g.GET("/level1", controller.Sys.Level1)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.Menu.Del)
		g.POST("/post", controller.Menu.Post)
		g.POST("/put", controller.Menu.Put)
	})
	s.Group("/api", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Api.Path)
		g.GET("/path/add", controller.Api.PathAdd)
		g.GET("/path/edit/:id", controller.Api.PathEdit)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.Api.Del)
		g.POST("/post", controller.Api.Post)
		g.POST("/put", controller.Api.Put)
	})
	s.Group("/dict", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Dict.Path)
		g.GET("/path/add", controller.Dict.PathAdd)
		g.GET("/path/edit/:id", controller.Dict.PathEdit)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.Dict.Del)
		g.POST("/post", controller.Dict.Post)
		g.POST("/put", controller.Dict.Put)
	})
}
