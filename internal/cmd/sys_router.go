package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.RouterGroup) {
	s.Group("/adminMessage", func(g *ghttp.RouterGroup) {
		g.Middleware(sys.AuthAdmin)
		g.GET("/path", controller.AdminMessage.Path)
		g.GET("/path/add", controller.AdminMessage.PathAdd)
		g.GET("/path/edit/:id", controller.AdminMessage.PathEdit)
		g.GET("/unreadMsgCount", controller.AdminMessage.UnreadMsgCount)
		g.GET("/clearUnread", controller.AdminMessage.ClearUnreadMsg)
		g.Middleware(sys.LockAction, sys.AdminAction)
		g.GET("/path/del/:id", controller.AdminMessage.Del)
		g.GET("/clear", controller.AdminMessage.Clear)
		g.POST("/post", controller.AdminMessage.Post)
		g.POST("/put", controller.AdminMessage.Put)
	})
}
