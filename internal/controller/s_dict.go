package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/dict"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// ---Dict-----------------------------------------------------------------

type cDict struct{ cBase }

var Dict = &cDict{
	cBase{Table: "s_dict", ReqPath: "/admin/dict", FileDir: "/sys/dict"},
}

func (c cDict) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = bo.Search{
			T1: "s_dict", OrderBy: "t1.group,t1.id desc", SearchFields: "t1.*", Fields: []bo.Field{
				{Name: "k", Type: 2}, {Name: "v", Type: 2}, {Name: "desc", Type: 2}, {Name: "group", Type: 1}, {Name: "status", Type: 1}, {Name: "type", Type: 1},
			},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
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
		"path": path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cDict) AddIndex(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cDict) Post(r *ghttp.Request) {
	var (
		d     = entity.Dict{}
		ctx   = r.Context()
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
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
func (c cDict) Del(r *ghttp.Request) {
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
func (c cDict) PathEdit(r *ghttp.Request) {
	var (
		table = c.Table
		ctx   = r.Context()
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
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
func (c cDict) Put(r *ghttp.Request) {
	var (
		d     = entity.Dict{}
		table = c.Table
		ctx   = r.Context()
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
	if d.K == "white_ips" {
		if err := dict.SetWhiteIps(r.Context(), d.V); err != nil {
			res.Err(err, r)
		}
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
