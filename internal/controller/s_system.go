// =================================================================================
// This is auto-generated by FreeKey Admin at 2022-07-22 20:14:53. For more information see https://github.com/1211ciel/ciel-admin
// =================================================================================

package controller

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xpwd"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"

	captcha "github.com/mojocn/base64Captcha"
)

// ---Home----------------------------------------------------------------------
type home struct{}

var Home = &home{}

func (c home) IndexPage(r *ghttp.Request) {
	r.Response.RedirectTo(g.Config().MustGet(r.Context(), "home").String())
}

// ---Menu-----------------------------------------------------------------

type cMenu struct{ bo.Search }

var Menu = cMenu{Search: bo.Search{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []bo.Field{
		{Name: "pid", QueryName: "menu_pid", SearchType: 1},
		{Name: "name", QueryName: "menu_name", SearchType: 2},
		{Name: "path", QueryName: "menu_path", SearchType: 2},
	},
}}

func (c cMenu) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/menu/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cMenu) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/menu/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cMenu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/menu/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cMenu) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/menu/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cMenu) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("menu_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/menu/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cMenu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	g.Log().Notice(nil, m)
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/menu/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

// ---Api-----------------------------------------------------------------

type cApi struct{ bo.Search }

var Api = &cApi{Search: bo.Search{
	T1: "s_api", OrderBy: "t1.group,t1.id desc", SearchFields: "t1.*",
	Fields: []bo.Field{
		{Name: "method", SearchType: 1, QueryName: "api_method"}, {Name: "group", SearchType: 2, QueryName: "api_group"}, {Name: "status", SearchType: 1, QueryName: "api_status"},
		{Name: "desc", SearchType: 2, QueryName: "api_desc"},
	},
}}

func (c cApi) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	apiGroup, err := sys.DictApiGroup(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/api/index.html", g.Map{
		"list":      data,
		"page":      r.GetPage(total, c.Size).GetContent(3),
		"node":      node,
		"msg":       sys.MsgFromSession(r),
		"path":      r.URL.Path,
		"api_group": apiGroup,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cApi) PathAdd(r *ghttp.Request) {
	apiGroup, err := sys.DictApiGroup(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/api/add.html", g.Map{
		"msg":       sys.MsgFromSession(r),
		"api_group": apiGroup,
	})
}
func (c cApi) Post(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/api/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/api/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cApi) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	apiGroup, err := sys.DictApiGroup(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("api_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/api/edit.html", g.Map{"msg": sys.MsgFromSession(r), "api_group": apiGroup})
}
func (c cApi) Put(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/api/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

// ---Role-----------------------------------------------------------------

type (
	cRole     struct{ bo.Search }
	cRoleApi  struct{ bo.Search }
	cRoleMenu struct{ bo.Search }
)

var (
	Role = &cRole{Search: bo.Search{
		T1: "s_role", OrderBy: "t1.id desc", SearchFields: "t1.*",
		Fields: []bo.Field{},
	}}
	RoleMenu = &cRoleMenu{Search: bo.Search{
		T1:           "s_role_menu",
		T2:           "s_role  t2 on t1.rid = t2.id",
		T3:           "s_menu t3 on t1.mid = t3.id",
		OrderBy:      "t1.id desc",
		SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
		Fields: []bo.Field{
			{Name: "rid", SearchType: 1},
		},
	}}
	RoleApi = &cRoleApi{Search: bo.Search{
		T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
		OrderBy:      "t3.group",
		SearchFields: "t1.*,t2.name r_name,t3.url url,t3.group,t3.method,t3.desc ", Fields: []bo.Field{
			{Name: "id"},
			{Name: "rid", SearchType: 1},
			{Name: "aid"},
			{Name: "t2.name", QueryName: "r_name", SearchType: 2},
			{Name: "t3.url"},
		},
	}}
)

func (c cRole) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/role/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cRole) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/role/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cRole) Post(r *ghttp.Request) {
	d := entity.Role{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/role/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRole) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/role/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRole) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("role_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/role/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cRole) Put(r *ghttp.Request) {
	d := entity.Role{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/role/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRole) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c cRole) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c cRoleMenu) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	node.Name = "角色菜单"
	node.Path = "/admin/roleMenu/path"
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/roleMenu/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cRoleMenu) PathAdd(r *ghttp.Request) {
	menus, err := sys.RoleNoMenu(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/roleMenu/add.html", g.Map{"msg": sys.MsgFromSession(r), "menus": menus})
}
func (c cRoleMenu) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Mid []int
	}
	_ = r.ParseForm(&d)
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.AddRoleMenu(r.Context(), d.Rid, d.Mid); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/roleMenu/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRoleMenu) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/roleMenu/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRoleApi) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	node.Path = "/admin/roleApi/path"
	node.Name = "角色禁用API"
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/roleApi/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cRoleApi) PathAdd(r *ghttp.Request) {
	apis, err := sys.RoleNoApi(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/roleApi/add.html", g.Map{"msg": sys.MsgFromSession(r), "apis": apis})
}
func (c cRoleApi) Post(r *ghttp.Request) {
	var d struct {
		Rid int
		Aid []int
	}
	_ = r.Parse(&d)
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.AddRoleApi(r.Context(), d.Rid, d.Aid); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/roleApi/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRoleApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/roleApi/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cRoleApi) Clear(r *ghttp.Request) {
	err := sys.ClearRoleApi(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---Dict-----------------------------------------------------------------

type cDict struct{ bo.Search }

var Dict = &cDict{Search: bo.Search{
	T1: "s_dict", OrderBy: "t1.group,t1.id desc", SearchFields: "t1.*",
	Fields: []bo.Field{
		{Name: "k", SearchType: 2, QueryName: "dict_k"}, {Name: "v", SearchType: 2, QueryName: "dict_v"}, {Name: "desc", SearchType: 2, QueryName: "dict_desc"}, {Name: "group", SearchType: 1, QueryName: "dict_group"}, {Name: "status", SearchType: 1, QueryName: "dict_status"}, {Name: "type", SearchType: 1, QueryName: "dict_type"},
	},
}}

func (c cDict) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/dict/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cDict) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/dict/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cDict) Post(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/dict/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cDict) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/dict/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cDict) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("dict_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/dict/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cDict) Put(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	if d.K == "white_ips" {
		if err := sys.SetWhiteIps(r.Context(), d.V); err != nil {
			res.Err(err, r)
		}
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/dict/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

type cFile struct{ bo.Search }

var File = &cFile{Search: bo.Search{
	T1: "s_file", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []bo.Field{
		{Name: "url", SearchType: 2, QueryName: "file_url"}, {Name: "group", SearchType: 1, QueryName: "file_group"}, {Name: "status", SearchType: 1, QueryName: "file_status"},
	},
}}

func (c cFile) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/file/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cFile) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/file/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cFile) Post(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/file/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cFile) Del(r *ghttp.Request) {
	f, err := sys.GetFileById(r.Context(), xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	path, err := g.Cfg().Get(r.Context(), "server.rootFilePath")
	if err != nil {
		res.Err(err, r)
	}
	p := gfile.Pwd() + path.String() + "/" + f.Url
	if err = sys.RemoveFile(r.Context(), p); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err = sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/file/path?", xurl.ToUrlParams(r.GetQueryMap())))
}

func (c cFile) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("file_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/file/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cFile) Put(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/file/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cFile) Upload(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "上传成功")
	if err := sys.UploadFile(r.Context(), r); err != nil {
		msg = fmt.Sprintf(consts.MsgPrimary, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo("/admin/file/path/add?" + xurl.ToUrlParams(r.GetQueryMap()))
}

type cOperationLog struct{ bo.Search }

var OperationLog = &cOperationLog{Search: bo.Search{
	T1: "s_operation_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
	Fields: []bo.Field{
		{Name: "t2.uname", SearchType: 2, QueryName: "operationLog_uname"}, {Name: "content", SearchType: 2, QueryName: "operationLog_content"},
	},
}}

func (c cOperationLog) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/operationLog/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cOperationLog) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/operationLog/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cOperationLog) Post(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/operationLog/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cOperationLog) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/operationLog/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cOperationLog) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("operationLog_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/operationLog/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cOperationLog) Put(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/operationLog/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cOperationLog) Clear(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "操作成功")
	if err := sys.OperationLogClear(r.Context()); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo("/admin/operationLog/path")
}

