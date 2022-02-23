package controller

import (
	"ciel-begin/internal/model/entity"
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---role-------------------------------------------------------------------
type role struct {
	*config.SearchConf
}

func Role() *role {
	return &role{SearchConf: &config.SearchConf{
		PageTitle: "Role", PageUrl: "/role/list", UrlPrefix: "/role",
		T1: "s_role", Fields: []*config.Field{
			{Field: "id", Type: "text", EditHidden: true},
			{Field: "name", Type: "text", Search: true},
			{Field: "created_at", Type: "text", EditHidden: true},
			{Field: "updated_at", Type: "text", EditHidden: true},
		},
	}}
}
func (c *role) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, "/sys/role.html", total, data, c)
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
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
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
