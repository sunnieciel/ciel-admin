package controller

import (
	"ciel-admin/internal/service"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---roleApi-------------------------------------------------------------------
type roleApi struct {
	*config.SearchConf
}

var RoleApi = &roleApi{SearchConf: &config.SearchConf{
	PageUrl: "roleApi/list",
	T1:      "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
	SearchFields: "t1.*,t2.name r_name,t3.url url ,t3.group,t3.method,t3.desc ", Fields: []*config.Field{
		{Field: "id"},
		{Field: "rid"},
		{Field: "aid"},
		{Field: "t2.name", QueryField: "r_name"},
		{Field: "t3.url"},
	},
}}

func (c *roleApi) Path(r *ghttp.Request) {
	res.Page(r, "/sys/roleApi.html")
}

func (c *roleApi) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
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
