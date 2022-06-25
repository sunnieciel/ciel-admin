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
			registerInterface(s)     // 注册对外提供功能的接口
			registerGenFileRouter(s) // 注册生成的代码路由

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.GET("/", controller.Home.IndexPage)
				group.GET("/login", controller.Admin.LoginPage)
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
				g.DELETE("/batch", controller.Role.Del)
				g.POST("/", controller.Role.Post)
				g.PUT("/", controller.Role.Put)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleApi.Path)
				g.GET("/", controller.RoleApi.List)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/batch", controller.RoleApi.Del)
				g.POST("/", controller.RoleApi.Post)
			})
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleMenu.Path)
				g.GET("/", controller.RoleMenu.List)
				g.GET("/:id", controller.RoleMenu.GetById)
				g.Middleware(sys.LockAction, sys.AdminAction)
				g.DELETE("/batch", controller.RoleMenu.Del)
				g.POST("/", controller.RoleMenu.Post)
				g.PUT("/", controller.RoleMenu.Put)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.POST("/login", controller.Admin.Login)
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Admin.Path)
				g.GET("/", controller.Admin.List)
				g.GET("/:id", controller.Admin.GetById)
				g.DELETE("/batch", controller.Admin.Del)
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
				g.DELETE("/batch", controller.Dict.Del)
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
				g.DELETE("/batch", controller.OperationLog.Del)
				g.POST("/", controller.OperationLog.Post)
				g.PUT("/", controller.OperationLog.Put)
			})
			s.Group("/gen", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Gen.Path)
				g.GET("/tables", controller.Gen.Tables)
				g.GET("/fields", controller.Gen.Fields)
				g.Middleware(sys.LockAction)
				g.POST("/", controller.Gen.GenFile)
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
