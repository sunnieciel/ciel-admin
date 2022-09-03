package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/role"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type cRole struct{ cBase }

var Role = &cRole{cBase{"s_role", "/admin/role", "/sys/role"}}

func (c cRole) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		s       = bo.Search{T1: "s_role", OrderBy: "t1.id desc", SearchFields: "t1.*"}
		reqPath = r.URL.Path
	)
	node, err := sys.NodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(ctx, s)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/index.html"), g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": reqPath,
	}, r)
}
func (c cRole) AddIndex(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": sys.MsgFromSession(r)}, r)
}

func (c cRole) EditIndex(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		file  = fmt.Sprint(c.FileDir, "/edit.html")
		d     = g.Map{"msg": sys.MsgFromSession(r)}
		id    = xparam.ID(r)
	)
	data, err := sys.GetById(ctx, table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(file, d, r)
}

func (c cRole) Post(r *ghttp.Request) {
	var (
		d     = entity.Role{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := sys.Add(ctx, table, &d); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cRole) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
		ctx   = r.Context()
		path  = fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cRole) Put(r *ghttp.Request) {
	var (
		d     = entity.Role{}
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
	path := fmt.Sprintf("%s/edit/%v?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cRole) RoleNoMenus(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid")
		ctx = r.Context()
	)
	data, err := role.NoMenu(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c cRole) RoleNoApis(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid")
		ctx = r.Context()
	)
	data, err := role.NoApi(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

func (c cRole) Clear(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid").Int()
		ctx = r.Context()
	)
	res.OkSession("清除成功", r)
	if err := role.Clear(ctx, rid); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo("", r)
}
