package controller

import (
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---roleMenu-------------------------------------------------------------------
type roleMenu struct {
	*config.SearchConf
}

func RoleMenu() *roleMenu {
	return &roleMenu{SearchConf: &config.SearchConf{
		PageTitle: "RoleMenu", UrlPrefix: "/roleMenu", NoEdit: true,
		T1:           "s_role_menu",
		T2:           "s_role  t2 on t1.rid = t2.id",
		T3:           "s_menu t3 on t1.mid = t3.id",
		SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
		Fields: []*config.Field{
			{Field: "id", EditHidden: true, Hidden: true},
			{Field: "rid", Title: "Role ID", Search: true},
			{Field: "t2.name", QueryFiled: "role_name", Title: "Role Name", EditHidden: true},
			{Field: "mid", Title: "Menu ID", Type: "text"},
			{Field: "t3.name", QueryFiled: "menu_name", Title: "Menu Name", EditHidden: true},
		},
	}}
}
func (c *roleMenu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, "sys/roleMenu.html", total, data, c)
}
func (c *roleMenu) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Mid []int
	}
	_ = r.ParseForm(&d)
	if err := service.Role().AddRoleMenu(r.Context(), d.Rid, d.Mid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *roleMenu) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *roleMenu) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := service.Role().RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *roleMenu) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := service.Role().RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
