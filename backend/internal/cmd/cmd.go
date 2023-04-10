package cmd

import (
	"context"
	"freekey-backend/internal/controller"
	"freekey-backend/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			service.Sys.Init(ctx)
			s := g.Server()
			routers(s)
			s.Run()
			return nil
		},
	}
)

func routers(s *ghttp.Server) {
	s.Group("/v1", func(g *ghttp.RouterGroup) { v1Rooters(g) })
	s.Group("/backend", func(g *ghttp.RouterGroup) { sysRouters(g) })
}

func v1Rooters(g *ghttp.RouterGroup) {
	g.Middleware(service.Sys.MiddlewareHandlerResponse, service.Sys.MiddlewareCORS, service.Sys.MiddleIpRateLimit)
	g.Group("/sys", func(g *ghttp.RouterGroup) {
		g.GET("/dict", controller.Sys.GetDictByKey)
		g.POST("/uploadFile", controller.Sys.UploadFiles)
	})
	g.Group("/user", func(g *ghttp.RouterGroup) {
		g.POST("/register", controller.Biz.Register)
		g.POST("/login", controller.Biz.Login)
		g.Middleware(service.Biz.MiddlewareUserAuth)
		g.GET("/getUserInfo", controller.Biz.GetUserInfo)
		g.PUT("/pass", controller.Biz.UpdateUserPass)
		g.PUT("/nickname", controller.Biz.UpdateUserNickname)
		g.PUT("/icon", controller.Biz.UpdateUserIcon)
	})
	g.Group("/wallet", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Biz.MiddlewareUserAuth)
		g.POST("/topUp", controller.Biz.CreateTopUp)
		g.GET("/listTopUp", controller.Biz.ListTopUpForWeb)
		g.GET("/listChangeLog", controller.Biz.ListWalletChangeLogForWeb)
	})
}
func sysRouters(g *ghttp.RouterGroup) {
	g.Middleware(
		service.Sys.MiddlewareCORS,
		service.Sys.MiddlewareAdminActionLog,
		//service.Sys.MiddlewareWhiteIp, // ip白名单检测
	)
	g.Group("/ws", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Sys.MiddlewareAdminAuth)
		g.GET("/", controller.Sys.WsGetConnectForAdmin) // 管理员连接
		g.POST("/sendMsg", controller.Sys.WsSendMsg)
		g.POST("/noticeAdmins", controller.Sys.WsNoticeAdmins)
	})
	g.Middleware(ghttp.MiddlewareHandlerResponse)
	g.GET("/getCaptcha", controller.Sys.GetCaptcha)
	g.POST("/login", controller.Sys.AdminLogin)
	g.Middleware(service.Sys.MiddlewareAdminAuth)
	g.Group("/menu", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetMenuById)
		g.GET("/list", controller.Sys.ListMenu)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.POST("/", controller.Sys.AddMenu)
		g.DELETE("/", controller.Sys.DelMenu)
		g.PUT("/", controller.Sys.UpdateMenu)
		g.PUT("/sort", controller.Sys.SortMenu)
	})
	g.Group("/api", func(g *ghttp.RouterGroup) {
		g.POST("/", controller.Sys.AddAPI)
		g.POST("/addGroup", controller.Sys.AddAPIGroup)
		g.GET("/", controller.Sys.GetAPIById)
		g.GET("/list", controller.Sys.ListAPI)
		g.DELETE("/", controller.Sys.DelAPI)
		g.PUT("/", controller.Sys.UpdateAPI)
	})
	g.Group("/role", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetRoleById)
		g.GET("/list", controller.Sys.ListRole)
		g.GET("/getOptions", controller.Sys.GetOptions)
		g.POST("/", controller.Sys.AddRole)
		g.DELETE("/", controller.Sys.DelRole)
		g.PUT("/", controller.Sys.UpdateRole)
	})
	g.Group("/roleMenu", func(g *ghttp.RouterGroup) {
		g.GET("/list", controller.Sys.ListRoleMenu)
		g.GET("/listRoleNoMenus", controller.Sys.ListRoleNoMenus)
		g.POST("/addRoleMenus", controller.Sys.AddRoleMenus)
		g.DELETE("/", controller.Sys.DelRoleMenu)
		g.DELETE("/clear", controller.Sys.ClearRoleMenu)
	})
	g.Group("/roleApi", func(g *ghttp.RouterGroup) {
		g.GET("/list", controller.Sys.ListRoleApi)
		g.GET("/listRoleNoApis", controller.Sys.ListRoleNoApis)
		g.POST("/addRoleApis", controller.Sys.AddRoleApis)
		g.DELETE("/", controller.Sys.DelRoleApi)
		g.DELETE("/clear", controller.Sys.ClearRoleApi)
	})
	g.Group("/admin", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetAdminById)
		g.GET("/list", controller.Sys.ListAdmin)
		g.GET("/getInfo", controller.Sys.GetAdminInfo)
		g.GET("/getMenu", controller.Sys.GetMenuByPath)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.POST("/", controller.Sys.AddAdmin)
		g.DELETE("/", controller.Sys.DelAdmin)
		g.PUT("/", controller.Sys.UpdateAdmin)
		g.PUT("/updateUname", controller.Sys.UpdateAdminUname)
		g.PUT("/updatePass", controller.Sys.UpdateAdminPass)
		g.PUT("/updateSelfPass", controller.Sys.UpdateAdminPassBySelf)
	})
	g.Group("/dict", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetDictById)
		g.GET("/list", controller.Sys.ListDict)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.POST("/", controller.Sys.AddDict)
		g.DELETE("/", controller.Sys.DelDict)
		g.PUT("/", controller.Sys.UpdateDict)
	})
	g.Group("/operationLog", func(g *ghttp.RouterGroup) {
		g.GET("/list", controller.Sys.ListOperationLog)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Sys.DelOperationLog)
		g.DELETE("/delClear", controller.Sys.ClearOperationLog)
	})
	g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {
		g.GET("/list", controller.Sys.ListLoginLog)
		g.DELETE("/", controller.Sys.DelLoginLog)
		g.DELETE("/delClear", controller.Sys.ClearLoginLog)
	})
	g.Group("/file", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetFileById)
		g.GET("/list", controller.Sys.ListFile)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Sys.DelFile)
		g.PUT("/", controller.Sys.UpdateFile)
		g.POST("/upload", controller.Sys.UploadFiles)
	})
	g.Group("/banner", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Sys.GetBannerById)
		g.GET("/list", controller.Sys.ListBanner)
		g.POST("/", controller.Sys.AddBanner)
		g.DELETE("/", controller.Sys.DelBanner)
		g.PUT("/", controller.Sys.UpdateBanner)
	})
	g.Group("/user", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetUserById)
		g.GET("/list", controller.Biz.ListUser)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Biz.DelUser)
		g.PUT("/", controller.Biz.UpdateUser)
		g.PUT("/updateUname", controller.Biz.UpdateUserUname)
		g.PUT("/updatePass", controller.Biz.UpdateUserPassByAdmin)
	})
	g.Group("/userLoginLog", func(g *ghttp.RouterGroup) {
		g.GET("/list", controller.Biz.ListUserLoginLog)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Biz.DelUserLoginLog)
		g.DELETE("/delClear", controller.Biz.ClearUserLoginLog)
	})
	g.Group("/wallet", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetWalletById)
		g.GET("/list", controller.Biz.ListWallet)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.PUT("/", controller.Biz.UpdateWallet)
		g.PUT("/updatePass", controller.Biz.UpdateWalletPassByAdmin)
		g.PUT("/updateByAdmin", controller.Biz.UpdateWalletMoneyByAdmin)
		g.GET("report", controller.Biz.GetWalletReport)
	})
	g.Group("/walletChangeType", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetWalletChangeTypeById)
		g.GET("/list", controller.Biz.ListWalletChangeType)
		g.GET("/listOptions", controller.Biz.ListWalletChangeTypeOptions)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.POST("/", controller.Biz.AddWalletChangeType)
		g.DELETE("/", controller.Biz.DelWalletChangeType)
		g.PUT("/", controller.Biz.UpdateWalletChangeType)
	})
	g.Group("/walletChangeLog", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetWalletChangeLogById)
		g.GET("/list", controller.Biz.ListWalletChangeLog)
		g.POST("/", controller.Biz.AddWalletChangeLog)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Biz.DelWalletChangeLog)
		g.PUT("/", controller.Biz.UpdateWalletChangeLog)
	})
	g.Group("/walletStatisticsLog", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetWalletStatisticsLogById)
		g.GET("/list", controller.Biz.ListWalletStatisticsLog)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.POST("/", controller.Biz.AddWalletStatisticsLog)
		g.DELETE("/", controller.Biz.DelWalletStatisticsLog)
		g.PUT("/", controller.Biz.UpdateWalletStatisticsLog)
	})
	// --- TopUp -----------------------------------------------------------------
	g.Group("/topUp", func(g *ghttp.RouterGroup) {
		g.GET("/", controller.Biz.GetTopUpById)
		g.GET("/list", controller.Biz.ListTopUp)
		g.Middleware(service.Sys.MiddlewareActionLock)
		g.DELETE("/", controller.Biz.DelTopUp)
		g.PUT("/", controller.Biz.UpdateTopUp)
		g.PUT("/updateByAdmin", controller.Biz.UpdateTopUpByAdmin)
	})

}
