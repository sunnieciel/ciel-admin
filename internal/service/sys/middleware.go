package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xredis"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
	"time"
)

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func AuthAdmin(r *ghttp.Request) {
	user, err := GetAdmin(r)
	if err != nil || user == nil {
		r.Response.RedirectTo("/login")
		return
	}
	b := CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI, r.Method)
	if !b {
		res.Err(errors.New("没有权限"), r)
	}
	r.Middleware.Next()
}
func UserAuth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
}
func LockAction(r *ghttp.Request) {
	uid := r.Get(Uid).Uint64()
	if uid == 0 {
		getAdmin, err := GetAdmin(r)
		if err != nil {
			res.Err(err, r)
		}
		uid = uint64(getAdmin.Admin.Id)
		if uid == 0 {
			err := errors.New("uid is empty")
			glog.Error(nil, err)
			res.Err(err, r)
		}
	}
	lock, err := xredis.UserLock(uid)
	if err != nil {
		res.Err(err, r)
	}
	r.Middleware.Next()
	lock.Unlock()
}
func AdminAction(r *ghttp.Request) {
	user, err := GetAdmin(r)
	if err != nil || user == nil {
		r.Response.RedirectTo("/login")
		return
	}
	uid := user.Admin.Id
	content := ""
	method := r.Method
	ctx := r.Context()
	uri := r.Router.Uri
	ip := r.GetClientIp()
	begin := time.Now().UnixMilli()

	switch method {
	case "GET", "DELETE":
		content = r.GetUrl()
	case "POST", "PUT":
		content = r.GetBodyString()
		if content == "" {
			content = r.Request.PostForm.Encode()
		}
		if content == "" {
			content = r.Request.Form.Encode()
		}
	}
	r.Middleware.Next()
	response := r.Response.BufferString()
	useTime := time.Now().UnixMilli() - begin
	data := g.Map{
		"uid":      uid,
		"content":  content,
		"method":   method,
		"uri":      uri,
		"response": response,
		"use_time": useTime,
		"ip":       ip,
	}
	_, err = dao.OperationLog.Ctx(ctx).Insert(data)
	if err != nil {
		glog.Error(ctx, err)
	}
}
