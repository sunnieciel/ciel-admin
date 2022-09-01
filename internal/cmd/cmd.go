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
			s.EnableAdmin("/debut/admin")
			s.Group("/", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.Home.IndexPage)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.WhiteIpMiddleware) // 白名单过滤  在字典表中为空时，这里不会进行检查的
				registerGenFileRouter(g)            // 注册生成的代码路由
				g.Group("/menu", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.Menu.Index)             // 主页面
					g.GET("/add", controller.Menu.AddIndex)       // 添加页面
					g.GET("/edit/:id", controller.Menu.EditIndex) // 修改页面
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.Menu.Del) // 删除请求
					g.POST("/post", controller.Menu.Post)  // 添加请求
					g.POST("/put", controller.Menu.Put)    // 修改请求
				})
				g.Group("/api", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.Api.Index)
					g.GET("/add", controller.Api.AddIndex)
					g.GET("/edit/:id", controller.Api.EditIndex)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.Api.Del)
					g.POST("/post", controller.Api.Post)
					g.POST("/put", controller.Api.Put)
				})
				g.Group("/roleMenu", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.RoleMenu.Path)
					g.GET("/add", controller.RoleMenu.PathAdd)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.RoleMenu.Del)
					g.POST("/post", controller.RoleMenu.Post)
				})
				g.Group("/role", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.Role.Index)
					g.GET("/add", controller.Role.AddIndex)
					g.GET("/edit/:id", controller.Role.EditIndex)
					g.GET("/nomenus", controller.Role.RoleNoMenus)
					g.GET("/noapis", controller.Role.RoleNoApis)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.Role.Del)
					g.POST("/post", controller.Role.Post)
					g.POST("/put", controller.Role.Put)
				})
				g.Group("/roleApi", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.RoleApi.Index)
					g.GET("/add", controller.RoleApi.AddIndex)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.RoleApi.Del)
					g.GET("/clear/:rid", controller.RoleApi.Clear)
					g.POST("/post", controller.RoleApi.Post)
				})
				g.Group("/dict", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.Dict.Index)
					g.GET("/add", controller.Dict.AddIndex)
					g.GET("/edit/:id", controller.Dict.PathEdit)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.Dict.Del)
					g.POST("/post", controller.Dict.Post)
					g.POST("/put", controller.Dict.Put)
				})
				g.Group("/admin", func(g *ghttp.RouterGroup) {
					g.GET("/getCaptcha", controller.Sys.GetCaptcha) // 获取验证码
					g.POST("/login", controller.Admin.Login)
					g.Middleware(admin.AuthMiddleware)
					g.GET("/logout", controller.Admin.Logout)
					g.GET("/", controller.Admin.Index)
					g.GET("/add", controller.Admin.AddIndex)
					g.GET("/edit/:id", controller.Admin.EditIndex)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.PUT("/updatePwd", controller.Admin.UpdatePwd)
					g.PUT("/updatePwdWithoutOldPwd", controller.Admin.UpdatePwdWithoutOldPwd)
					g.PUT("/updateUname", controller.Admin.UpdateUname)
					g.GET("/del/:id", controller.Admin.Del)
					g.POST("/post", controller.Admin.Post)
					g.POST("/put", controller.Admin.Put)
				})
				g.Group("/operationLog", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.OperationLog.Index)
					g.GET("/add", controller.OperationLog.AddIndex)
					g.GET("/edit/:id", controller.OperationLog.EditIndex)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.OperationLog.Del)
					g.POST("/post", controller.OperationLog.Post)
					g.POST("/put", controller.OperationLog.Put)
					g.GET("/clear", controller.OperationLog.Clear)
				})
				g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/path", controller.AdminLoginLog.Path)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/path/del/:id", controller.AdminLoginLog.Del)
					g.GET("/clear", controller.AdminLoginLog.Clear)
				})
				g.Group("/file", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.File.Index)
					g.GET("/add", controller.File.AddIndex)
					g.GET("/edit/:id", controller.File.EditIndex)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.File.Del)
					g.POST("/post", controller.File.Post)
					g.POST("/put", controller.File.Put)
					g.POST("/upload", controller.File.Upload)
				})
				g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.AdminLoginLog.Path)
					g.GET("/add", controller.AdminLoginLog.PathAdd)
					g.GET("/edit/:id", controller.AdminLoginLog.PathEdit)
					g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
					g.GET("/del/:id", controller.AdminLoginLog.Del)
					g.POST("/post", controller.AdminLoginLog.Post)
					g.POST("/put", controller.AdminLoginLog.Put)
				})
				g.Group("/sys", func(g *ghttp.RouterGroup) {
					g.GET("/noticeAdmin", controller.Ws.NoticeAdmin)
					g.GET("/document", controller.Sys.DocumentIndex)
					g.Middleware(admin.AuthMiddleware)
					g.GET("/ws", controller.Ws.GetAdminWs)
				})
				g.Group("/gen", func(g *ghttp.RouterGroup) {
					g.Middleware(admin.AuthMiddleware)
					g.GET("/", controller.Gen.Index)
					g.POST("/table", controller.Gen.Gen)
				})
				g.GET("/login", controller.Admin.LoginPage)
				g.GET("/to/:path", controller.Sys.To)
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
