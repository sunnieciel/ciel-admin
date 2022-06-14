package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service/sys"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
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
			registerInterface(s)     // 注册对外提供功能的接口
			registerGenFileRouter(s) // 注册生成的代码路由

			s.Group("/", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.Home.IndexPage)
				g.GET("/login", controller.Admin.LoginPage)
			})
			s.Group("/menu", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Menu.Path)
				g.GET("/", controller.Menu.List)
				g.GET("/:id", controller.Menu.GetById)
				g.GET("/level1", controller.Menu.ListLevel1) // 获取一级菜单
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.Menu.Del)
				g.POST("/", controller.Menu.Post)
				g.PUT("/", controller.Menu.Put)
			})
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Api.Path)
				g.GET("/", controller.Api.List)
				g.GET("/:id", controller.Api.GetById)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.Api.Del)
				g.POST("/", controller.Api.Post)
				g.PUT("/", controller.Api.Put)
			})
			s.Group("/role", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Role.Path)
				g.GET("/", controller.Role.List)
				g.GET("/:id", controller.Role.GetById)
				g.GET("/roles", controller.Role.Roles) // select all role info
				g.GET("/nomenus", controller.RoleMenu.RoleNoMenus)
				g.GET("/noapis", controller.RoleMenu.RoleNoApis)
				g.GET("/currentMenus", controller.RoleMenu.CurrentMenus)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.Role.Del)
				g.POST("/", controller.Role.Post)
				g.PUT("/", controller.Role.Put)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleApi.Path)
				g.GET("/", controller.RoleApi.List)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.RoleApi.Del)
				g.POST("/", controller.RoleApi.Post)
			})
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleMenu.Path)
				g.GET("/", controller.RoleMenu.List)
				g.GET("/:id", controller.RoleMenu.GetById)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.RoleMenu.Del)
				g.POST("/", controller.RoleMenu.Post)
				g.PUT("/", controller.RoleMenu.Put)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.POST("/login", controller.Admin.Login)
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Admin.Path)
				g.GET("/", controller.Admin.List)
				g.GET("/:id", controller.Admin.GetById)
				g.DELETE("/:id", controller.Admin.Del)
				g.GET("/logout", controller.Admin.Logout)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.POST("/", controller.Admin.Post)
				g.PUT("/", controller.Admin.Put)
				g.PUT("/updatePwd", controller.Admin.UpdatePwd)
				g.PUT("/updatePwdWithoutOldPwd", controller.Admin.UpdatePwdWithoutOldPwd)
				g.PUT("/updateUname", controller.Admin.UpdateUname)
			})
			s.Group("/dict", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Dict.Path)
				g.GET("/", controller.Dict.List)
				g.GET("/:id", controller.Dict.GetById)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.Dict.Del)
				g.POST("/", controller.Dict.Post)
				g.PUT("/", controller.Dict.Put)
			})
			s.Group("/file", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.File.Path)
				g.GET("/", controller.File.List)
				g.GET("/:id", controller.File.GetById)
				g.Middleware(sys.LockAction)
				g.DELETE("/:id", controller.File.Del)
				g.POST("/", controller.File.Post)
				g.PUT("/", controller.File.Put)
				g.Middleware(sys.AdminAction)
				g.POST("/upload", controller.File.Upload)
			})
			s.Group("/operationLog", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.OperationLog.Path)
				g.GET("/", controller.OperationLog.List)
				g.GET("/:id", controller.OperationLog.GetById)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/:id", controller.OperationLog.Del)
				g.POST("/", controller.OperationLog.Post)
				g.PUT("/", controller.OperationLog.Put)
			})
			s.Group("/gen", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Gen.Path)
				g.GET("/tables", controller.Gen.Tables)
				g.GET("/fields", controller.Gen.Fields)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.POST("/", controller.Gen.GenFile)
			})

			s.Group("/loginLog", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.LoginLog.Path)
				g.GET("/", controller.LoginLog.List)
				g.GET("/:id", controller.LoginLog.GetById)
				g.Middleware(sys.LockAction)
				g.DELETE("/:id", controller.LoginLog.Del)
				g.POST("/", controller.LoginLog.Post)
				g.PUT("/", controller.LoginLog.Put)
			})
			s.Group("/sys", func(g *ghttp.RouterGroup) {
				g.GET("/ws", controller.Ws.GetAdminWs)
				g.GET("/noticeAdmin", controller.Ws.NoticeAdmin)
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Sys.Path)
			})
			go func() {
				var ctx = context.Background()
				time.Sleep(time.Second * 1)
				port, err := g.Cfg().Get(ctx, "server.address")
				if err != nil {
					panic(err)
				}
				rootIp, err := g.Cfg().Get(ctx, "server.rootIp")
				glog.Infof(nil, "Server start at :http://%s%s/login", rootIp, port)
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
