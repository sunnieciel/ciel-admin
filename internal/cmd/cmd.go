package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/admin"
	"ciel-admin/internal/service/sys"
	"ciel-admin/internal/service/view"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
	_ "net/http/pprof"
	"time"
)

var (
	Main = gcmd.Command{
		Name:        "ciel",
		Usage:       "ciel",
		Brief:       "start http server",
		Description: "hello i'm ciel",
		Arguments:   nil,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化服务
			g.Log().SetFlags(glog.F_FILE_LONG | glog.F_TIME_DATE | glog.F_TIME_MILLI)
			g.View().BindFuncMap(view.BindFuncMap())
			sys.Init(ctx)
			s := g.Server()
			registerInterface(s) // 注册对外提供功能的接口
			s.Group("/", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.Home.IndexPage)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.WhiteIpMiddleware) // 白名单过滤  在字典表中为空时，这里不会进行检查的
				controller.Menu.RegisterRouter(g)
				controller.Api.RegisterRouter(g)
				controller.RoleMenu.RegisterRouter(g)
				controller.Role.RegisterRouter(g)
				controller.RoleApi.RegisterRouter(g)
				controller.Dict.RegisterRouter(g)
				controller.Admin.RegisterRouter(g)
				controller.OperationLog.RegisterRouter(g)
				controller.AdminLoginLog.RegisterRouter(g)
				controller.File.RegisterRouter(g)
				controller.Sys.RegisterRouter(g)
				controller.Gen.RegisterRouter(g)
				registerGenFileRouter(g) // 注册生成的代码路由
				g.GET("/login", controller.Admin.LoginPage)
				g.Middleware(admin.AuthMiddleware)
				g.GET("/quotations", controller.Sys.Quotations)
			})
			go func() {
				var ctx = context.Background()
				time.Sleep(time.Second * 1)
				port, err := g.Cfg().Get(ctx, "server.address")
				if err != nil {
					panic(err)
				}
				rootIp, err := g.Cfg().Get(ctx, "server.rootIp")
				g.Log().Infof(nil, "Server start at :http://%s%s/admin/login", rootIp, port)
			}()
			s.Run()
			return nil
		},
		FuncWithValue: nil,
		HelpFunc:      nil,
		Examples:      "",
		Additional:    "",
		Strict:        false,
		Config:        "",
	}
)
