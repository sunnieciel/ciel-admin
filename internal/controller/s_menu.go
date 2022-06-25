package controller

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---Menu-----------------------------------------------------------------

type cMenu struct{ *config.Search }

var Menu = &cMenu{Search: &config.Search{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []*config.Field{
		{Name: "pid", QueryName: "menu_pid", SearchType: 1},
		{Name: "name", QueryName: "menu_name", SearchType: 2},
		{Name: "path", QueryName: "menu_path", SearchType: 2},
	},
}}

func (c *cMenu) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/menu/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cMenu) PathAdd(r *ghttp.Request) {
	r.Response.WriteTpl("/sys/menu/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cMenu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cMenu) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cMenu) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	r.Session.Set("menu_edit", data.Map())
	r.Response.WriteTpl("/sys/menu/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cMenu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
