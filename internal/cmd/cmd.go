package cmd

import (
	"ciel-admin/internal/controller"
	"ciel-admin/internal/service"
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
			service.System().Init()
			g.View().BindFuncMap(service.View().BindFuncMap())
			s := g.Server()
			s.Group("/", func(g *ghttp.RouterGroup) {
				g.GET("/", controller.Home.IndexPage)
				g.GET("/login", controller.Admin.LoginPage)
			})
			s.Group("/menu", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.Menu.List)
				g.GET("/getById", controller.Menu.GetById)
				g.DELETE("/del", controller.Menu.Del)
				g.POST("/post", controller.Menu.Post)
				g.PUT("/put", controller.Menu.Put)
			})
			s.Group("/api", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.Api.List)
				g.GET("/getById", controller.Api.GetById)
				g.DELETE("/del", controller.Api.Del)
				g.POST("/post", controller.Api.Post)
				g.PUT("/put", controller.Api.Put)
			})
			s.Group("/role", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.Role.List)
				g.GET("/getById", controller.Role.GetById)
				g.DELETE("/del", controller.Role.Del)
				g.GET("/nomenus", controller.RoleMenu.RoleNoMenus)
				g.GET("/noapis", controller.RoleMenu.RoleNoApis)
				g.POST("/post", controller.Role.Post)
				g.PUT("/put", controller.Role.Put)
			})
			s.Group("/roleApi", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.RoleApi.List)
				g.DELETE("/del", controller.RoleApi.Del)
				g.POST("/post", controller.RoleApi.Post)
			})
			s.Group("/roleMenu", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.RoleMenu.List)
				g.DELETE("/del", controller.RoleMenu.Del)
				g.POST("/post", controller.RoleMenu.Post)
			})
			s.Group("/admin", func(g *ghttp.RouterGroup) {
				g.POST("/login", controller.Admin.Login)
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/logout", controller.Admin.Logout)
				g.GET("/list", controller.Admin.List)
				g.GET("/getById", controller.Admin.GetById)
				g.DELETE("/del", controller.Admin.Del)
				g.POST("/post", controller.Admin.Post)
				g.PUT("/put", controller.Admin.Put)
				g.POST("/updatePwd", controller.Admin.UpdatePwd)
			})
			s.Group("/dict", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.Dict.List)
				g.GET("/getById", controller.Dict.GetById)
				g.DELETE("/del", controller.Dict.Del)
				g.POST("/post", controller.Dict.Post)
				g.PUT("/put", controller.Dict.Put)
			})
			s.Group("/file", func(g *ghttp.RouterGroup) {
				g.Middleware(service.Middleware().AuthAdmin)
				g.GET("/list", controller.File.List)
				g.GET("/getById", controller.File.GetById)
				g.DELETE("/del", controller.File.Del)
				g.POST("/post", controller.File.Post)
				g.PUT("/put", controller.File.Put)
				g.POST("/upload", controller.File.Upload)
			})
			go func() {
				var ctx = context.Background()
				time.Sleep(time.Second * 1)
				port, err := g.Cfg().Get(ctx, "server.address")
				if err != nil {
					panic(err)
				}
				glog.Infof(nil, "Server start at :http://localhost%s/login", port)
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
