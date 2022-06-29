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

// ---Home----------------------------------------------------------------------
type home struct{}

var Home = &home{}

func (c *home) IndexPage(r *ghttp.Request) {
	r.Response.RedirectTo(g.Config().MustGet(r.Context(), "home").String())
}

// ---Menu-----------------------------------------------------------------

type cMenu struct{ *bo.Search }

var Menu = &cMenu{Search: &bo.Search{
	T1: "s_menu", OrderBy: "t1.sort desc,t1.id desc",
	Fields: []*bo.Field{
		{Name: "pid", QueryName: "menu_pid", SearchType: 1},
		{Name: "name", QueryName: "menu_name", SearchType: 2},
		{Name: "path", QueryName: "menu_path", SearchType: 2},
	},
}}

func (c *cMenu) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cMenu) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/menu/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cMenu) Post(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cMenu) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cMenu) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("menu_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/menu/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cMenu) Put(r *ghttp.Request) {
	d := entity.Menu{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/menu/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

// ---Api-----------------------------------------------------------------

type cApi struct{ *bo.Search }

var Api = &cApi{Search: &bo.Search{
	T1: "s_api", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*bo.Field{
		{Name: "method", SearchType: 1, QueryName: "api_method"}, {Name: "group", SearchType: 2, QueryName: "api_group"}, {Name: "status", SearchType: 1, QueryName: "api_status"},
	},
}}

func (c *cApi) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/sys/api/index.html", g.Map{
		"list": data,
		"page": r.GetPage(total, c.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cApi) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/api/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cApi) Post(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/api/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/api/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cApi) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("api_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/api/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cApi) Put(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/api/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

// ---Role-----------------------------------------------------------------

type (
	cRole     struct{ *bo.Search }
	cRoleApi  struct{ *bo.Search }
	cRoleMenu struct{ *bo.Search }
)

var (
	Role = &cRole{Search: &bo.Search{
		T1: "s_role", OrderBy: "t1.id desc", SearchFields: "t1.*",
		Fields: []*bo.Field{},
	}}
	RoleMenu = &cRoleMenu{Search: &bo.Search{
		T1:           "s_role_menu",
		T2:           "s_role  t2 on t1.rid = t2.id",
		T3:           "s_menu t3 on t1.mid = t3.id",
		OrderBy:      "t1.id desc",
		SearchFields: "t1.*,t2.name role_name ,t3.name menu_name",
		Fields: []*bo.Field{
			{Name: "rid", SearchType: 1},
		},
	}}
	RoleApi = &cRoleApi{Search: &bo.Search{
		T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
		SearchFields: "t1.*,t2.name r_name,t3.url url ,t3.group,t3.method,t3.desc ", Fields: []*bo.Field{
			{Name: "id"},
			{Name: "rid", SearchType: 1},
			{Name: "aid"},
			{Name: "t2.name", QueryName: "r_name", SearchType: 2},
			{Name: "t3.url"},
		},
	}}
)

func (c *cRole) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cRole) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/role/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cRole) Post(r *ghttp.Request) {
	d := entity.Role{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/role/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRole) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/role/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRole) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("role_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/role/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cRole) Put(r *ghttp.Request) {
	d := entity.Role{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/role/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRole) RoleNoMenus(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoMenu(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRole) RoleNoApis(r *ghttp.Request) {
	rid := r.GetQuery("rid")
	data, err := sys.RoleNoApi(r.Context(), rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *cRoleMenu) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	node.Name = "角色菜单"
	node.Path = "/roleMenu/path"
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cRoleMenu) PathAdd(r *ghttp.Request) {
	menus, err := sys.RoleNoMenu(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/roleMenu/add.html", g.Map{"msg": sys.MsgFromSession(r), "menus": menus})
}
func (c *cRoleMenu) Post(r *ghttp.Request) {
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
	r.Response.RedirectTo(fmt.Sprint("/roleMenu/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRoleMenu) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/roleMenu/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRoleApi) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	node.Path = "/roleApi/path"
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cRoleApi) PathAdd(r *ghttp.Request) {
	apis, err := sys.RoleNoApi(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/roleApi/add.html", g.Map{"msg": sys.MsgFromSession(r), "apis": apis})
}
func (c *cRoleApi) Post(r *ghttp.Request) {
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
	r.Response.RedirectTo(fmt.Sprint("/roleApi/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRoleApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/roleApi/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cRoleApi) Clear(r *ghttp.Request) {
	err := sys.ClearRoleApi(r.Context(), r.Get("rid"))
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

// ---Dict-----------------------------------------------------------------

type cDict struct{ *bo.Search }

var Dict = &cDict{Search: &bo.Search{
	T1: "s_dict", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*bo.Field{
		{Name: "k", SearchType: 2, QueryName: "dict_k"}, {Name: "v", SearchType: 2, QueryName: "dict_v"}, {Name: "desc", SearchType: 2, QueryName: "dict_desc"}, {Name: "group", SearchType: 1, QueryName: "dict_group"}, {Name: "status", SearchType: 1, QueryName: "dict_status"}, {Name: "type", SearchType: 1, QueryName: "dict_type"},
	},
}}

func (c *cDict) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cDict) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/dict/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cDict) Post(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/dict/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cDict) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/dict/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cDict) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("dict_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/dict/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cDict) Put(r *ghttp.Request) {
	d := entity.Dict{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/dict/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

type cFile struct{ *bo.Search }

var File = &cFile{Search: &bo.Search{
	T1: "s_file", OrderBy: "t1.id desc", SearchFields: "t1.*",
	Fields: []*bo.Field{
		{Name: "url", SearchType: 2, QueryName: "file_url"}, {Name: "group", SearchType: 1, QueryName: "file_group"}, {Name: "status", SearchType: 1, QueryName: "file_status"},
	},
}}

func (c *cFile) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cFile) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/file/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cFile) Post(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/file/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
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
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/file/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cFile) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("file_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/file/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cFile) Put(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/file/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cFile) Upload(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "上传成功")
	if err := sys.UploadFile(r.Context(), r); err != nil {
		msg = fmt.Sprintf(consts.MsgPrimary, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo("/file/path/add?" + xurl.ToUrlParams(r.GetQueryMap()))
}

type cOperationLog struct{ *bo.Search }

var OperationLog = &cOperationLog{Search: &bo.Search{
	T1: "s_operation_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
	Fields: []*bo.Field{
		{Name: "t2.uname", SearchType: 2, QueryName: "operationLog_uname"}, {Name: "content", SearchType: 2, QueryName: "operationLog_content"},
	},
}}

func (c *cOperationLog) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cOperationLog) PathAdd(r *ghttp.Request) {
	_ = r.Response.WriteTpl("/sys/operationLog/add.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cOperationLog) Post(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/operationLog/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cOperationLog) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/operationLog/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cOperationLog) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("operationLog_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/operationLog/edit.html", g.Map{"msg": sys.MsgFromSession(r)})
}
func (c *cOperationLog) Put(r *ghttp.Request) {
	d := entity.OperationLog{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/operationLog/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cOperationLog) Clear(r *ghttp.Request) {
	msg := fmt.Sprintf(consts.MsgPrimary, "操作成功")
	if err := sys.OperationLogClear(r.Context()); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo("/operationLog/path")
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

//  ---admin-------------------------------------------------------------------

type cAdmin struct{ *bo.Search }

var Admin = &cAdmin{Search: &bo.Search{
	T1: "s_admin", T2: "s_role t2 on t1.rid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.name role_name",
	Fields: []*bo.Field{
		{Name: "rid", SearchType: 1, QueryName: "admin_rid"}, {Name: "status", SearchType: 1, QueryName: "admin_status"},
	},
}}

func (c *cAdmin) Path(r *ghttp.Request) {
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
	}); err != nil {
		res.Err(err, r)
	}
}
func (c *cAdmin) PathAdd(r *ghttp.Request) {
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Response.WriteTpl("/sys/admin/add.html", g.Map{"msg": sys.MsgFromSession(r), "roles": roles})
}
func (c *cAdmin) Post(r *ghttp.Request) {
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
	r.Response.RedirectTo(fmt.Sprint("/admin/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cAdmin) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c *cAdmin) PathEdit(r *ghttp.Request) {
	roles, err := sys.Roles(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	g.Log().Notice(nil, data["rid"])

	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("admin_edit", data.Map())
	_ = r.Response.WriteTpl("/sys/admin/edit.html", g.Map{"msg": sys.MsgFromSession(r), "roles": roles})
}
func (c *cAdmin) Put(r *ghttp.Request) {
	d := entity.Admin{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, &d); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/admin/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
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

// ---Gen Code-------------------------------------------------------------------

var Gen = &gen{}

type gen struct{}

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
	g.Dump(d)
	if err := sys.GenFile(r.Context(), &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
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
