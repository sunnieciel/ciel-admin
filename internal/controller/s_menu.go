package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/admin"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type cMenu struct{ cBase }

var Menu = cMenu{cBase{"s_menu", "/admin/menu", "/sys/menu"}}

func (c cMenu) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = sys.MsgFromSession(r)
		s       = bo.Search{T1: c.Table, OrderBy: "t1.sort desc,t1.id desc", Fields: []bo.Field{
			{Name: "pid", Type: 1},
			{Name: "name", Type: 2},
			{Name: "path", Type: 2},
		}}
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
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cMenu) AddIndex(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cMenu) EditIndex(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": sys.MsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := sys.GetById(r.Context(), table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cMenu) Post(r *ghttp.Request) {
	var (
		d = entity.Menu{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := sys.Add(r.Context(), c.Table, &d); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(r.Context(), table, id); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) Put(r *ghttp.Request) {
	var (
		d     = entity.Menu{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := sys.Update(r.Context(), table, d.Id, m); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/menu", func(g *ghttp.RouterGroup) {
		g.Middleware(admin.AuthMiddleware)
		g.GET("/", c.Index)             // 主页面
		g.GET("/add", c.AddIndex)       // 添加页面
		g.GET("/edit/:id", c.EditIndex) // 修改页面
		g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
		g.GET("/del/:id", c.Del) // 删除请求
		g.POST("/post", c.Post)  // 添加请求
		g.POST("/put", c.Put)    // 修改请求
	})
}
