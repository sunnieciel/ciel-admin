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

type cOperationLog struct{ cBase }

var OperationLog = &cOperationLog{cBase{Table: "s_operation_log", ReqPath: "/admin/operationLog", FileDir: "/sys/operationLog"}}

func (c cOperationLog) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = bo.Search{
			T1: "s_operation_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
			Fields: []bo.Field{
				{Name: "t2.uname", Type: 2, QueryName: "uname"}, {Name: "content", Type: 2},
			},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
		msg  = sys.MsgFromSession(r)
	)
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), s)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  msg,
		"path": path,
	}, r)
}
func (c cOperationLog) AddIndex(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
		d    = g.Map{"msg": sys.MsgFromSession(r)}
	)
	res.Tpl(file, d, r)
}
func (c cOperationLog) EditIndex(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		id    = xparam.ID(r)
		table = c.Table
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
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
func (c cOperationLog) Post(r *ghttp.Request) {
	var (
		d     = entity.OperationLog{}
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
func (c cOperationLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		ctx   = r.Context()
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cOperationLog) Put(r *ghttp.Request) {
	var (
		d     = entity.OperationLog{}
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
	path := fmt.Sprintf("%s/edit/%d/?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cOperationLog) Clear(r *ghttp.Request) {
	res.OkSession("操作成功", r)
	if err := sys.OperationLogClear(r.Context()); err != nil {
		res.ErrSession(err, r)
	}
	path := c.ReqPath
	res.RedirectTo(path, r)
}
