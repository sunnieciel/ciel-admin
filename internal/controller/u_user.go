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

type user struct {
	*config.SearchConf
}

var User = &user{SearchConf: &config.SearchConf{
	 T1:"u_user", T2:"u_login_log t2 on t1.id = t2.uid",OrderBy: "t1.id desc",SearchFields: "t1.*,t2.ip login_ip",
	Fields: []*config.Field{
		{Field: "id",QueryField: ""},{Field: "t2.ip", Like: true,QueryField: "login_ip"},{Field: "uname", Like: true,QueryField: ""},
	},
}}

func (c *user) Path(r *ghttp.Request) {
	icon, err := service.System().Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/u/u_user.html", g.Map{"icon": icon})
}
func (c *user) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *user) GetById(r *ghttp.Request) {
	data, err := service.System().GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *user) Post(r *ghttp.Request) {
	d := entity.User{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *user) Put(r *ghttp.Request) {
	d := entity.User{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *user) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
