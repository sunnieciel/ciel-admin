package ws

import (
	"ciel-admin/internal/logic"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

func GetUserWs(r *ghttp.Request) {
	logic.Ws.GetUserWs(r)
}
func GetAdminWs(r *ghttp.Request) {
	logic.Ws.GetAdminWs(r)
}
func NoticeAllUser(ctx context.Context, msg interface{}) error {
	return logic.Ws.NoticeAllUser(ctx, msg)
}
func NoticeAdmin(ctx context.Context, msg interface{}, toUid int) error {
	return logic.Ws.NoticeAdmin(ctx, msg, toUid)
}
func NoticeAllAdmin(ctx context.Context, msg interface{}) error {
	return logic.Ws.NoticeAllAdmin(ctx, msg)
}
func NoticeUser(ctx context.Context, uid int, msg interface{}) error {
	return logic.Ws.NoticeUser(ctx, uid, msg)
}
