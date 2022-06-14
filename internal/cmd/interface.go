package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerInterface(s *ghttp.Server) {
	s.Group("/v1", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.CORS)
		g.Group("/dict", func(g *ghttp.RouterGroup) {
			g.GET("/key/:key", controller.Dict.GetByKey)
		})
	})
}
