package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"time"
)

var (
	Main = gcmd.Command{
		Name:        "main",
		Usage:       "main",
		Brief:       "start http server",
		Description: "",
		Arguments:   nil,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化服务
			sys.Init()
			g.View().BindFuncMap(sys.BindFuncMap())
			s := g.Server()
			registerInterface(s)                         // 注册对外提供功能的接口
			registerGenFileRouter(s)                     // 注册生成的代码路由
			s.BindMiddlewareDefault(sys.MiddlewareXIcon) // 默认中间件
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.GET("/", controller.Home.IndexPage)
				group.GET("/login", controller.Admin.LoginPage)
				group.Middleware(sys.AuthAdmin)
				group.GET("/to/:name", controller.Sys.To)
			})
			s.Group("/menu", func(g *ghttp.RouterGroup) {
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
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Api.Path)
				g.GET("/path/add", controller.Api.PathAdd)
				g.GET("/path/edit/:id", controller.Api.PathEdit)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.GET("/path/del/:id", controller.Api.Del)
				g.POST("/post", controller.Api.Post)
				g.POST("/put", controller.Api.Put)
			})
			s.Group("/dict", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Dict.Path)
				g.GET("/path/add", controller.Dict.PathAdd)
				g.GET("/path/edit/:id", controller.Dict.PathEdit)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.GET("/path/del/:id", controller.Dict.Del)
				g.POST("/post", controller.Dict.Post)
				g.POST("/put", controller.Dict.Put)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
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
			s.Group("/role", func(g *ghttp.RouterGroup) {
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
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleMenu.Path)
				g.GET("/path/add", controller.RoleMenu.PathAdd)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.GET("/path/del/:id", controller.RoleMenu.Del)
				g.POST("/post", controller.RoleMenu.Post)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleApi.Path)
				g.GET("/path/add", controller.RoleApi.PathAdd)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.GET("/path/del/:id", controller.RoleApi.Del)
				g.GET("/clear/:rid", controller.RoleApi.Clear)
				g.POST("/post", controller.RoleApi.Post)
			})
			s.Group("/file", func(g *ghttp.RouterGroup) {
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
			s.Group("/operationLog", func(g *ghttp.RouterGroup) {
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
			s.Group("/gen", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Gen.Path)
				g.GET("/tables", controller.Gen.Tables)
				g.GET("/fields", controller.Gen.Fields)
				g.Middleware(sys.LockAction)
				g.POST("/", controller.Gen.GenFile)
			})
			go func() {
				var ctx = context.Background()
				time.Sleep(time.Second * 1)
				port, err := g.Cfg().Get(ctx, "server.address")
				if err != nil {
					panic(err)
				}
				rootIp, err := g.Cfg().Get(ctx, "server.rootIp")
				g.Log().Infof(nil, "Server start at :http://%s%s/login", rootIp, port)
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
