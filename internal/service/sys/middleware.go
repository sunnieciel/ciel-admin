package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xredis"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
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
			g.Log().Error(nil, err)
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
	response := ""

	switch method {
	case "GET":
		content = r.GetUrl()
	case "DELETE":
		content = fmt.Sprintf("删除记录ID %s", r.Get("id").String())
	case "POST", "PUT":
		content = fmt.Sprint(r.GetFormMap())
		if content == "" {
			content = r.Request.PostForm.Encode()
		}
		if content == "" {
			content = r.Request.Form.Encode()
		}
	}
	r.Middleware.Next()
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
		g.Log().Error(ctx, err)
	}
}
func MiddlewareXIcon(r *ghttp.Request) {
	if r.GetHeader("Sec-Fetch-Dest") == "image" {
		r.Response.Write("ok")
		r.Exit()
	}
	r.Middleware.Next()
}
