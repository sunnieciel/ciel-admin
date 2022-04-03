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

type dict struct {
	*config.SearchConf
}

var Dict = &dict{SearchConf: &config.SearchConf{
	PageUrl: "/dict/list", T1: "s_dict",
	Fields: []*config.Field{
		{Field: "id"},
		{Field: "k", Like: true},
		{Field: "v", Like: true},
		{Field: "desc", Like: true},
		{Field: "group"},
		{Field: "type"},
		{Field: "status"},
	},
}}

func (c *dict) Path(r *ghttp.Request) {
	icon, err := service.System().Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/dict.html", g.Map{"icon": icon})
}
func (c *dict) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *dict) GetById(r *ghttp.Request) {
	data, err := service.System().GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *dict) Post(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *dict) Put(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *dict) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
