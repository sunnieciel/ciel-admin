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

type dict struct {
	*config.SearchConf
}

func Dict() *dict {
	return &dict{SearchConf: &config.SearchConf{
		PageTitle: "Dict", UrlPrefix: "/dict", T1: "s_dict",
		Fields: []*config.Field{
			{Field: "id", EditHidden: true},
			{Field: "k", Search: true, Like: true},
			{Field: "v", Search: true, Like: true},
			{Field: "desc", Search: true, Like: true},
			{Field: "group", Search: true, Like: true},
			{Field: "type", Search: true, Type: "select", Items: []*config.Item{
				{Value: 1, Text: "TEXT"},
				{Value: 2, Text: "IMG"},
				{Value: 3, Text: "HTML"},
			}},
			{Field: "status", Type: "select", Items: []*config.Item{{Value: "1", Text: "NO"}, {Value: "2", Text: "OFF"}}, Required: true},
			{Field: "created_at", EditHidden: true},
			{Field: "updated_at", EditHidden: true},
		},
	}}
}
func (c *dict) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, consts.DefaultPage, total, data, c)
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