// ---System-----------------------------------------------------------------

type cSys struct{}

var Sys = &cSys{}

func (s cSys) Level1(r *ghttp.Request) {
	level1, err := sys.MenusLevel1(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(level1, r)
}
func (s cSys) GetDictByKey(r *ghttp.Request) {
	data, err := sys.DictGetByKey(r.Context(), r.Get("key").String())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (s cSys) To(r *ghttp.Request) {
	name := r.Get("path")
	if name.IsEmpty() || name.String() == "null" {
		res.Err(fmt.Errorf("filename prefix cannot be empty"), r)
	}
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	if node.FilePath == "" {
		res.Err(fmt.Errorf("node file path is empty"), r)
	}
	_ = r.Response.WriteTpl(node.FilePath, g.Map{
		"node": node,
		"path": r.URL.Path,
	})
}
func (s cSys) GetCaptcha(r *ghttp.Request) {
	var driver = sys.NewDriver().ConvertFonts()
	c := captcha.NewCaptcha(driver, sys.Store)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	id := r.GetQuery("id").String()
	item, _ := c.Driver.DrawCaptcha(content)
	c.Store.Set(id, answer)
	res.OkData(item.EncodeB64string(), r)
}
func (s cSys) Quotations(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), "/to/quotations")
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/tool/quotations.html", g.Map{"node": node}); err != nil {
		res.Err(err, r)
	}
}

