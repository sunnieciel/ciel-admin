package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type cAdminLoginLog struct{ cBase }

var AdminLoginLog = &cAdminLoginLog{cBase{Table: "s_admin_login_log", ReqPath: "/admin/adminLoginLog", FileDir: "/sys/adminLoginLog"}}

func (c cAdminLoginLog) Path(r *ghttp.Request) {
	var (
		s = bo.Search{
			T1: "s_admin_login_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
			Fields: []bo.Field{
				{Name: "uid", Type: 1}, {Name: "t2.uname", Type: 2, QueryName: "uname"},
			},
		}
		ctx  = r.Context()
		path = r.URL.Path
		file = fmt.Sprintf("%s/index.html", c.FileDir)
	)
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(ctx, s)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cAdminLoginLog) PathAdd(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cAdminLoginLog) Post(r *ghttp.Request) {
	var (
		d     = entity.AdminLoginLog{}
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := sys.Add(ctx, table, &d); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdminLoginLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
		ctx   = r.Context()
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdminLoginLog) PathEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
		table = c.Table
		id    = xparam.ID(r)
	)
	data, err := sys.GetById(ctx, table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cAdminLoginLog) Put(r *ghttp.Request) {
	var (
		d     = entity.AdminLoginLog{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := sys.Update(ctx, table, d.Id, m); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

func (c cAdminLoginLog) Clear(r *ghttp.Request) {
	var (
		path = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("操作成功", r)
	if err := sys.ClearAdminLog(r.Context()); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
