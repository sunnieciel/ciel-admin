package controller

import (
	"ciel-begin/internal/consts"
	"ciel-begin/internal/model/entity"
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"github.com/gogf/gf/v2/net/ghttp"
)

type menu struct {
	*config.SearchConf
}

func Menu() *menu {
	return &menu{SearchConf: &config.SearchConf{
		PageTitle: "Menu", UrlPrefix: "/menu", T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
		Fields: []*config.Field{
			{Field: "id", EditHidden: true},
			{Field: "pid", Type: "number", Search: true, Required: true},
			{Field: "name", Search: true, Required: true, Like: true},
			{Field: "path", Search: true, Like: true},
			{Field: "sort", Type: "number", Step: 0.01, Required: true},
			{Field: "type", Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "Normal"}, {Value: "2", Text: "Group"}}, Required: true},
			{Field: "status", Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "NO"}, {Value: "2", Text: "OFF"}}, Required: true},
			{Field: "created_at", EditHidden: true},
			{Field: "updated_at", EditHidden: true},
		},
	}}
}
func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, consts.DefaultPage, total, data, c)
}
func (c *menu) GetById(r *ghttp.Request) {
	data, err := service.System().GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *menu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
