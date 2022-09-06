package cmd

import (
	"ciel-admin/internal/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

func registerGenFileRouter(s *ghttp.RouterGroup) {
	controller.User.RegisterRouter(s)              // 用户
	controller.UserLoginLog.RegisterRouter(s)      // 登录日志
	controller.Gold.RegisterRouter(s)              // 金币
	controller.TopUpCategory.RegisterRouter(s)     // 充值
	controller.GoldChangeLog.RegisterRouter(s)     // 账变日志
	controller.GoldStatisticsLog.RegisterRouter(s) // 账变统计
	controller.GoldReport.RegisterRouter(s)        // 账变报表
}
