package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xpwd"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

//  ---admin-------------------------------------------------------------------

type cAdmin struct{ cBase }

var Admin = &cAdmin{
	cBase{Table: "s_admin", ReqPath: "/admin/admin", FileDir: "/sys/admin/"},
}

func (c cAdmin) Index(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		s     = bo.Search{
			T1: table, T2: "s_role t2 on t1.rid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.name role_name",
			Fields: []bo.Field{{Name: "rid", Type: 1}, {Name: "status", Type: 1}},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
	)
	roles, err := sys.Roles(ctx)
	if err != nil {
		res.Err(err, r)
	}
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(ctx, s)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list":  data,
		"page":  r.GetPage(total, s.Size).GetContent(3),
		"node":  node,
		"msg":   sys.MsgFromSession(r),
		"roles": roles,
		"path":  r.URL.Path,
	}, r)
}
func (c cAdmin) AddIndex(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	roles, err := sys.Roles(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r), "roles": roles}, r)
}
func (c cAdmin) Post(r *ghttp.Request) {
	var (
		d     = entity.Admin{}
		ctx   = r.Context()
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	d.Pwd = xpwd.GenPwd(d.Pwd)
	res.OkSession("添加成功", r)
	if err := sys.Add(ctx, table, &d); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdmin) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdmin) EditIndex(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
		id    = xparam.ID(r)
	)
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	data, err := sys.GetById(ctx, table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r), "roles": roles}, r)
}
func (c cAdmin) Put(r *ghttp.Request) {
	var (
		d     = entity.Admin{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "pwd")
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := sys.Update(ctx, table, d.Id, m); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cAdmin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c cAdmin) Login(r *ghttp.Request) {
	var d struct {
		Uname string `form:"uname"`
		Pwd   string `form:"pwd"`
		Id    string `form:"id"`   // 获取二维码时的id
		Code  string `from:"code"` // 二维码
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Login(r.Context(), d.Id, d.Code, d.Uname, d.Pwd, r.GetClientIp()); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) Logout(r *ghttp.Request) {
	err := sys.Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		OldPwd string `v:"required"`
		NewPwd string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwd(r.Context(), d.OldPwd, d.NewPwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) UpdateUname(r *ghttp.Request) {
	var d struct {
		Uname string `v:"required"`
		Id    int64  `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminUname(r.Context(), d.Id, d.Uname); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) UpdatePwdWithoutOldPwd(r *ghttp.Request) {
	var d struct {
		Pwd string `v:"required"`
		Id  string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwdWithoutOldPwd(r.Context(), d.Id, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
