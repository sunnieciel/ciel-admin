package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type cRoleMenu struct{ cBase }

var RoleMenu = &cRoleMenu{
	cBase{Table: "s_role_menu", ReqPath: "/admin/roleMenu", FileDir: "/sys/roleMenu"}}

func (c cRoleMenu) Path(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = bo.Search{
			T1:           "s_role_menu",
			T2:           "s_role  t2 on t1.rid = t2.id",
			T3:           "s_menu t3 on t1.mid = t3.id",
			OrderBy:      "t1.id desc",
			SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
			Fields: []bo.Field{
				{Name: "rid", Type: 1},
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

	node.Name = "角色菜单"
	node.Path = c.ReqPath
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(ctx, s)
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
func (c cRoleMenu) PathAdd(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		rid  = r.Get("rid")
		file = fmt.Sprintf("%s/add.html", c.FileDir)
		msg  = sys.MsgFromSession(r)
	)
	menus, err := sys.RoleNoMenu(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{"msg": msg, "menus": menus}, r)
}
func (c cRoleMenu) Post(r *ghttp.Request) {
	var (
		d struct {
			Rid int
			Mid []int
		}
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	_ = r.ParseForm(&d)
	res.OkSession("添加成功", r)
	if err := sys.AddRoleMenu(ctx, d.Rid, d.Mid); err != nil {
		res.ErrSession(err, r)
	}
	r.Response.RedirectTo(path)
}
func (c cRoleMenu) Del(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		id    = r.Get("id")
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
