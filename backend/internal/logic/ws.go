package logic

import (
	"context"
	"encoding/json"
	"errors"
	"freekey-backend/utility/utils/res"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Ws     = lWs{}
	users  = gmap.New(true)
	admins = gmap.New(true)
)

type lWs struct{}

func (l lWs) GetAdminWs(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	ws, err := r.WebSocket()
	if err != nil {
		res.Err(err, r)
	}
	uid := Sys.GetAdminIdFromCtx(ctx)
	if uid == 0 {
		res.Err(errors.New("链接失败，获取UID为空"), r)
	}

	value := admins.Get(uid)
	if value != nil {
		if err = value.(*ghttp.WebSocket).Close(); err != nil {
			g.Log().Error(r.Context(), err)
		}
		admins.Remove(uid)
	}
	admins.Set(uid, ws)
	l.printAdminWs()
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			if err = ws.Close(); err != nil {
				g.Log().Error(ctx, err)
			}
			admins.Remove(uid)
			l.printAdminWs()
			return
		}
		g.Log().Info(gctx.New(), "ws:lSys msg ", messageType, msg)
	}
}
func (l lWs) NoticeAdmin(ctx context.Context, msg interface{}, uid uint64) error {
	to := admins.Get(uid)
	if to != nil {
		marshal, _ := json.Marshal(msg)
		if err := to.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) NoticeAdmins(ctx context.Context, msg interface{}) error {
	marshal, _ := json.Marshal(msg)
	for _, id := range admins.Keys() {
		if err := admins.Get(id).(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) GetUserWs(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		res.Err(err, r)
	}
	uid := Biz.GetUserIdFromCtx(r.Context())
	users.Set(uid, ws)
	l.printUserWs()
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			users.Remove(uid)
			l.printUserWs()
			return
		}
		g.Log().Info(gctx.New(), "ws:lSys msg ", messageType, msg)
	}
}
func (l lWs) NoticeUser(ctx context.Context, uid int, msg interface{}) error {
	marshal, _ := json.Marshal(msg)
	item := users.Get(uid)
	if item != nil {
		if err := item.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) NoticeUsers(ctx context.Context, msg interface{}) error {
	if users.Size() == 0 {
		return nil
	}
	marshal, _ := json.Marshal(msg)
	for _, item := range users.Values() {
		if err := item.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}

func (l lWs) printUserWs() {
	g.Log().Infof(gctx.New(), "user连接个数%v %v", len(users.Map()), users.Keys())
}
func (l lWs) printAdminWs() {
	g.Log().Infof(gctx.New(), "admin连接个数%v %v", len(admins.Map()), admins.Keys())
}
