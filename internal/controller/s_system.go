package controller

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
)

type (
	home          struct{}
	cSys          struct{}
	gen           struct{}
	api           struct{ *config.Search }
	role          struct{ *config.Search }
	cRoleApi      struct{ *config.Search }
	cRoleMenu     struct{ *config.Search }
	cDict         struct{ *config.Search }
	cFile         struct{ *config.Search }
	cAdmin        struct{ *config.Search }
	cMenu         struct{ *config.Search }
	cOperationLog struct{ *config.Search }
	ws            struct{}
)

var (
	Home = &home{}
	Sys  = &cSys{}
	Gen  = &gen{}
	Ws   = &ws{}
)

// ---home-------------------------------------------------------------------

func (c *home) IndexPage(r *ghttp.Request) {
	res.Page(r, "/index.html", g.Map{"icon": "/resource/image/v2ex.png"})
}

// ---system-----------------------------------------------------------------

func (s cSys) Path(r *ghttp.Request) {
	path := r.GetQuery("path")
	res.Page(r, path.String())
}

// ---Menu-----------------------------------------------------------------

var Menu = &cMenu{Search: &config.Search{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []*config.Field{
		{Name: "pid", SearchType: 1},
		{Name: "name", SearchType: 2},
		{Name: "path", SearchType: 2},
	},
}}

func (c *cMenu) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_menu.html", g.Map{"icon": icon})
}
func (c *cMenu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cMenu) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cMenu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cMenu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cMenu) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

func (c *cMenu) ListLevel1(r *ghttp.Request) {
	level1, err := sys.MenusLevel1(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(level1, r)
}

// ---api-------------------------------------------------------------------

var Api = &api{Search: &config.Search{
	T1: "s_api", Fields: []*config.Field{
		{Name: "method", SearchType: 1},
		{Name: "group", SearchType: 1},
		{Name: "desc", SearchType: 2},
		{Name: "status", SearchType: 1},
	},
}}

func (c *api) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *api) Post(r *ghttp.Request) {
	d := entity.Api{}
	if d.Status == 0 {
		res.Err(errors.New("状态不能为空"), r)
	}
	_ = r.Parse(&d)
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) Put(r *ghttp.Request) {
	d := entity.Api{}
	_ = r.Parse(&d)
	if d.Status == 0 {
		res.Err(errors.New("状态不能为空"), r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *api) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_api.html", g.Map{"icon": icon})
}

// ---role-------------------------------------------------------------------

var Role = &role{Search: &config.Search{
	T1: "s_role", Fields: []*config.Field{
		{Name: "id"},
		{Name: "name"},
		{Name: "created_at"},
		{Name: "updated_at"},
	},
}}

func (c *role) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *role) Post(r *ghttp.Request) {
	d := entity.Role{}
	_ = r.Parse(&d)
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) Put(r *ghttp.Request) {
	d := entity.Role{}
	_ = r.Parse(&d)
	if err := sys.Update(r.Context(), c.T1, d.Id, g.Map{"name": d.Name}); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *role) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

func (c *role) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_role.html", g.Map{"icon": icon})
}

func (c *role) Roles(r *ghttp.Request) {
	data, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

// ---roleApi-------------------------------------------------------------------

var RoleApi = &cRoleApi{Search: &config.Search{
	T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
	SearchFields: "t1.*,t2.name r_name,t3.url url ,t3.group,t3.method,t3.desc ", Fields: []*config.Field{
		{Name: "id"},
		{Name: "rid"},
		{Name: "aid"},
		{Name: "t2.name", QueryName: "r_name"},
		{Name: "t3.url"},
	},
}}

func (c *cRoleApi) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cRoleApi) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Aid []int
	}
	_ = r.Parse(&d)
	if err := sys.AddRoleApi(r.Context(), d.Rid, d.Aid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cRoleApi) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_role_api.html", g.Map{"icon": icon})
}
func (c *cRoleApi) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---roleMenu-------------------------------------------------------------------

