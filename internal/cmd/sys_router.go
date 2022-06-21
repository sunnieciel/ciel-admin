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
		g.Middleware(sys.LockAction, sys.AdminAction)
		g.DELETE("/batch", controller.User.Del)
		g.POST("/", controller.User.Post)
		g.PUT("/", controller.User.Put)
	})
	s.Group("/loginLog", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.LoginLog.Path)
		g.GET("/", controller.LoginLog.List)
		g.GET("/:id", controller.LoginLog.GetById)
		g.Middleware(sys.LockAction, sys.AdminAction)
		g.DELETE("/batch", controller.LoginLog.Del)
		g.POST("/", controller.LoginLog.Post)
		g.PUT("/", controller.LoginLog.Put)
	})
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
		g.DELETE("/batch", controller.LoginLog.Del)
		g.POST("/", controller.LoginLog.Post)
		g.PUT("/", controller.LoginLog.Put)
	})
	s.Group("/node", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Node.Path)
		g.GET("/", controller.Node.List)
		g.GET("/:id", controller.Node.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/batch", controller.Node.Del)
		g.POST("/", controller.Node.Post)
		g.PUT("/", controller.Node.Put)
	})
	s.Group("/thing", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.Thing.Path)
		g.GET("/", controller.Thing.List)
		g.GET("/:id", controller.Thing.GetById)
		g.GET("/options", controller.Thing.Options)
		g.Middleware(sys.LockAction)
		g.DELETE("/batch", controller.Thing.Del)
		g.POST("/", controller.Thing.Post)
		g.PUT("/", controller.Thing.Put)
	})
	s.Group("/thingRecord", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.ThingRecord.Path)
		g.GET("/", controller.ThingRecord.List)
		g.GET("/:id", controller.ThingRecord.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/batch", controller.ThingRecord.Del)
		g.POST("/", controller.ThingRecord.Post)
		g.PUT("/", controller.ThingRecord.Put)
	})
}
