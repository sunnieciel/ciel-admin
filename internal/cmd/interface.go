package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 对外提供的接口放在此文件中

func registerInterface(s *ghttp.Server) {
	s.Group("/v1", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.CORS)
		g.Group("/dict", func(g *ghttp.RouterGroup) {
			g.GET("/key/:key", controller.Sys.GetDictByKey)
		})
	})
}