var RoleMenu = &cRoleMenu{Search: &config.Search{
	T1:           "s_role_menu",
	T2:           "s_role  t2 on t1.rid = t2.id",
	T3:           "s_menu t3 on t1.mid = t3.id",
	OrderBy:      "t1.id desc",
	SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
	Fields: []*config.Field{
		{Name: "rid", SearchType: 1},
	},
}}

func (c *cRoleMenu) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_role_menu.html", g.Map{"icon": icon})
}
func (c *cRoleMenu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cRoleMenu) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRoleMenu) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Mid []int
	}
	_ = r.ParseForm(&d)
	if err := sys.AddRoleMenu(r.Context(), d.Rid, d.Mid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cRoleMenu) Put(r *ghttp.Request) {
	d := entity.RoleMenu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cRoleMenu) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cRoleMenu) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRoleMenu) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRoleMenu) CurrentMenus(r *ghttp.Request) {
	getAdmin, err := sys.GetAdmin(r)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(getAdmin.Menus, r)
}

//  ---admin-------------------------------------------------------------------

var Admin = &cAdmin{Search: &config.Search{
	T1: "s_admin", T2: "s_role t2 on t1.rid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.id,t1.uname,t1.rid,t1.status,t1.created_at,t1.updated_at,t2.name role_name",
	Fields: []*config.Field{
		{Name: "uname", SearchType: 2, QueryName: "uname"}, {Name: "t2.id", SearchType: 1, QueryName: "rid"}, {Name: "status", SearchType: 1, QueryName: "status"},
	},
}}

func (c *cAdmin) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_admin.html", g.Map{"icon": icon})
}
func (c *cAdmin) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cAdmin) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cAdmin) Post(r *ghttp.Request) {
	d := entity.Admin{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) Put(r *ghttp.Request) {
	d := entity.Admin{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c *cAdmin) Login(r *ghttp.Request) {
	var d struct {
		Uname string `form:"uname"`
		Pwd   string `form:"pwd"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Login(r.Context(), d.Uname, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) Logout(r *ghttp.Request) {
	err := sys.Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		OldPwd string `v:"required"`
		NewPwd string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwd(r.Context(), d.OldPwd, d.NewPwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdateUname(r *ghttp.Request) {
	var d struct {
		Uname string `v:"required"`
		Id    int64  `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminUname(r.Context(), d.Id, d.Uname); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cAdmin) UpdatePwdWithoutOldPwd(r *ghttp.Request) {
	var d struct {
		Pwd string `v:"required"`
		Id  string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.UpdateAdminPwdWithoutOldPwd(r.Context(), d.Id, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- Dict ------------------------------------------------------------------

var Dict = &cDict{Search: &config.Search{
	T1: "s_dict", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*config.Field{
		{Name: "k", SearchType: 2, QueryName: "k"}, {Name: "group", SearchType: 1, QueryName: "group"}, {Name: "type", SearchType: 1, QueryName: "type"}, {Name: "status", SearchType: 1, QueryName: "status"},
	},
}}

func (c *cDict) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_dict.html", g.Map{"icon": icon})
}
func (c *cDict) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cDict) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cDict) Post(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cDict) Put(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cDict) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

func (c *cDict) GetByKey(r *ghttp.Request) {
	s := r.Get("key").String()
	if s == "" {
		res.Err(consts.ErrParamEmpty, r)
	}
	v, err := sys.DictGetByKey(r.Context(), s)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(v, r)
}

// --- File -------------------------------------------------------------------

var File = &cFile{Search: &config.Search{
	T1: "s_file", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*config.Field{
		{Name: "url", SearchType: 2, QueryName: "url"}, {Name: "group", SearchType: 1, QueryName: "group"}, {Name: "status", SearchType: 1, QueryName: "status"},
	},
}}

func (c *cFile) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_file.html", g.Map{"icon": icon})
}
func (c *cFile) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cFile) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cFile) Post(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cFile) Put(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cFile) Del(r *ghttp.Request) {
	f, err := sys.GetFileById(r.Context(), xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	path, err := g.Cfg().Get(r.Context(), "server.rootFilePath")
	if err != nil {
		res.Err(err, r)
	}
	p := gfile.Pwd() + path.String() + "/" + f.Url
	if gfile.Exists(p) && gfile.IsFile(p) {
		_ = gfile.Remove(p)
	}
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cFile) Upload(r *ghttp.Request) {
	if err := sys.UploadFile(r.Context(), r); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- OperationLog ------------------------------------------------------------------------

var OperationLog = &cOperationLog{Search: &config.Search{
	T1: "s_operation_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
	Fields: []*config.Field{
		{Name: "t2.uname", SearchType: 2, QueryName: "uname"}, {Name: "content", SearchType: 2, QueryName: "content"}, {Name: "method", SearchType: 1, QueryName: "method"}, {Name: "uri", SearchType: 2, QueryName: "uri"}, {Name: "ip", SearchType: 2, QueryName: "ip"},
	},
}}

func (c *cOperationLog) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/s_operation_log.html", g.Map{"icon": icon})
}
func (c *cOperationLog) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *cOperationLog) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cOperationLog) Post(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cOperationLog) Put(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *cOperationLog) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---Gen Code-------------------------------------------------------------------

