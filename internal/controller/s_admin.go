package controller

import (
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xpwd"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

//  ---admin-------------------------------------------------------------------
type admin struct{ *config.SearchConf }

var Admin = &admin{
	SearchConf: &config.SearchConf{
		PageUrl:      "/admin/list",
		T1:           "s_admin",
		SearchFields: "id,rid,uname,status,created_at,updated_at",
		Fields: []*config.Field{
			{Field: "id"},
			{Field: "uname", Like: true},
			{Field: "pwd"},
			{Field: "rid"},
			{Field: "status"},
			{Field: "created_at"},
			{Field: "updated_at"},
		},
	}}

func (c *admin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}

func (c *admin) Path(r *ghttp.Request) {
	res.Page(r, "/sys/admin.html")
}
func (c *admin) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *admin) GetById(r *ghttp.Request) {
	data, err := service.System().GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	gMap := data.GMap()
	gMap.Remove("pwd")
	res.OkData(gMap.Map(), r)
}
func (c *admin) Post(r *ghttp.Request) {
	d := entity.Admin{}
	_ = r.Parse(&d)
	m := gconv.Map(d)
	if d.Pwd != "" {
		m["pwd"] = xpwd.GenPwd(d.Pwd)
	} else {
		delete(m, "pwd")
	}
	if err := service.System().Add(r.Context(), c.T1, m); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Put(r *ghttp.Request) {
	d := entity.Admin{}
	_ = r.Parse(&d)
	m := gconv.Map(d)
	if d.Pwd != "" {
		m["pwd"] = xpwd.GenPwd(d.Pwd)
	} else {
		delete(m, "pwd")
	}
	if err := service.System().Update(r.Context(), c.T1, d.Id, m); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Login(r *ghttp.Request) {
	var d struct {
		Uname string `form:"uname"`
		Pwd   string `form:"pwd"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Admin().Login(r.Context(), d.Uname, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Logout(r *ghttp.Request) {
	err := service.Admin().Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		OldPwd string `v:"required"`
		NewPwd string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	err := service.Admin().UpdateAdminPwd(r.Context(), d.OldPwd, d.NewPwd)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
