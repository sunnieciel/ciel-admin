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

type cApi struct {
	Table   string
	FileDir string //主页文件位置
	ReqPath string
}

var Api = &cApi{
	Table:   "s_api",
	FileDir: "/sys/api",
	ReqPath: "/admin/api",
}

func (c cApi) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.Request.URL.Path
		s       = bo.Search{
			T1: c.Table, OrderBy: "t1.group,t1.id desc",
			Fields: []bo.Field{
				{Name: "method", Type: 1},
				{Name: "group", Type: 2},
				{Name: "status", Type: 1},
				{Name: "desc", Type: 2},
			},
		}
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
	apiGroup, err := sys.DictApiGroup(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/index.html"), g.Map{
		"list":      data,
		"page":      r.GetPage(total, s.Size).GetContent(3),
		"node":      node,
		"msg":       sys.MsgFromSession(r),
		"path":      reqPath,
		"api_group": apiGroup,
	}, r)
}
func (c cApi) AddIndex(r *ghttp.Request) {
	apiGroup, err := sys.DictApiGroup(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{
		"msg":       sys.MsgFromSession(r),
		"api_group": apiGroup,
	}, r)
}
func (c cApi) EditIndex(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Table, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	apiGroup, err := sys.DictApiGroup(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/edit.html"), g.Map{"msg": sys.MsgFromSession(r), "api_group": apiGroup}, r)
}
func (c cApi) Post(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := sys.Add(r.Context(), c.Table, &d); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap())), r)
}
func (c cApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	res.OkSession("删除成功", r)
	if err := sys.Del(r.Context(), c.Table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap())), r)
}
func (c cApi) Put(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := sys.Update(r.Context(), c.Table, d.Id, m); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())), r)
}
