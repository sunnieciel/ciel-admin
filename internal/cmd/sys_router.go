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
		g.GET("/list", controller.User.List)
		g.GET("/getById", controller.User.GetById)
		g.Middleware(sys.LockAction)
		g.DELETE("/del", controller.User.Del)
		g.POST("/post", controller.User.Post)
		g.PUT("/put", controller.User.Put)
	})
}
