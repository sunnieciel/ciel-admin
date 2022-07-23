package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"ciel-admin/internal/service/sys/view"
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
			sys.Init(ctx)
			g.View().BindFuncMap(view.BindFuncMap())
			s := g.Server()
			registerInterface(s) // 注册对外提供功能的接口
			s.EnableAdmin("/debut/admin")
			s.Group("/", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.Home.IndexPage)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.MiddlewareWhiteIp) // 白名单过滤  在字典表中为空时，这里不会进行检查的
				registerGenFileRouter(g)            // 注册生成的代码路由
				g.Group("/", func(g *ghttp.RouterGroup) {
					g.GET("/login", controller.Admin.LoginPage)
					g.GET("/to/:name", controller.Sys.To)
					g.GET("/quotations", controller.Sys.Quotations)
					g.Middleware(sys.AuthAdmin)
				})
				g.Group("/menu", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.Menu.Path)
					g.GET("/path/add", controller.Menu.PathAdd)
					g.GET("/path/edit/:id", controller.Menu.PathEdit)
					g.GET("/level1", controller.Sys.Level1)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.Menu.Del)
					g.POST("/post", controller.Menu.Post)
					g.POST("/put", controller.Menu.Put)
				})
				g.Group("/api", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.Api.Path)
					g.GET("/path/add", controller.Api.PathAdd)
					g.GET("/path/edit/:id", controller.Api.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.Api.Del)
					g.POST("/post", controller.Api.Post)
					g.POST("/put", controller.Api.Put)
				})
				g.Group("/dict", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.Dict.Path)
					g.GET("/path/add", controller.Dict.PathAdd)
					g.GET("/path/edit/:id", controller.Dict.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.Dict.Del)
					g.POST("/post", controller.Dict.Post)
					g.POST("/put", controller.Dict.Put)
				})
				g.Group("/admin", func(g *ghttp.RouterGroup) {
					g.GET("/getCaptcha", controller.Sys.GetCaptcha) // 获取验证码
					g.POST("/login", controller.Admin.Login)
					g.Middleware(sys.AuthAdmin)
					g.GET("/logout", controller.Admin.Logout)
					g.GET("/path", controller.Admin.Path)
					g.GET("/path/add", controller.Admin.PathAdd)
					g.GET("/path/edit/:id", controller.Admin.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.PUT("/updatePwd", controller.Admin.UpdatePwd)
					g.PUT("/updatePwdWithoutOldPwd", controller.Admin.UpdatePwdWithoutOldPwd)
					g.PUT("/updateUname", controller.Admin.UpdateUname)
					g.GET("/path/del/:id", controller.Admin.Del)
					g.POST("/post", controller.Admin.Post)
					g.POST("/put", controller.Admin.Put)
				})
				g.Group("/role", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.Role.Path)
					g.GET("/path/add", controller.Role.PathAdd)
					g.GET("/path/edit/:id", controller.Role.PathEdit)
					g.GET("/nomenus", controller.Role.RoleNoMenus)
					g.GET("/noapis", controller.Role.RoleNoApis)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.Role.Del)
					g.POST("/post", controller.Role.Post)
					g.POST("/put", controller.Role.Put)
				})
				g.Group("/roleMenu", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.RoleMenu.Path)
					g.GET("/path/add", controller.RoleMenu.PathAdd)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.RoleMenu.Del)
					g.POST("/post", controller.RoleMenu.Post)
				})
				g.Group("/roleApi", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.RoleApi.Path)
					g.GET("/path/add", controller.RoleApi.PathAdd)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.RoleApi.Del)
					g.GET("/clear/:rid", controller.RoleApi.Clear)
					g.POST("/post", controller.RoleApi.Post)
				})
				g.Group("/operationLog", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.OperationLog.Path)
					g.GET("/path/add", controller.OperationLog.PathAdd)
					g.GET("/path/edit/:id", controller.OperationLog.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.OperationLog.Del)
					g.POST("/post", controller.OperationLog.Post)
					g.POST("/put", controller.OperationLog.Put)
					g.GET("/clear", controller.OperationLog.Clear)
				})
				g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.AdminLoginLog.Path)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.AdminLoginLog.Del)
					g.GET("/clear", controller.AdminLoginLog.Clear)
				})
				g.Group("/file", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.File.Path)
					g.GET("/path/add", controller.File.PathAdd)
					g.GET("/path/edit/:id", controller.File.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.File.Del)
					g.POST("/post", controller.File.Post)
					g.POST("/put", controller.File.Put)
					g.POST("/upload", controller.File.Upload)
				})
				g.Group("/node", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.Node.Path)
					g.GET("/path/add", controller.Node.PathAdd)
					g.GET("/path/edit/:id", controller.Node.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.Node.Del)
					g.POST("/post", controller.Node.Post)
					g.POST("/put", controller.Node.Put)
				})
				g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
					g.Middleware(sys.AuthAdmin)
					g.GET("/path", controller.AdminLoginLog.Path)
					g.GET("/path/add", controller.AdminLoginLog.PathAdd)
					g.GET("/path/edit/:id", controller.AdminLoginLog.PathEdit)
					g.Middleware(sys.LockAction, sys.AdminAction)
					g.GET("/path/del/:id", controller.AdminLoginLog.Del)
					g.POST("/post", controller.AdminLoginLog.Post)
					g.POST("/put", controller.AdminLoginLog.Put)
				})
				g.Group("/sys", func(g *ghttp.RouterGroup) {
					g.GET("/noticeAdmin", controller.Ws.NoticeAdmin)
					g.Middleware(sys.AuthAdmin)
					g.GET("/ws", controller.Ws.GetAdminWs)
				})
				//s.EnableHTTPS("./server.crt", "./server.key")
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
