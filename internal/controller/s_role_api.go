package controller

import (
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---roleApi-------------------------------------------------------------------
type roleApi struct {
	*config.SearchConf
}

func RoleApi() *roleApi {
	return &roleApi{SearchConf: &config.SearchConf{
		PageTitle: "RoleApi", PageUrl: "roleApi/list", UrlPrefix: "/roleApi",
		T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
		SearchFields: "t1.*,t2.name r_name,t3.url url", Fields: []*config.Field{
			{Field: "id", Type: "text", Hidden: true, EditHidden: true},
			{Field: "rid"},
			{Field: "aid"},
			{Field: "t2.name", QueryFiled: "r_name", Title: "Role", EditHidden: true},
			{Field: "t3.url", Title: "url", QueryFiled: "url", EditHidden: true},
		},
	}}
}
func (c *roleApi) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, "/sys/roleApi.html", total, data, c)
}
func (c *roleApi) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Aid []int
	}
	_ = r.Parse(&d)
	if err := service.Role().AddRoleApi(r.Context(), d.Rid, d.Aid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *roleApi) Del(r *ghttp.Request) {
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
