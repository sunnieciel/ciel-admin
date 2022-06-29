package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.Server) {

	s.Group("/thingRecord", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.ThingRecord.Path)
		g.GET("/path/add", controller.ThingRecord.PathAdd)
		g.GET("/path/edit/:id", controller.ThingRecord.PathEdit)
		g.Middleware(sys.LockAction)
		g.GET("/path/del/:id", controller.ThingRecord.Del)
		g.POST("/post", controller.ThingRecord.Post)
		g.POST("/put", controller.ThingRecord.Put)
	})
}