//  ---admin-------------------------------------------------------------------

type cAdmin struct{ bo.Search }

var Admin = &cAdmin{Search: bo.Search{
	T1: "s_admin", T2: "s_role t2 on t1.rid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.name role_name",
	Fields: []bo.Field{
		{Name: "rid", SearchType: 1, QueryName: "admin_rid"}, {Name: "status", SearchType: 1, QueryName: "admin_status"},
	},
}}

func (c cAdmin) Path(r *ghttp.Request) {
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/admin/index.html", g.Map{
		"list":  data,
		"page":  r.GetPage(total, c.Size).GetContent(3),
		"node":  node,
		"msg":   sys.MsgFromSession(r),
		"roles": roles,
		"path":  r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cAdmin) PathAdd(r *ghttp.Request) {
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/admin/add.html", g.Map{"msg": sys.MsgFromSession(r), "roles": roles})
}
func (c cAdmin) Post(r *ghttp.Request) {
	d := entity.Admin{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	d.Pwd = xpwd.GenPwd(d.Pwd)
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/admin/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdmin) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/admin/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdmin) PathEdit(r *ghttp.Request) {
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("admin_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/admin/edit.html", g.Map{"msg": sys.MsgFromSession(r), "roles": roles})
}
func (c cAdmin) Put(r *ghttp.Request) {
	d := entity.Admin{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "pwd")
	delete(m, "createdAt")
	g.Log().Notice(nil, m)
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/admin/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdmin) LoginPage(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c cAdmin) Login(r *ghttp.Request) {
	var d struct {
		Uname string `form:"uname"`
		Pwd   string `form:"pwd"`
		Id    string `form:"id"`   // 获取二维码时的id
		Code  string `from:"code"` // 二维码
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Login(r.Context(), d.Id, d.Code, d.Uname, d.Pwd, r.GetClientIp()); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) Logout(r *ghttp.Request) {
	err := sys.Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) UpdatePwd(r *ghttp.Request) {
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
func (c cAdmin) UpdateUname(r *ghttp.Request) {
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
func (c cAdmin) UpdatePwdWithoutOldPwd(r *ghttp.Request) {
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

var AdminLoginLog = &cAdminLoginLog{Search: bo.Search{
	T1: "s_admin_login_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
	Fields: []bo.Field{
		{Name: "uid", SearchType: 1, QueryName: "adminLoginLog_uid"}, {Name: "t2.uname", SearchType: 2, QueryName: "adminLoginLog_uname"}, {Name: "area", SearchType: 2, QueryName: "adminLoginLog_area"},
	},
}}

type cAdminLoginLog struct{ bo.Search }

func (c cAdminLoginLog) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/adminLoginLog/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cAdminLoginLog) PathAdd(r *ghttp.Request) {
	r.Response.WriteTpl("/sys/adminLoginLog/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cAdminLoginLog) Post(r *ghttp.Request) {
	d := entity.AdminLoginLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminLoginLog/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdminLoginLog) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminLoginLog/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdminLoginLog) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("adminLoginLog_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/adminLoginLog/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cAdminLoginLog) Put(r *ghttp.Request) {
	d := entity.AdminLoginLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminLoginLog/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

func (c cAdminLoginLog) Clear(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "操作成功")
	if err := sys.ClearAdminLog(r.Context()); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminLoginLog/path?", xurl.ToUrlParams(r.GetQueryMap())))
	res.Ok(r)
}

// --adminMessage ---------------------------------------------------------------------
type cAdminMessage struct{ bo.Search }

var AdminMessage = &cAdminMessage{Search: bo.Search{
	T1: "s_admin_message", T2: "s_admin t2 on t2.id=t1.from_uid", T3: "s_admin t3 on t3.id=t1.to_uid", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname from_uname,t3.uname to_uname",
	Fields: []bo.Field{
		{Name: "t2.from_uname", SearchType: 2, QueryName: "adminMessage_from_uname"}, {Name: "t3.uname", SearchType: 2, QueryName: "adminMessage_to_uname"}, {Name: "group", SearchType: 1, QueryName: "adminMessage_group"}, {Name: "type", SearchType: 1, QueryName: "adminMessage_type"}, {Name: "content", SearchType: 2, QueryName: "adminMessage_content"}, {Name: "link", SearchType: 2, QueryName: "adminMessage_link"},
	},
}}

func (c cAdminMessage) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	adminOptions, err := sys.GetAllAdminOptions(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/adminMessage/index.html", g.Map{
		"list":         data,
		"page":         r.GetPage(total, c.Size).GetContent(3),
		"node":         node,
		"msg":          sys.MsgFromSession(r),
		"adminOptions": adminOptions.String(),
		"path":         r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cAdminMessage) PathAdd(r *ghttp.Request) {
	r.Response.WriteTpl("/sys/adminMessage/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cAdminMessage) Post(r *ghttp.Request) {
	d := entity.AdminMessage{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	admin, err := sys.GetAdmin(r)
	if err != nil {
		res.Err(err, r)
	}
	d.FromUid = admin.Admin.Id
	d.Status = 1
	msg := fmt.Sprintf(consts.MsgPrimary, "发送成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	d.Content = ""
	_ = r.Session.Set("adminMessage_edit", d)
	if err = sys.NoticeAdmin(r.Context(), "{'msg':'hello'}", d.ToUid); err != nil {
		res.Err(err, r)
	}
	if err = sys.AddAdminUnReadMsg(r.Context(), d.ToUid); err != nil {
		res.Err(err, r)
	}
	r.Response.RedirectTo(fmt.Sprint("/admin/adminMessage/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdminMessage) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminMessage/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdminMessage) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("adminMessage_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/adminMessage/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c cAdminMessage) Put(r *ghttp.Request) {
	d := entity.AdminMessage{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/adminMessage/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cAdminMessage) UnreadMsgCount(r *ghttp.Request) {
	count, err := sys.GetAdminUnreadMsgCount(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(count, r)
}
func (c cAdminMessage) ClearUnreadMsg(r *ghttp.Request) {
	err := sys.ClearUnreadMsg(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdminMessage) Clear(r *ghttp.Request) {
	if err := sys.ClearAdminMessage(r.Context(), r.Get("group").String()); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// --Node ---------------------------------------------------------------------
type cNode struct{ bo.Search }

var Node = &cNode{Search: bo.Search{
	T1: "f_node", T2: "s_admin t2 on t1.uid= t2.id", OrderBy: "t1.year desc,t1.month desc,t1.day desc,t1.id desc", SearchFields: "t1.*,t2.uname",
	Fields: []bo.Field{
		{Name: "t1.year", SearchType: 1, QueryName: "node_year"},
		{Name: "t1.month", SearchType: 1, QueryName: "node_month"},
		{Name: "t1.day", SearchType: 1, QueryName: "node_day"},
		{Name: "t1.category", SearchType: 1, QueryName: "node_category"},
		{Name: "t1.summary", SearchType: 2, QueryName: "node_summary"},
		{Name: "level", SearchType: 1, QueryName: "node_level"}, {Name: "tag", SearchType: 2, QueryName: "node_tag"}, {Name: "main_things", SearchType: 2, QueryName: "node_main_things"},
	},
}}

func (c cNode) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/f/node/index.html", g.Map{
		"list":     data,
		"page":     r.GetPage(total, c.Size).GetContent(3),
		"node":     node,
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
		"path":     r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}

func (c cNode) PathAdd(r *ghttp.Request) {
	r.Response.WriteTpl("/f/node/add.html", g.Map{
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
	})
}
func (c cNode) Post(r *ghttp.Request) {
	data := entity.Node{}
	if err := r.Parse(&data); err != nil {
		res.Err(err, r)
	}
	y := fmt.Sprint(gtime.Now().Year())
	m := gtime.Now().Month()
	d := time.Now().Day()
	happenDate := r.Get("happen_date").String()
	if happenDate != "" {
		split := strings.Split(happenDate, ",")
		switch len(split) {
		case 1:
			y = gconv.String(split[0])
		case 2:
			y = gconv.String(split[0])
			m = gconv.Int(split[1])
		case 3:
			y = gconv.String(split[0])
			m = gconv.Int(split[1])
			d = gconv.Int(split[2])
		}
	}
	data.Year = y
	data.Month = m
	data.Day = d
	admin, _ := sys.GetAdmin(r)
	data.Uid = admin.Admin.Id
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	id, err := sys.AddGetID(r.Context(), c.T1, &data)
	if err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/node/path/edit/", id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cNode) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/node/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cNode) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("node_edit", data.Map())
	_ = r.Response.WriteTpl("/f/node/edit.html", g.Map{
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
	})
}
func (c cNode) Put(r *ghttp.Request) {
	d := entity.Node{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/node/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func getCategory(r *ghttp.Request) []g.Map {
	key, err := sys.DictGetByKey(r.Context(), "node-category")
	if err != nil {
		res.Err(err, r)
	}
	category := make([]g.Map, 0)
	for _, i := range strings.Split(key, "\n") {
		temp := strings.Split(i, ".")
		category = append(category, map[string]interface{}{
			"value": strings.TrimSpace(temp[0]),
			"label": strings.TrimSpace(temp[1]),
		})
	}
	return category
}

// --- Ws ------------------------------------------------------------------------

type ws struct{}

var Ws = &ws{}

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
