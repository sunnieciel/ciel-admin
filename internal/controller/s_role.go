package controller

import (
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---role-------------------------------------------------------------------
type role struct {
	*config.SearchConf
}

var Role = &role{SearchConf: &config.SearchConf{
	PageUrl: "/role/list",
	T1:      "s_role", Fields: []*config.Field{
		{Field: "id"},
		{Field: "name"},
		{Field: "created_at"},
		{Field: "updated_at"},
	},
}}

func (c *role) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *role) Post(r *ghttp.Request) {
	d := entity.Role{}
	_ = r.Parse(&d)
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) Put(r *ghttp.Request) {
	d := entity.Role{}
	_ = r.Parse(&d)
	if err := service.System().Update(r.Context(), c.T1, d.Id, g.Map{"name": d.Name}); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) GetById(r *ghttp.Request) {
	id := r.GetQuery("id")
	data, err := service.System().GetById(r.Context(), c.T1, id)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

func (c *role) Path(r *ghttp.Request) {
	icon, err := service.System().Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/role.html", g.Map{"icon": icon})
}
