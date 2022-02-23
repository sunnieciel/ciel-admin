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

// ---api-------------------------------------------------------------------
type api struct {
	*config.SearchConf
}

func Api() *api {
	return &api{SearchConf: &config.SearchConf{
		PageTitle: "Api", PageUrl: "/api/list",
		T1: "s_api", UrlPrefix: "/api", Fields: []*config.Field{
			{Field: "id", EditHidden: true},
			{Field: "url", Required: true},
			{Field: "method", Type: "select", Required: true, Items: []*config.Item{
				{Text: "GET", Value: "GET"},
				{Text: "POST", Value: "POST"},
				{Text: "PUT", Value: "PUT"},
				{Text: "DELETE", Value: "DELETE"}},
			},
			{Field: "group", Required: true},
			{Field: "desc"},
			{Field: "status", Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "ON"}, {Value: "2", Text: "OFF"}}, Required: true},
		},
	}}
}
func (c *api) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, consts.DefaultPage, total, data, c)
}
func (c *api) Post(r *ghttp.Request) {
	d := entity.Api{}
	_ = r.Parse(&d)
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) Put(r *ghttp.Request) {
	d := entity.Api{}
	_ = r.Parse(&d)
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) GetById(r *ghttp.Request) {
	id := r.GetQuery("id")
	data, err := service.System().GetById(r.Context(), c.T1, id)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
