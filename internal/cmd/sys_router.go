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
}
