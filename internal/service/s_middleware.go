package service

import (
	"ciel-admin/internal/consts"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xjwt"
	"ciel-admin/utility/utils/xredis"
	"errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"net/http"
)

// ---sMiddleware-----------------------------------------------------------------------
type sMiddleware struct{}

var (
	insMiddleware = new(sMiddleware)
	Uid           = "userInfoKey"
)

func Middleware() *sMiddleware { return insMiddleware }
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (s *sMiddleware) AuthAdmin(r *ghttp.Request) {
	user, err := Session().GetAdmin(r)
	if err != nil || user == nil {
		r.Response.RedirectTo("/login")
		return
	}
	b := Role().CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI, r.Method)
	if !b {
		res.Err(errors.New("没有权限"), r)
	}
	r.Middleware.Next()
}
func (s *sMiddleware) UserAuth(c *ghttp.Request) {
	userInfo, err := xjwt.UserInfo(c)
	if err != nil {
		c.Response.WriteStatus(http.StatusForbidden, consts.ErrAuth.Error())
		c.Exit()
	}
	c.SetParam(Uid, userInfo.Uid)
	c.Middleware.Next()
}

func (s *sMiddleware) LockAction(r *ghttp.Request) {
	uid := r.Get(Uid).Uint64()
	if uid == 0 {
		getAdmin, err := Session().GetAdmin(r)
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
