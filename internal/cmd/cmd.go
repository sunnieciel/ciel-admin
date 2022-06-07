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
				g.Middleware(sys.LockAction)
				g.DELETE("/:id", controller.Menu.Del)
				g.POST("/", controller.Menu.Post)
				g.PUT("/", controller.Menu.Put)
			})
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Api.Path)
				g.GET("/", controller.Api.List)
				g.GET("/:id", controller.Api.GetById)
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
				g.DELETE("/:id", controller.Role.Del)
				g.GET("/nomenus", controller.RoleMenu.RoleNoMenus)
				g.GET("/noapis", controller.RoleMenu.RoleNoApis)
				g.GET("/currentMenus", controller.RoleMenu.CurrentMenus)
				g.POST("/", controller.Role.Post)
				g.PUT("/", controller.Role.Put)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleApi.Path)
				g.GET("/list", controller.RoleApi.List)
				g.DELETE("/del", controller.RoleApi.Del)
				g.POST("/post", controller.RoleApi.Post)
			})
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.RoleMenu.Path)
				g.GET("/list", controller.RoleMenu.List)
				g.DELETE("/del", controller.RoleMenu.Del)
				g.POST("/post", controller.RoleMenu.Post)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.POST("/login", controller.Admin.Login)
				g.Middleware(sys.AuthAdmin)
				g.GET("/logout", controller.Admin.Logout)
				g.GET("/path", controller.Admin.Path)
				g.GET("/", controller.Admin.List)
				g.GET("/:id", controller.Admin.GetById)
				g.DELETE("/:id", controller.Admin.Del)
				g.POST("/", controller.Admin.Post)
				g.PUT("/", controller.Admin.Put)
				g.POST("/updatePwd", controller.Admin.UpdatePwd)
			})
			s.Group("/dict", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Dict.Path)
				g.GET("/list", controller.Dict.List)
				g.GET("/getById", controller.Dict.GetById)
				g.DELETE("/del", controller.Dict.Del)
				g.POST("/post", controller.Dict.Post)
				g.PUT("/put", controller.Dict.Put)
			})
			s.Group("/file", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.File.Path)
				g.GET("/list", controller.File.List)
				g.GET("/getById", controller.File.GetById)
				g.DELETE("/del", controller.File.Del)
				g.POST("/post", controller.File.Post)
				g.PUT("/put", controller.File.Put)
				g.POST("/upload", controller.File.Upload)
			})
			s.Group("/gen", func(g *ghttp.RouterGroup) {
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Gen.Path)
				g.GET("/tables", controller.Gen.Tables)
				g.GET("/fields", controller.Gen.Fields)
				g.POST("/", controller.Gen.GenFile)
			})
			s.Group("/sys", func(g *ghttp.RouterGroup) {
				g.GET("/ws", controller.Ws.GetAdminWs)
				g.GET("/noticeAdmin", controller.Ws.NoticeAdmin)
				g.Middleware(sys.AuthAdmin)
				g.GET("/path", controller.Sys.Path)
				g.GET("/path/github", controller.Sys.PathGithub)
				g.GET("/path/oschina", controller.Sys.OsChina)
				g.GET("/path/douban", controller.Sys.Douban)
			})
			s.Group("/rss", func(g *ghttp.RouterGroup) {
				g.GET("/v2ex", controller.Rss.V2ex)
				g.GET("/fetch", controller.Rss.Fetch)
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
