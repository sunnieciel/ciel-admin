package controller

import (
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xpwd"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	home     struct{}
	cSys     struct{}
	rss      struct{}
	gen      struct{}
	api      struct{ *config.SearchConf }
	role     struct{ *config.SearchConf }
	roleApi  struct{ *config.SearchConf }
	admin    struct{ *config.SearchConf }
	roleMenu struct{ *config.SearchConf }
	dict     struct{ *config.SearchConf }
	file     struct{ *config.SearchConf }
	menu     struct{ *config.SearchConf }
	ws       struct{}
)

var (
	Home = &home{}
	Sys  = &cSys{}
	Rss  = &rss{}
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
func (s cSys) PathGithub(r *ghttp.Request) {
	res.Page(r, "/sys/rss/github.html", g.Map{"icon": "/resource/image/github.png"})
}
func (s cSys) OsChina(r *ghttp.Request) {
	res.Page(r, "/sys/rss/oschina.html", g.Map{"icon": "/resource/image/github.png"})
}
func (s cSys) Douban(r *ghttp.Request) {
	res.Page(r, "/sys/rss/douban.html", g.Map{"icon": "/resource/image/github.png"})
}

// ---Menu-------------------------------------------------------------------

var Menu = &menu{SearchConf: &config.SearchConf{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []*config.Field{
		{Field: "pid"},
		{Field: "status"},
		{Field: "name", Like: true},
		{Field: "path", Like: true},
	},
}}

func (c *menu) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/menu2.html", g.Map{"icon": icon})
}
func (c *menu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *menu) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
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
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *menu) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---api-------------------------------------------------------------------

var Api = &api{SearchConf: &config.SearchConf{
	PageUrl: "/api/list",
	T1:      "s_api", Fields: []*config.Field{
		{Field: "id"},
		{Field: "url"},
		{Field: "method"},
		{Field: "group"},
		{Field: "desc"},
		{Field: "status"},
	},
}}

func (c *api) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *api) Post(r *ghttp.Request) {
	d := entity.Api{}
	_ = r.Parse(&d)
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *api) Put(r *ghttp.Request) {
	d := entity.Api{}
	_ = r.Parse(&d)
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
	id := r.GetQuery("id")
	data, err := sys.GetById(r.Context(), c.T1, id)
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
	res.Page(r, "/sys/api.html", g.Map{"icon": icon})
}

// ---role-------------------------------------------------------------------

var Role = &role{SearchConf: &config.SearchConf{
	PageUrl: "/role/list",
	T1:      "s_role", Fields: []*config.Field{
		{Field: "id"},
		{Field: "name"},
		{Field: "created_at"},
		{Field: "updated_at"},
	},
}}

func (c *role) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
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
	id := r.GetQuery("id")
	data, err := sys.GetById(r.Context(), c.T1, id)
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
	res.Page(r, "/sys/role.html", g.Map{"icon": icon})
}

func (c *role) Roles(r *ghttp.Request) {
	data, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

// ---roleApi-------------------------------------------------------------------

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
	total, data, err := sys.List(r.Context(), c.SearchConf)
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
	if err := sys.AddRoleApi(r.Context(), d.Rid, d.Aid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *roleApi) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---roleMenu-------------------------------------------------------------------

var RoleMenu = &roleMenu{SearchConf: &config.SearchConf{
	PageUrl:      "/roleMenu/list",
	T1:           "s_role_menu",
	T2:           "s_role  t2 on t1.rid = t2.id",
	T3:           "s_menu t3 on t1.mid = t3.id",
	SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
	Fields: []*config.Field{
		{Field: "id"},
		{Field: "rid"},
		{Field: "t2.name", QueryField: "role_name", Like: true},
		{Field: "mid"},
		{Field: "t3.name", QueryField: "menu_name"},
	},
}}

func (c *roleMenu) Path(r *ghttp.Request) {
	res.Page(r, "sys/roleMenu.html")
}
func (c *roleMenu) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *roleMenu) Post(r *ghttp.Request) {
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
func (c *roleMenu) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *roleMenu) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *roleMenu) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *roleMenu) CurrentMenus(r *ghttp.Request) {
	getAdmin, err := sys.GetAdmin(r)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(getAdmin.Menus, r)
}

//  ---admin-------------------------------------------------------------------

var Admin = &admin{
	SearchConf: &config.SearchConf{
		PageUrl:      "/admin/list",
		T1:           "s_admin",
		T2:           "s_role t2 on t1.rid = t2.id",
		SearchFields: "t1.id,t1.rid,t1.uname,t1.status,t1.created_at,t1.updated_at,t2.name role_name",
		Fields: []*config.Field{
			{Field: "id"},
			{Field: "uname", Like: true},
			{Field: "rid"},
			{Field: "status"},
		},
	}}

func (c *admin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c *admin) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/admin.html", g.Map{"icon": icon})
}
func (c *admin) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *admin) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
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
	if err := sys.Add(r.Context(), c.T1, m); err != nil {
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
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
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
	if err := sys.Login(r.Context(), d.Uname, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *admin) Logout(r *ghttp.Request) {
	err := sys.Logout(r.Context())
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
	if err := sys.UpdateAdminPwd(r.Context(), d.OldPwd, d.NewPwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- Dict ------------------------------------------------------------------

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
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/dict.html", g.Map{"icon": icon})
}
func (c *dict) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *dict) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
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
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *dict) Put(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *dict) Del(r *ghttp.Request) {
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --- File -------------------------------------------------------------------

var File = &file{SearchConf: &config.SearchConf{
	PageUrl: "/file/list", T1: "s_file",
	Fields: []*config.Field{
		{Field: "id"},
		{Field: "img"},
		{Field: "group", Like: true},
		{Field: "status"},
		{Field: "url"},
		{Field: "created_at"},
		{Field: "updated_at"},
	},
}}

func (c *file) Path(r *ghttp.Request) {
	icon, err := sys.Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/file.html", g.Map{"icon": icon})
}
func (c *file) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := sys.List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.OkPage(page, size, total, data, r)
}
func (c *file) GetById(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *file) Post(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *file) Put(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *file) Del(r *ghttp.Request) {
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
func (c *file) Upload(r *ghttp.Request) {
	if err := sys.UploadFile(r.Context(), r); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---Rss-------------------------------------------------------------------

func (c rss) Fetch(r *ghttp.Request) {
	data, err := sys.FetchRss(r.Context(), r.GetQuery("url").String())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

func (c rss) V2ex(r *ghttp.Request) {
	data, err := sys.FetchRss(r.Context(), "https://www.v2ex.com/index.xml")
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
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
