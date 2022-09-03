package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/do"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xcaptcha"
	"ciel-admin/utility/utils/xpwd"
	"ciel-admin/utility/utils/xredis"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"time"
)

var (
	Admin = admin{}
)

type admin struct{}

func (admin) Menus(ctx context.Context, rid int, pid int) ([]*bo.Menu, error) {
	var d = make([]*bo.Menu, 0)
	menus, err := dao.RoleMenu.Menus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}

func (a admin) Login(ctx context.Context, id string, code string, uname string, pwd string, ip string) error {
	if !xcaptcha.Store.Verify(id, code, true) {
		return errors.New("验证码错误")
	}
	admin, err := dao.Admin.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(admin.Pwd, pwd) {
		return consts.ErrLogin
	}

	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	menus, err := Admin.Menus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	adminInfo := bo.Admin{Admin: admin, Menus: menus}
	if err = Session.SetAdmin(ctx, &adminInfo); err != nil {
		return err
	}
	if _, err = dao.AdminLoginLog.Ctx(ctx).Insert(do.AdminLoginLog{Uid: admin.Id, Ip: ip}); err != nil {
		return err
	}
	return nil
}

func (a admin) UpdatePwd(ctx context.Context, pwd string, pwd2 string) error {
	adminBo, err := Session.GetAdmin(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return err
	}
	u, err := dao.Admin.GetByUname(ctx, adminBo.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = Session.RemoveAdmin(ctx)
	if err != nil {
		return err
	}
	return dao.Admin.Update(ctx, u)
}

func (a admin) UpdateUname(ctx context.Context, id interface{}, uname interface{}) error {
	count, err := dao.Admin.Ctx(ctx).Count("uname", uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if _, err = dao.Admin.Ctx(ctx).Update(g.Map{"uname": uname}, "id", id); err != nil {
		return err
	}
	return nil
}

func (a admin) UpdatePwdWithoutOldPwd(ctx context.Context, id interface{}, pwd interface{}) error {
	_, err := dao.Admin.Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd.(string))}, "id", id)
	if err != nil {
		return err
	}
	return nil
}

func (a admin) ClearAllLoginLog(ctx context.Context) error {
	if _, err := dao.AdminLoginLog.Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}
	return nil
}

func (a admin) AuthMiddleware(r *ghttp.Request) {
	user, err := Session.GetAdmin(r.Session)
	if err != nil || user == nil {
		r.Response.RedirectTo("/admin/login")
		return
	}
	if !Role.CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI) {
		switch r.Method {
		case "GET", "DELETE", "POST":
			if err = r.Session.Set("msg", fmt.Sprintf(consts.MsgWarning, "权限不足")); err != nil {
				res.Err(err, r)
			}
			r.Response.RedirectTo(g.Config().MustGet(r.Context(), "home").String())
			r.Exit()
		default:
			res.Err(fmt.Errorf("权限不足"), r)
		}
	}
	r.Middleware.Next()
}

func (a admin) LockMiddleware(r *ghttp.Request) {
	var uid uint64
	getAdmin, err := Session.GetAdmin(r.Session)
	if err != nil {
		res.Err(err, r)
	}
	uid = uint64(getAdmin.Admin.Id)
	if uid == 0 {
		err := errors.New("uid is empty")
		g.Log().Error(nil, err)
		res.Err(err, r)
	}
	lock, err := xredis.UserLock(uid)
	if err != nil {
		res.Err(err, r)
	}
	r.Middleware.Next()
	lock.Unlock()
}

func (a admin) ActionMiddleware(r *ghttp.Request) {
	user, err := Session.GetAdmin(r.Session)
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

func (a admin) Add(ctx context.Context, in entity.Admin) error {
	if in.Pwd == "" {
		return consts.ErrPassEmpty
	}
	if in.Nickname == "" {
		in.Nickname = in.Uname
	}
	if in.Email != "" {
		if err := g.Validator().Rules("email").Data(in.Email).Run(ctx); err != nil {
			return consts.ErrFormatEmail
		}
	}
	count, err := dao.Admin.Ctx(ctx).Count("uname", in.Uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameAlreadyExist
	}
	in.Pwd = xpwd.GenPwd(in.Pwd)
	if _, err = dao.Admin.Ctx(ctx).Insert(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
