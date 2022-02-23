package controller

import (
	"ciel-begin/internal/consts"
	"ciel-begin/internal/model/entity"
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"ciel-begin/utility/utils/xpwd"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

//  ---admin-------------------------------------------------------------------
type admin struct {
	*config.SearchConf
	loginPage string
}

func Admin() *admin {
	c := admin{
		loginPage: "login.html",
		SearchConf: &config.SearchConf{
			PageTitle: "Admin", PageUrl: "/admin/list", UrlPrefix: "/admin",
			T1:           "s_admin",
			SearchFields: "id,rid,uname,status,created_at,updated_at",
			Fields: []*config.Field{
				{Field: "id", EditHidden: true},
				{Field: "uname", Required: true},
				{Field: "pwd", Hidden: true},
				{Field: "rid", Title: "Role Id", Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "Super Admin"}, {Value: "2", Text: "Admin"}}, Required: true},
				{Field: "status", Title: "status", Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "ON"}, {Value: "2", Text: "OFF"}}, Required: true},
				{Field: "created_at", EditHidden: true},
				{Field: "updated_at", EditHidden: true},
			},
		}}
	return &c
}
func (c *admin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}

func (c *admin) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, consts.DefaultPage, total, data, c)
}
func (c *admin) GetById(r *ghttp.Request) {
	id := r.GetQuery("id")
	data, err := service.System().GetById(r.Context(), c.T1, id)
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
	r.Response.RedirectTo("/login")
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
