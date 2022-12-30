package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	_ "net/http/pprof"
	"time"
)

var (
	Main = gcmd.Command{
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			service.System.Init(ctx)
			routers(s)
			otherWorks(ctx)
			s.Run()
			return
		},
		Strict: false,
	}
)

func otherWorks(ctx context.Context) {
	go func() {
		time.Sleep(time.Second * 1)
		port, err := g.Cfg().Get(ctx, "server.address")
		if err != nil {
			panic(err)
		}
		rootIp, err := g.Cfg().Get(ctx, "server.rootIp")
		g.Log().Infof(nil, "Server start at :http://%s%s/admin/login", rootIp, port)
	}()
}
func routers(s *ghttp.Server) {
	s.Group("/admin", func(g *ghttp.RouterGroup) { registerGenFileRouter(g) })
	s.Group("/v1", func(g *ghttp.RouterGroup) { v1Rooters(g) })
}
func v1Rooters(g *ghttp.RouterGroup) {
	g.Middleware(ghttp.MiddlewareHandlerResponse, service.System.MiddlewareCORS)
	g.Group("/user", func(g *ghttp.RouterGroup) {
		g.POST("/register", controller.User.Register)
		g.POST("/login", controller.User.Login)
		g.GET("/icons", controller.User.Icons)
		g.Middleware(service.User.MiddlewareAuth)
		g.GET("/info", controller.User.Info)
		g.POST("/updatePass", controller.User.UpdatePassByUser)
		g.POST("/updateNickname", controller.User.UpdateNickname)
		g.POST("/updateIcon", controller.User.UpdateIcon)
	})
	g.Group("/wallet", func(g *ghttp.RouterGroup) {
		g.Middleware(service.User.MiddlewareAuth)
		g.POST("/setPass", controller.Wallet.SetPass)
		g.POST("/updatePass", controller.Wallet.UpdatePass)
		g.GET("/topUpCategory", controller.Wallet.TopUpCategory)
		g.POST("/createTopUp", controller.Wallet.CreateTopUp)
		g.GET("/listTopUp", controller.Wallet.ListTopUp)
		g.GET("/listChangeTypes", controller.Wallet.ListChangeTypes)
		g.GET("/listChangeLogs", controller.Wallet.ListChangeLogs)
		g.GET("/getInfo", controller.Wallet.GetInfo)
	})
	g.Group("/sys", func(g *ghttp.RouterGroup) {
		g.GET("/allDict", controller.Sys.ListAllDict)
		g.GET("/dict", controller.Sys.GetDictByKey)
		g.GET("/banners", controller.Banner.List)
		g.Middleware(service.User.MiddlewareAuth)
		g.POST("/uploadImg", controller.Sys.UploadImg)
	})
}

// Note: Please keep this function in the end of the file. When you generation code, it can auto register routers.
func registerGenFileRouter(s *ghttp.RouterGroup) {
	s.GET("/login", controller.Admin.IndexLogin)                                   // login page
	s.Middleware(service.System.MiddlewareWhiteIp, service.Admin.MiddlewareUnread) // 白名单过滤  在字典表中为空时，这里不会进行检查的
	controller.Menu.RegisterRouter(s)                                              // Menu
	controller.Api.RegisterRouter(s)                                               // Api
	controller.Role.RegisterRouter(s)                                              // Role
	controller.RoleMenu.RegisterRouter(s)                                          // RoleMenu
	controller.RoleApi.RegisterRouter(s)                                           // RoleApi
	controller.Dict.RegisterRouter(s)                                              // Dict
	controller.Admin.RegisterRouter(s)                                             // Admin
	controller.OperationLog.RegisterRouter(s)                                      // OperationLog
	controller.AdminLoginLog.RegisterRouter(s)                                     // AdminLoginLog
	controller.AdminMessage.RegisterRouter(s)                                      // Admin message
	controller.File.RegisterRouter(s)                                              // File
	controller.Sys.RegisterRouter(s)                                               // Sys
	controller.Gen.RegisterRouter(s)                                               // Gen
	controller.User.RegisterRouter(s)                                              // 用户
	controller.UserLoginLog.RegisterRouter(s)                                      // 登录日志
	controller.Wallet.RegisterRouter(s)                                            // 金币
	controller.WalletChangeType.RegisterRouter(s)                                  // 账变类型
	controller.WalletChangeLog.RegisterRouter(s)                                   // 账变日志
	controller.WalletStatisticsLog.RegisterRouter(s)                               // 账变统计
	controller.WalletReport.RegisterRouter(s)                                      // 账变报表
	controller.WalletTopUpApplication.RegisterRouter(s)                            // 充值订单
	controller.Banner.RegisterRouter(s)                                            // banner
}
