package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/service/admin"
	"ciel-admin/internal/service/role"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xredis"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
	"time"
)

func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// AuthAdmin auth admin
func AuthAdmin(r *ghttp.Request) {
	user, err := admin.GetFromSession(r.Session)
	if err != nil || user == nil {
		r.Response.RedirectTo("/admin/login")
		return
	}
	if !role.CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI) {
		switch r.Method {
		case "GET", "DELETE", "POST":
			r.Session.Set("msg", fmt.Sprintf(consts.MsgWarning, "权限不足"))
			r.Response.RedirectTo(g.Config().MustGet(r.Context(), "home").String())
			r.Exit()
		default:
			res.Err(fmt.Errorf("权限不足"), r)
		}
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
		getAdmin, err := admin.GetFromSession(r.Session)
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
	user, err := admin.GetFromSession(r.Session)
	if err != nil || user == nil {
		res.Err(fmt.Errorf("用户信息错误"), r)
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
	g.Log().Info(ctx, uri)
	if uri == "/admin/operationLog/clear" {
		r.Middleware.Next()
		return
	}

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

// MiddlewareWhiteIp white ip
func MiddlewareWhiteIp(r *ghttp.Request) {
	ips := consts.WhiteIps
	if ips != "" {
		if !gstr.Contains(consts.WhiteIps, r.GetClientIp()) {
			r.Response.WriteStatus(http.StatusForbidden, fmt.Sprintf("%s ip error", r.GetClientIp()))
			r.Exit()
		}
	}
	r.Middleware.Next()
}