func (c gen) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/gen.html", g.Map{"icon": icon})
}
func (c gen) Tables(r *ghttp.Request) {
	data, err := sys.Tables(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c gen) Fields(r *ghttp.Request) {
	var d struct {
		Table string `v:"required#名表不能为空"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	data, err := sys.Fields(r.Context(), d.Table)
	res.OkData(data, r)
}

func (c gen) GenFile(r *ghttp.Request) {
	var d bo.GenConf
	// set genConf
	genConf := r.Get("genConf")
	if err := genConf.Struct(&d); err != nil {
		res.Err(err, r)
	}
	if err := d.SetUrlPrefix(); err != nil {
		res.Err(err, r)
	}
	// set fields
	d.Fields = make([]*bo.GenFiled, 0)
	for _, v := range r.Get("fields").MapStrVarDeep() {
		f := &bo.GenFiled{}
		if err := v.Struct(f); err != nil {
			res.Err(err, r)
		}
		if f.FieldType == "select" {
			f.Options = make([]*bo.FieldOption, 0)
			for _, v := range v.Map()["Options"].(map[string]interface{}) {
				f.Options = append(f.Options, &bo.FieldOption{
					Value: v.(map[string]interface{})["Value"].(string),
					Type:  v.(map[string]interface{})["Type"].(string),
					Label: v.(map[string]interface{})["Name"].(string),
				})
				if gstr.IsNumeric(fmt.Sprint(v.(map[string]interface{})["Value"])) {
					sort.Slice(f.Options, func(i, j int) bool { return gconv.Int(f.Options[i].Value) < gconv.Int(f.Options[j].Value) })
				}
			}
		}
		d.Fields = append(d.Fields, f)
	}
	sort.Slice(d.Fields, func(i, j int) bool {
		return d.Fields[i].Index < d.Fields[j].Index
	})
	if len(d.Fields) == 0 {
		res.Err(errors.New("字段不能为空"), r)
	}
	if err := sys.GenFile(r.Context(), &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- Ws ------------------------------------------------------------------------

func (w ws) GetUserWs(r *ghttp.Request) {
	sys.GetUserWs(r)
}
func (w ws) GetAdminWs(r *ghttp.Request) {
	sys.GetAdminWs(r)
}
func (w ws) NoticeUser(r *ghttp.Request) {
	var d struct {
		Uid     int `v:"required"`
		OrderId int `v:"required"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = sys.NoticeUser(gctx.New(), d.Uid, d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (w ws) NoticeAdmin(r *ghttp.Request) {
	var d struct {
		Msg string `v:"required" json:"msg"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = sys.NoticeAllAdmin(r.Context(), d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
