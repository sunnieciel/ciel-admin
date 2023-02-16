package controller

import (
	v1 "ciel-admin/api/v1"
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xcaptcha"
	"ciel-admin/utility/utils/xfile"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	captcha "github.com/mojocn/base64Captcha"
	"math"
)

var (
	Menu                   = &cMenu{cBase{"s_menu", "/admin/menu", "/sys/menu", "sys"}}
	Api                    = &cApi{cBase{"s_api", "/admin/api", "/sys/api", "sys"}}
	Role                   = &cRole{cBase{"s_role", "/admin/role", "/sys/role", "sys"}}
	RoleApi                = &cRoleApi{cBase{Table: "s_role_api", ReqPath: "/admin/roleApi", FileDir: "/sys/roleApi"}}
	RoleMenu               = &cRoleMenu{cBase{Table: "s_role_menu", ReqPath: "/admin/roleMenu", FileDir: "/sys/roleMenu"}}
	Admin                  = &cAdmin{cBase{Table: "s_admin", ReqPath: "/admin/admin", FileDir: "/sys/admin/"}}
	AdminLoginLog          = &cAdminLoginLog{cBase{Table: "s_admin_login_log", ReqPath: "/admin/adminLoginLog", FileDir: "/sys/adminLoginLog"}}
	OperationLog           = &cOperationLog{cBase{Table: "s_operation_log", ReqPath: "/admin/operationLog", FileDir: "/sys/operationLog"}}
	AdminMessage           = &cAdminMessage{cBase{"s_admin_message", "/admin/adminMessage", "/sys/adminMessage", "sys"}}
	Dict                   = &cDict{cBase{Table: "s_dict", ReqPath: "/admin/dict", FileDir: "/sys/dict"}}
	File                   = &cFile{cBase{Table: "s_file", ReqPath: "/admin/file", FileDir: "/sys/file"}}
	Gen                    = &cGen{}
	Sys                    = &cSys{}
	Ws                     = &cWs{}
	Banner                 = &cBanner{cBase{"c_banner", "/admin/banner", "/common/banner", "sys"}}
	User                   = &cUser{cBase{"u_user", "/admin/user", "/user/user", "sys"}}
	UserLoginLog           = &cUserLoginLog{cBase{"u_user_login_log", "/admin/userLoginLog", "/user/userLoginLog", "sys"}}
	Wallet                 = &cWallet{cBase{"u_wallet", "/admin/wallet", "/user/wallet", "sys"}}
	WalletChangeLog        = &cWalletChangeLog{cBase{"u_wallet_change_log", "/admin/walletChangeLog", "/user/walletChangeLog", "sys"}}
	WalletReport           = &cWalletReport{}
	WalletStatisticsLog    = &cWalletStatisticsLog{cBase{"u_wallet_statistics_log", "/admin/walletStatisticsLog", "/statistics/walletStatisticsLog", "sys"}}
	WalletChangeType       = &cWalletChangeType{cBase{"u_wallet_change_type", "/admin/walletChangeType", "/user/walletChangeType", "sys"}}
	WalletTopUpApplication = &cWalletTopUpApplication{cBase{"u_wallet_top_up_application", "/admin/walletTopUpApplication", "/user/walletTopUpApplication", "sys"}}
)

type cBase struct {
	Table   string
	ReqPath string
	FileDir string
	DBGroup string
}

type cMenu struct{ cBase }

func (c cMenu) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/menu", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)             // 主页面
		g.GET("/add", c.IndexAdd)       // 添加页面
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)                  // 删除请求
		g.POST("/add", c.Add)                     // 添加请求
		g.POST("/update", c.Update)               // 修改请求
		g.PUT("/setGroupSort", c.UpdateGroupSort) // 设置分组排序
	})
}
func (c cMenu) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, OrderBy: "t1.sort ,t1.id desc", Fields: []model.Field{
			{Name: "pid", Type: 1},
			{Name: "name", Type: 2},
			{Name: "path", Type: 2},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r, 50)
	if s.Size == 10 {
		s.Size = 50
	}
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cMenu) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cMenu) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cMenu) Add(r *ghttp.Request) {
	var (
		d = entity.Menu{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) Update(r *ghttp.Request) {
	var (
		d     = entity.Menu{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cMenu) UpdateGroupSort(r *ghttp.Request) {
	var (
		ctx = r.Context()
		d   struct {
			Id   uint64
			Sort int
		}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System.UpdateMenuSort(ctx, d.Sort, d.Id); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

type cApi struct{ cBase }

func (c cApi) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/api", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
	})
}
func (c cApi) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.Request.URL.Path
		s       = model.Search{
			T1: c.Table, OrderBy: "t1.group,t1.type,t1.id desc",
			Fields: []model.Field{
				{Name: "method", Type: 1},
				{Name: "group", Type: 2},
				{Name: "type", Type: 1},
				{Name: "desc", Type: 2},
			},
		}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	apiGroup, err := service.Dict.GetApiGroupOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/index.html"), g.Map{
		"list":      data,
		"page":      r.GetPage(total, s.Size).GetContent(3),
		"node":      node,
		"msg":       service.System.GetMsgFromSession(r),
		"path":      reqPath,
		"api_group": apiGroup,
	}, r)
}
func (c cApi) IndexAdd(r *ghttp.Request) {
	apiGroup, err := service.Dict.GetApiGroupOptions(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{
		"msg":       service.System.GetMsgFromSession(r),
		"api_group": apiGroup,
	}, r)
}
func (c cApi) IndexEdit(r *ghttp.Request) {
	data, err := service.System.GetById(r.Context(), c.Table, xparam.ID(r), "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	apiGroup, err := service.Dict.GetApiGroupOptions(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/edit.html"), g.Map{"msg": service.System.GetMsgFromSession(r), "api_group": apiGroup}, r)
}
func (c cApi) Add(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap())), r)
}
func (c cApi) Del(r *ghttp.Request) {
	id := r.Get("id")
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), c.Table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap())), r)
}
func (c cApi) Update(r *ghttp.Request) {
	d := entity.Api{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), c.Table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())), r)
}

type cRole struct{ cBase }

func (c cRole) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/role", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.GET("/nomenus", c.ListRoleNoMenus)
		g.GET("/noapis", c.ListRoleNoApis)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/clear/:id", c.DelRoles)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
	})
}
func (c cRole) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		s       = model.Search{T1: "s_role", OrderBy: "t1.id desc", SearchFields: "t1.*"}
		reqPath = r.URL.Path
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(fmt.Sprint(c.FileDir, "/index.html"), g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  service.System.GetMsgFromSession(r),
		"path": reqPath,
	}, r)
}
func (c cRole) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cRole) IndexEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		file  = fmt.Sprint(c.FileDir, "/edit.html")
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		id    = xparam.ID(r)
	)
	data, err := service.System.GetById(ctx, table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(file, d, r)
}
func (c cRole) Add(r *ghttp.Request) {
	var (
		d     = entity.Role{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(ctx, table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cRole) ListRoleNoMenus(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid")
		ctx = r.Context()
	)
	data, err := service.Role.ListRoleNoMenus(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c cRole) ListRoleNoApis(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid")
		ctx = r.Context()
	)
	data, err := service.Role.ListRoleNoApis(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c cRole) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
		ctx   = r.Context()
		path  = fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cRole) DelRoles(r *ghttp.Request) {
	var (
		rid = r.GetQuery("rid").Int()
		ctx = r.Context()
	)
	res.OkSession("清除成功", r)
	if err := service.Role.DelMenus(ctx, rid); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo("", r)
}
func (c cRole) Update(r *ghttp.Request) {
	var (
		d     = entity.Role{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%v?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cRoleApi struct{ cBase }

func (c cRoleApi) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/roleApi", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.GET("/clear/:rid", c.DelRoleApis)
		g.POST("/add", c.Add)
	})
}
func (c cRoleApi) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = model.Search{
			T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
			OrderBy:      "t3.group,t3.type",
			SearchFields: "t1.*,t2.name r_name,t3.url url,t3.group,t3.method,t3.desc,t3.type", Fields: []model.Field{
				{Name: "id"},
				{Name: "rid", Type: 1},
				{Name: "aid"},
				{Name: "t3.group", QueryName: "group", Type: 1},
				{Name: "t3.type", QueryName: "type", Type: 1},
				{Name: "t2.name", QueryName: "r_name", Type: 2},
				{Name: "t3.url"},
			},
		}
		path = r.URL.Path
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		rid  = r.GetQuery("rid").Int()
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	roleInfo, err := service.Role.GetById(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	node.Name = "角色禁用API"
	node.Desc = fmt.Sprintf("如果你不希望<span class='color-red strong'>【%s】</span>角色访问某些api功能，可以点击添加按钮，将他们添加到下面的列表中", roleInfo.Name)
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(r.Context(), s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	groups, err := service.Dict.GetApiGroupOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl(file, g.Map{
		"list":   data,
		"page":   r.GetPage(total, s.Size).GetContent(3),
		"node":   node,
		"msg":    service.System.GetMsgFromSession(r),
		"path":   path,
		"groups": groups,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cRoleApi) IndexAdd(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		rid  = r.Get("rid")
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	apis, err := service.Role.ListRoleNoApis(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	roleData, err := service.Role.GetById(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}

	groups, err := service.Dict.GetApiGroupOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	d := g.Map{
		"msg":    service.System.GetMsgFromSession(r),
		"apis":   apis,
		"role":   roleData,
		"groups": groups,
	}
	_ = r.Response.WriteTpl(file, d)
}
func (c cRoleApi) Add(r *ghttp.Request) {
	var (
		d struct {
			Rid int
			Aid []int
		}
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	_ = r.Parse(&d)
	res.OkSession("添加成功", r)
	if err := service.Role.AddApi(ctx, d.Rid, d.Aid); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cRoleApi) Del(r *ghttp.Request) {
	var (
		id   = r.Get("id")
		ctx  = r.Context()
		path = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, c.Table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	g.Log().Infof(ctx, path)
	res.RedirectTo(path, r)
}
func (c cRoleApi) DelRoleApis(r *ghttp.Request) {
	var (
		ctx = r.Context()
		rid = r.Get("rid")
		t   = r.Get("type").Int()
	)
	err := service.Role.DelApis(ctx, rid, t)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

type cRoleMenu struct{ cBase }

func (c cRoleMenu) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/roleMenu", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.IndexPath)
		g.GET("/add", c.IndexPathAdd)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
	})
}
func (c cRoleMenu) IndexPath(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = model.Search{
			T1:           "s_role_menu",
			T2:           "s_role  t2 on t1.rid = t2.id",
			T3:           "s_menu t3 on t1.mid = t3.id",
			OrderBy:      "t3.sort ",
			SearchFields: "t1.*,t2.name role_name ,t3.name menu_name,t3.pid ",
			Fields: []model.Field{
				{Name: "rid", Type: 1},
			},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
		msg  = service.System.GetMsgFromSession(r)
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}

	node.Name = "角色菜单"
	node.Path = c.ReqPath
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  msg,
		"path": path,
	}, r)
}
func (c cRoleMenu) IndexPathAdd(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		rid  = r.Get("rid")
		file = fmt.Sprintf("%s/add.html", c.FileDir)
		msg  = service.System.GetMsgFromSession(r)
	)
	menus, err := service.Role.ListRoleNoMenus(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	roleData, err := service.Role.GetById(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	for _, i := range menus {
		g.Log().Info(ctx, i)
	}
	res.Tpl(file, g.Map{"msg": msg, "menus": menus, "role": roleData}, r)
}
func (c cRoleMenu) Del(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		id    = r.Get("id")
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cRoleMenu) Add(r *ghttp.Request) {
	var (
		d struct {
			Rid int
			Mid []int
		}
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	_ = r.ParseForm(&d)
	res.OkSession("添加成功", r)
	if err := service.Role.AddMenu(ctx, d.Rid, d.Mid); err != nil {
		res.ErrSession(err, r)
	}
	r.Response.RedirectTo(path)
}

type cAdmin struct{ cBase }

func (c cAdmin) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/admin", func(g *ghttp.RouterGroup) {
		g.GET("/getCaptcha", Sys.GetCaptcha) // 获取验证码
		g.POST("/login", c.Login)
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/logout", c.Logout)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.GET("/notifications", c.IndexNotifications) // Admin notifications page
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/notifications/del/:id", c.DelNotification)
		g.GET("/notifications/clear", c.DelNotifications)
		g.PUT("/updatePwd", c.UpdatePwd)
		g.PUT("/updatePwdWithoutOldPwd", c.UpdatePwdWithoutOldPwd)
		g.PUT("/updateUname", c.UpdateUname)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
	})
}
func (c cAdmin) Index(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		s     = model.Search{
			T1: table, T2: "s_role t2 on t1.rid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.name role_name",
			Fields: []model.Field{{Name: "rid", Type: 1}, {Name: "status", Type: 1}},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
	)
	roles, err := service.Role.GetRoleOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list":  data,
		"page":  r.GetPage(total, s.Size).GetContent(3),
		"node":  node,
		"msg":   service.System.GetMsgFromSession(r),
		"roles": roles,
		"path":  r.URL.Path,
	}, r)
}
func (c cAdmin) IndexAdd(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	roles, err := service.Role.GetRoleOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r), "roles": roles}, r)
}
func (c cAdmin) IndexEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
		id    = xparam.ID(r)
	)
	roles, err := service.Role.GetRoleOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	data, err := service.System.GetById(ctx, table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r), "roles": roles}, r)
}
func (c cAdmin) IndexLogin(r *ghttp.Request) {
	res.Page(r, "login.html")
}
func (c cAdmin) IndexNotifications(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	page, size := res.GetPage(r, 50)
	total, list, err := service.Admin.ListNotifications(ctx, page, size)
	if err != nil {
		res.Err(err, r)
	}
	p := r.GetPage(int(total), size)
	res.Tpl("sys/admin/notifications.html", g.Map{
		"msg":       service.System.GetMsgFromSession(r),
		"total":     total,
		"list":      list,
		"totalPage": p.TotalPage,
		"page":      p.GetContent(3),
	}, r)
}
func (c cAdmin) Add(r *ghttp.Request) {
	var (
		d    = entity.Admin{}
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.Admin.Add(ctx, d); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdmin) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdmin) DelNotification(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		id   = r.Get("id").Uint64()
		path = fmt.Sprintf("/admin/admin/notifications?%s", xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.Admin.DelNotification(ctx, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdmin) DelNotifications(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	if err := service.Admin.DelNotifications(ctx); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
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
	if err := service.Admin.Login(r.Context(), d.Id, d.Code, d.Uname, d.Pwd, r.GetClientIp()); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) Logout(r *ghttp.Request) {
	err := service.Admin.Logout(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cAdmin) Update(r *ghttp.Request) {
	var (
		d     = entity.Admin{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "pwd")
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cAdmin) UpdatePwd(r *ghttp.Request) {
	var d struct {
		OldPwd string `v:"required"`
		NewPwd string `v:"required"`
	}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Admin.UpdatePwd(r.Context(), d.OldPwd, d.NewPwd); err != nil {
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
	if err := service.Admin.UpdateUname(r.Context(), d.Id, d.Uname); err != nil {
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
	if err := service.Admin.UpdatePwdWithoutOld(r.Context(), d.Id, d.Pwd); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

type cAdminLoginLog struct{ cBase }

func (c cAdminLoginLog) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/adminLoginLog", func(g *ghttp.RouterGroup) {

		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
		g.GET("/clear", c.DelAdminLoginLogs)
	})
}
func (c cAdminLoginLog) Index(r *ghttp.Request) {
	var (
		s = model.Search{
			T1: "s_admin_login_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
			Fields: []model.Field{
				{Name: "uid", Type: 1}, {Name: "t2.uname", Type: 2, QueryName: "uname"},
			},
		}
		ctx  = r.Context()
		path = r.URL.Path
		file = fmt.Sprintf("%s/index.html", c.FileDir)
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  service.System.GetMsgFromSession(r),
		"path": r.URL.Path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cAdminLoginLog) IndexAdd(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cAdminLoginLog) IndexEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
		table = c.Table
		id    = xparam.ID(r)
	)
	data, err := service.System.GetById(ctx, table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cAdminLoginLog) Add(r *ghttp.Request) {
	var (
		d     = entity.AdminLoginLog{}
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(ctx, table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdminLoginLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
		ctx   = r.Context()
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdminLoginLog) DelAdminLoginLogs(r *ghttp.Request) {
	var (
		path = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("操作成功", r)
	if err := service.Admin.DelLoginLogs(r.Context()); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cAdminLoginLog) Update(r *ghttp.Request) {
	var (
		d     = entity.AdminLoginLog{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cOperationLog struct{ cBase }

func (c cOperationLog) RegisterRouter(g *ghttp.RouterGroup) {

	g.Group("/operationLog", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
		g.GET("/clear", c.DelOperationLogs)
	})
}
func (c cOperationLog) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = model.Search{
			T1: "s_operation_log", T2: "s_admin t2 on t1.uid = t2.id", OrderBy: "t1.id desc", SearchFields: "t1.*,t2.uname",
			Fields: []model.Field{
				{Name: "t2.uname", Type: 2, QueryName: "uname"}, {Name: "content", Type: 2},
			},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
		msg  = service.System.GetMsgFromSession(r)
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(r.Context(), s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  msg,
		"path": path,
	}, r)
}
func (c cOperationLog) IndexAdd(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
		d    = g.Map{"msg": service.System.GetMsgFromSession(r)}
	)
	res.Tpl(file, d, r)
}
func (c cOperationLog) IndexEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		id    = xparam.ID(r)
		table = c.Table
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
	)
	data, err := service.System.GetById(ctx, table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cOperationLog) Add(r *ghttp.Request) {
	var (
		d     = entity.OperationLog{}
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(ctx, table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cOperationLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		ctx   = r.Context()
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cOperationLog) DelOperationLogs(r *ghttp.Request) {
	res.OkSession("操作成功", r)
	if err := service.Admin.DelOperationLogs(r.Context()); err != nil {
		res.ErrSession(err, r)
	}
	path := c.ReqPath
	res.RedirectTo(path, r)
}
func (c cOperationLog) Update(r *ghttp.Request) {
	var (
		d     = entity.OperationLog{}
		ctx   = r.Context()
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("%s/edit/%d/?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cDict struct{ cBase }

func (c cDict) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/dict", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
	})
}
func (c cDict) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = model.Search{
			T1: "s_dict", OrderBy: "t1.group,t1.id desc", SearchFields: "t1.*", Fields: []model.Field{
				{Name: "k", Type: 2}, {Name: "v", Type: 2}, {Name: "desc", Type: 2}, {Name: "group", Type: 1}, {Name: "status", Type: 1}, {Name: "type", Type: 1},
			},
		}
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  service.System.GetMsgFromSession(r),
		"path": path,
	}); err != nil {
		res.Err(err, r)
	}
}
func (c cDict) IndexAdd(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cDict) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		ctx   = r.Context()
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
		id    = xparam.ID(r)
	)
	data, err := service.System.GetById(ctx, table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cDict) Add(r *ghttp.Request) {
	var (
		d     = entity.Dict{}
		ctx   = r.Context()
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(ctx, table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cDict) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cDict) Update(r *ghttp.Request) {
	var (
		d     = entity.Dict{}
		table = c.Table
		ctx   = r.Context()
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	if d.K == "white_ips" {
		if err := service.Dict.UpdateWhiteIps(r.Context(), d.V); err != nil {
			res.Err(err, r)
		}
	}
	path := fmt.Sprintf("%s/edit/%d?%s", c.ReqPath, d.Id, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cFile struct{ cBase }

func (c cFile) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/file", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
		g.POST("/upload", c.Upload)
	})
}
func (c cFile) Index(r *ghttp.Request) {
	var (
		s = model.Search{
			T1: "s_file", OrderBy: "t1.id desc", SearchFields: "t1.*", Fields: []model.Field{
				{Name: "url", Type: 2}, {Name: "group", Type: 1}, {Name: "status", Type: 1},
			},
		}
		ctx  = r.Context()
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(r.Context(), s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  service.System.GetMsgFromSession(r),
		"path": path,
	}, r)
}
func (c cFile) IndexAdd(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cFile) IndexEdit(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		id    = xparam.ID(r)
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
	)
	data, err := service.System.GetById(ctx, table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cFile) Add(r *ghttp.Request) {
	var (
		d     = entity.File{}
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(ctx, table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cFile) Upload(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("上传成功", r)
	if err := service.File.Uploads(ctx, r); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cFile) Del(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		p     = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
		id    = xparam.ID(r)
	)
	f, err := service.File.GetById(ctx, id)
	if err != nil {
		res.Err(err, r)
	}
	path, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		res.Err(err, r)
	}
	filePath := gfile.Pwd() + path.String() + "/" + f.Url
	if err = xfile.Remove(ctx, filePath); err != nil {
		res.Err(err, r)
	}
	res.OkSession("删除成功", r)
	if err = service.System.Del(ctx, table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(p, r)
}
func (c cFile) Update(r *ghttp.Request) {
	var (
		d     = entity.File{}
		ctx   = r.Context()
		table = c.Table
		id    = r.GetForm("id").Uint()
		path  = fmt.Sprintf("%s/edit/%v?%s", c.ReqPath, id, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(ctx, table, id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}

type cGen struct{}

func (c cGen) Index(r *ghttp.Request) {
	var (
		path = r.URL.Path
		ctx  = r.Context()
		db   = r.Get("db").String()
	)
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	tables, err := service.Gen.GetTables(ctx, db)
	if err != nil {
		res.ErrSession(err, r)
	}

	menuLeve1, err := service.Gen.GenMenuLevel1(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl("/sys/gen/index.html", g.Map{
		"node":       node,
		"path":       path,
		"msg":        service.System.GetMsgFromSession(r),
		"menuLevel1": menuLeve1,
		"tables":     tables,
	}, r)
}
func (c cGen) Gen(r *ghttp.Request) {
	var (
		ctx = r.Context()
		d   struct {
			Table     string `v:"required#请选择表名"`
			Group     string `v:"required#分组不能为空"`
			Menu      string `v:"required#菜单名不能为空"`
			Prefix    string
			ApiGroup  string `v:"required#API分组不能为空"`
			HtmlGroup string `v:"required#html文件文件夹分组不能为空"`
			Db        string
		}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Gen.Gen(ctx, d.Table, d.Group, d.Menu, d.Prefix, d.ApiGroup, d.HtmlGroup, d.Db); err != nil {
		res.Err(err, r)
	}
	res.OkMsg("生成成功", r)
}
func (c cGen) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/gen", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.POST("/table", c.Gen)
	})
}

type cAdminMessage struct{ cBase }

func (c cAdminMessage) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/adminMessage", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index) // 主页面
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
	})
}
func (c cAdminMessage) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, OrderBy: "t1.id desc", Fields: []model.Field{
			{Name: "id", Type: 1},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cAdminMessage) IndexAdd(r *ghttp.Request) {
	var file = fmt.Sprintf("%s/add.html", c.FileDir)
	res.Tpl(file, g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cAdminMessage) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cAdminMessage) Add(r *ghttp.Request) {
	var (
		d = struct {
			Uname string
			Title string
			Url   string
			Type  int
		}{}
		ctx = r.Context()
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	d.Type = 1
	res.OkSession("添加成功", r)
	if err := service.Admin.AddMessage(ctx, d.Uname, d.Title, d.Url, d.Type); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cAdminMessage) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cAdminMessage) Update(r *ghttp.Request) {
	var (
		d     = entity.AdminMessage{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cSys struct{}

func (c cSys) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/sys", func(g *ghttp.RouterGroup) {
		g.GET("/noticeAdmin", Ws.NoticeAdmin)
		g.GET("/document", c.IndexDocument)
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/ws", Ws.GetAdminWs)
	})
}
func (c cSys) IndexDocument(r *ghttp.Request) {
	res.Tpl("/sys/tool/document.html", nil, r)
}

func (c cSys) UploadImg(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	return service.File.Upload(ctx, req.Group)
}
func (c cSys) GetDictByKey(ctx context.Context, req *v1.DictReq) (res *v1.DictRes, err error) {
	data, err := service.Dict.GetByKey(ctx, req.Key)
	if err != nil {
		return nil, err
	}
	return &v1.DictRes{Value: data}, nil
}
func (c cSys) GetCaptcha(r *ghttp.Request) {
	var driver = xcaptcha.NewDriver().ConvertFonts()
	cc := captcha.NewCaptcha(driver, xcaptcha.Store)
	_, content, answer := cc.Driver.GenerateIdQuestionAnswer()
	id := r.GetQuery("id").String()
	item, _ := cc.Driver.DrawCaptcha(content)
	_ = cc.Store.Set(id, answer)
	res.OkData(item.EncodeB64string(), r)
}
func (c cSys) ListAllDict(ctx context.Context, req *v1.AllDictReq) (res v1.AllDictRes, err error) {
	allDict, err := service.System.ListAllDict(ctx)
	if err != nil {
		return nil, err
	}
	res = allDict
	return
}

type cWs struct{}

func (w cWs) GetUserWs(r *ghttp.Request) {
	service.Ws.GetUserWs(r)
}
func (w cWs) GetAdminWs(r *ghttp.Request) {
	service.Ws.GetAdminWs(r)
}
func (w cWs) NoticeUser(r *ghttp.Request) {
	var d struct {
		Uid     int `v:"required"`
		OrderId int `v:"required"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = service.Ws.NoticeUser(gctx.New(), d.Uid, d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (w cWs) NoticeAdmin(r *ghttp.Request) {
	var (
		d struct {
			Msg string `v:"required" json:"msg"`
		}
		ctx = r.Context()
	)
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	err = service.Ws.NoticeAdmins(ctx, d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

type cBanner struct{ cBase }

func (c cBanner) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/banner", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)             // 主页面
		g.GET("/add", c.IndexAdd)       // 添加页面
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
	})
}
func (c cBanner) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, OrderBy: "t1.id desc", Fields: []model.Field{
			{Name: "id", Type: 1},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cBanner) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cBanner) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cBanner) Add(r *ghttp.Request) {
	var (
		d = entity.Banner{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cBanner) List(ctx context.Context, req *v1.BannersReq) (res []*v1.BannerRes, err error) {
	return service.System.ListBanners(ctx)
}
func (c cBanner) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cBanner) Update(r *ghttp.Request) {
	var (
		d     = entity.Banner{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

type cUser struct{ cBase }

func (c cUser) RegisterRouter(s *ghttp.RouterGroup) {
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)
		g.GET("/add", c.IndexAdd)
		g.GET("/edit/:id", c.IndexEdit)
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)
		g.POST("/add", c.Add)
		g.POST("/update", c.Update)
		g.PUT("/updateUname", c.UpdateUname)
		g.PUT("/updatePass", c.UpdatePassByAdmin)
	})
}
func (c cUser) RegisterWebApi(s *ghttp.RouterGroup) {
	s.Group("/user", func(g *ghttp.RouterGroup) {
		g.PUT("/updateIcon", c.UpdateIcon) // 修改头像
	})
}
func (c cUser) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, OrderBy: "t1.id desc", Fields: []model.Field{
			{Name: "id", Type: 1},
			{Name: "uname", Type: 2},
			{Name: "status", Type: 1},
			{Name: "desc", Type: 2},
			{Name: "join_ip", Type: 2},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cUser) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cUser) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, "sys")
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cUser) IndexLogin(r *ghttp.Request) {
	res.Page(r, "/web/user_login.html")
}
func (c cUser) Add(r *ghttp.Request) {
	var (
		d = entity.User{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cUser) Del(r *ghttp.Request) {
	var (
		ctx = r.Context()
		id  = r.Get("id").Uint64()
	)
	res.OkSession("删除成功", r)
	if err := service.User.Del(ctx, id); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cUser) Update(r *ghttp.Request) {
	var (
		d     = entity.User{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	delete(m, "pass")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}

func (c cUser) Register(ctx context.Context, req *v1.RegisterReq) (*v1.LoginRes, error) {
	if len(req.Uname) < 3 || len(req.Uname) > 12 {
		return nil, consts.ErrUnameFormat
	}
	return service.User.Register(ctx, req)
}
func (c cUser) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	return service.User.Login(ctx, req)
}
func (c cUser) UpdateUname(r *ghttp.Request) {
	var (
		ctx = r.Context()
		d   struct {
			Uname string `v:"required"`
			Id    uint64 `v:"required"`
		}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.User.UpdateUname(ctx, d.Uname, d.Id); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cUser) UpdatePassByAdmin(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		pass = r.GetForm("pass").String()
		id   = r.GetForm("id").Uint64()
	)
	if err := service.User.UpdatePassByAdmin(ctx, pass, id); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cUser) UpdatePassByUser(ctx context.Context, req *v1.UpdatePassReq) (res *v1.DefaultRes, err error) {
	if err := service.User.UpdatePassByUser(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cUser) UpdateNickname(ctx context.Context, req *v1.UpdateNicknameReq) (res *v1.DefaultRes, err error) {
	if err := service.User.UpdateNickname(ctx, req); err != nil {
		return nil, err
	}
	return
}
func (c cUser) Icons(ctx context.Context, req *v1.IconsReq) (res *v1.IconsRes, err error) {
	icons, err := service.User.ListIcons(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.IconsRes{Icons: icons, ImgPrefix: consts.ImgPrefix}
	return
}
func (c cUser) UpdateIcon(ctx context.Context, req *v1.UpdateIconReq) (res *v1.DefaultRes, err error) {
	if err = service.User.UpdateIcon(ctx, req.Icon); err != nil {
		return nil, err
	}
	return
}
func (c cUser) RegisterIndex(r *ghttp.Request) {
	res.Page(r, "/web/user_register.html")
}
func (c cUser) Logout(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	service.User.Logout(ctx)
	res.RedirectTo("/", r)
}

func (c cUser) Info(ctx context.Context, req *v1.UserInfoReq) (*v1.LoginRes, error) {
	return service.User.GetUserInfo(ctx)
}

type cUserLoginLog struct{ cBase }

func (c cUserLoginLog) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{
			T1:           c.Table,
			T2:           "u_user t2 on t1.uid = t2.id",
			SearchFields: "t1.*,t2.uname uname",
			OrderBy:      "t1.id desc", Fields: []model.Field{
				{Name: "id", Type: 1},
				{Name: "ip", Type: 2},
				{Name: "t2.uname", QueryName: "uname", Type: 1},
			}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, "sys")
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cUserLoginLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, "sys"); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cUserLoginLog) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/userLoginLog", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index) // 主页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del) // 删除请求
		g.GET("/clear", c.Clear)
	})
}
func (c cUserLoginLog) Clear(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	res.OkSession("ok", r)
	if err := service.User.DelLoginLogs(ctx); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo("/admin/userLoginLog", r)
}

type cWallet struct{ cBase }

func (c cWallet) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/wallet", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)                 // 主页面
		g.GET("/edit/:id", c.IndexEdit)     // 修改页面
		g.GET("/topUp/:id", c.TopUpIndex)   // 充值页面
		g.GET("/deduct/:id", c.DeductIndex) //扣除页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
		g.PUT("/updatePassByAdmin", c.UpdatePassByAdmin)
		g.POST("/topUpByAdmin", c.TopUpByAdmin)   // 充值金币
		g.POST("/deductByAdmin", c.DeductByAdmin) // 扣除用户金币
	})
}

func (c cWallet) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{
			T1:           c.Table,
			T2:           "u_user t2 on t1.uid = t2.id",
			SearchFields: "t1.*,t2.uname",
			OrderBy:      "t1.id desc", Fields: []model.Field{
				{Name: "id", Type: 1},
				{Name: "status", Type: 1},
				{Name: "t2.uname", QueryName: "uname", Type: 2},
			}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cWallet) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cWallet) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cWallet) TopUpIndex(r *ghttp.Request) {
	var (
		id  = r.Get("id").Uint64()
		ctx = r.Context()
	)
	userInfo, err := service.User.GetById(ctx, id)
	if err != nil {
		res.ErrSession(err, r)
	}
	walletInfo, err := service.Wallet.GetById(ctx, id)
	if err != nil {
		res.ErrSession(err, r)
	}
	topUpOptions, err := service.Wallet.GetChangeTypeTopUpOptions(ctx)
	if err != nil {
		res.ErrSession(err, r)
	}
	res.Tpl("/user/wallet/topUp.html", g.Map{
		"user":         userInfo,
		"walletInfo":   walletInfo,
		"msg":          service.System.GetMsgFromSession(r),
		"topUpOptions": topUpOptions,
	}, r)
}
func (c cWallet) DeductIndex(r *ghttp.Request) {
	var (
		id  = r.Get("id").Uint64()
		ctx = r.Context()
	)
	userInfo, err := service.User.GetById(ctx, id)
	if err != nil {
		res.ErrSession(err, r)
	}
	walletInfo, err := service.Wallet.GetById(ctx, id)
	if err != nil {
		res.ErrSession(err, r)
	}
	deductOptions, err := service.Wallet.GetChangeTypeDeductOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl("/user/wallet/deduct.html", g.Map{
		"user":          userInfo,
		"walletInfo":    walletInfo,
		"msg":           service.System.GetMsgFromSession(r),
		"deductOptions": deductOptions,
	}, r)
}
func (c cWallet) Add(r *ghttp.Request) {
	var (
		d = entity.Wallet{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWallet) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWallet) Update(r *ghttp.Request) {
	var (
		d     = entity.Wallet{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWallet) UpdatePassByAdmin(r *ghttp.Request) {
	var (
		d struct {
			Uid  uint64
			Pass string
		}
		ctx = r.Context()
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.Wallet.UpdatePassByAdmin(ctx, d.Pass, d.Uid); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c cWallet) TopUpByAdmin(r *ghttp.Request) {
	var (
		d struct {
			Type   int
			Uid    uint64
			Wallet float64
			Desc   string
		}
		ctx = r.Context()
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("充值成功", r)
	if err := service.Wallet.UpdateTopUpByAdmin(ctx, d.Type, d.Uid, d.Wallet, d.Desc); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("/admin/wallet/topUp/%d?%s", d.Uid, xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWallet) DeductByAdmin(r *ghttp.Request) {
	var (
		d struct {
			Type   int
			Uid    uint64
			Amount float64
		}
		ctx  = r.Context()
		path = fmt.Sprintf("/admin/wallet/deduct/%d?%s", r.GetForm("uid").Uint64(), xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("扣除成功", r)
	if err := r.Parse(&d); err != nil {
		res.ErrSession(err, r)
		goto here
	}
	if err := service.Wallet.UpdateDeductByAdmin(ctx, d.Type, d.Uid, math.Abs(d.Amount)*-1); err != nil {
		res.ErrSession(err, r)
	}
here:
	res.RedirectTo(path, r)
}
func (c cWallet) SetPass(ctx context.Context, req *v1.WalletSetPassReq) (res *v1.DefaultRes, err error) {
	if err = service.Wallet.UpdateSetPass(ctx, req, service.User.GetUidFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cWallet) UpdatePass(ctx context.Context, req *v1.WalletUpdatePassReq) (res *v1.DefaultRes, err error) {
	if err = service.Wallet.UpdatePass(ctx, req, service.User.GetUidFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cWallet) TopUpCategory(ctx context.Context, req *v1.TopUpCategoryReq) (res []*v1.TopUpCategoryRes, err error) {
	return service.Wallet.ListTopUpCategory(ctx)
}
func (c cWallet) CreateTopUp(ctx context.Context, req *v1.CreateTopUpReq) (res *v1.DefaultRes, err error) {
	if err = service.Wallet.AddTopUp(ctx, req, service.User.GetUidFromCtx(ctx)); err != nil {
		return nil, err
	}
	return
}
func (c cWallet) ListTopUp(ctx context.Context, req *v1.ListTopUpReq) (res *v1.ListTopUpRes, err error) {
	items, page, err := service.Wallet.ListTopUp(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.ListTopUpRes{
		PageRes: page,
		Items:   items,
	}
	return
}
func (c cWallet) ListChangeTypes(ctx context.Context, req *v1.ListChangeTypesReq) (res []*v1.ListChangeTypesRes, err error) {
	return service.Wallet.ListChangeTypes(ctx)
}
func (c cWallet) ListChangeLogs(ctx context.Context, req *v1.ListChangeLogReq) (res *v1.ListChangeLogRes, err error) {
	items, pageRes, err := service.Wallet.ListChangeLogs(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.ListChangeLogRes{
		PageRes: pageRes,
		Items:   items,
	}
	return
}
func (c cWallet) GetInfo(ctx context.Context, req *v1.WalletInfoReq) (res *v1.WalletInfoRes, err error) {
	return service.Wallet.GetInfo(ctx)
}

type cWalletChangeLog struct{ cBase }

func (c cWalletChangeLog) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/walletChangeLog", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)             // 主页面
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/clear", c.Clear)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
	})
}
func (c cWalletChangeLog) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{
			T1: c.Table, T2: "u_user t2 on t1.uid = t2.id", SearchFields: "t1.*,t2.uname", OrderBy: "t1.id desc", Fields: []model.Field{
				{Name: "trans_id", Type: 2},
				{Name: "desc", Type: 2},
				{Name: "type", Type: 1},
				{Name: "t2.uname", QueryName: "uname", Type: 2},
			}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	changeTypeOptions, err := service.Wallet.GetChangeTypeOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node":              node,
		"list":              data,
		"page":              r.GetPage(total, s.Size).GetContent(3),
		"path":              reqPath, // 用于确定导航菜单
		"msg":               msg,
		"changeTypeOptions": changeTypeOptions,
	}, r)
}
func (c cWalletChangeLog) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cWalletChangeLog) Add(r *ghttp.Request) {
	var (
		d = entity.WalletChangeLog{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeLog) Update(r *ghttp.Request) {
	var (
		d     = entity.WalletChangeLog{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	delete(m, "type")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeLog) Clear(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	res.OkSession("ok", r)
	if err := service.Wallet.DelChangeLogs(ctx); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo("/admin/walletChangeLog", r)
}

type cWalletStatisticsLog struct{ cBase }

func (c cWalletStatisticsLog) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, T2: "u_user t2 on t1.uid = t2.id ", SearchFields: "t1.*,t2.uname", OrderBy: "t1.id desc", Fields: []model.Field{
			{Name: "id", Type: 1},
			{Name: "t2.uname", Type: 2, QueryName: "uname"},
			{Name: "created_date", Type: 5, QueryName: "begin"},
			{Name: "created_date", Type: 6, QueryName: "end"},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// Get the fields that need to be counted
	th, err := service.Wallet.GetStatisticsLogFieldsNeedToBeCountedOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
		"th":   th,
	}, r)
}
func (c cWalletStatisticsLog) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletStatisticsLog) Clear(r *ghttp.Request) {
	var (
		ctx = r.Context()
	)
	res.OkSession("ok", r)
	if err := service.Wallet.DelStatisticsLogs(ctx); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprintf("/admin/walletStatisticsLog")
	res.RedirectTo(path, r)
}
func (c cWalletStatisticsLog) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/walletStatisticsLog", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index) // 主页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/clear", c.Clear) // 清空
		g.GET("/del/:id", c.Del) // 删除请求
	})
}

type cWalletReport struct{}

func (c cWalletReport) RegisterRouter(g *ghttp.RouterGroup) {
	g.GET("/walletReport", c.WalletReport)
}
func (c cWalletReport) WalletReport(r *ghttp.Request) {
	var (
		d struct {
			Begin string
			End   string
			Uname string
		}
		ctx  = r.Context()
		path = r.URL.Path
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	node, err := service.System.GetNodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	report, err := service.Wallet.GetStatisticsLogReport(ctx, d.Begin, d.End, d.Uname)
	if err != nil {
		res.Err(err, r)
	}
	array, err := service.Wallet.GetStatisticsLogFieldsNeedToBeCountedOptionsIntoArray(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl("/statistics/walletReport/index.html", g.Map{
		"report": report,
		"node":   node,
		"path":   path,
		"array":  array,
	}, r)
}

type cWalletChangeType struct{ cBase }

func (c cWalletChangeType) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{T1: c.Table, OrderBy: "t1.id desc", Fields: []model.Field{
			{Name: "id", Type: 1},
			{Name: "type", Type: 1},
			{Name: "status", Type: 1},
		}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node": node,
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"path": reqPath, // 用于确定导航菜单
		"msg":  msg,
	}, r)
}
func (c cWalletChangeType) IndexAdd(r *ghttp.Request) {
	res.Tpl(fmt.Sprint(c.FileDir, "/add.html"), g.Map{"msg": service.System.GetMsgFromSession(r)}, r)
}
func (c cWalletChangeType) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = r.Get("id").Uint64()
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cWalletChangeType) Add(r *ghttp.Request) {
	var (
		d = entity.WalletChangeType{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeType) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeType) Update(r *ghttp.Request) {
	var (
		d     = entity.WalletChangeType{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := service.System.Update(r.Context(), table, d.Id, m, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletChangeType) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/walletChangeType", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)             // 主页面
		g.GET("/add", c.IndexAdd)       // 添加页面
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
	})
}

type cWalletTopUpApplication struct{ cBase }

func (c cWalletTopUpApplication) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/walletTopUpApplication", func(g *ghttp.RouterGroup) {
		g.Middleware(service.Admin.MiddlewareAuth)
		g.GET("/", c.Index)             // 主页面
		g.GET("/edit/:id", c.IndexEdit) // 修改页面
		g.Middleware(service.Admin.MiddlewareLock, service.Admin.MiddlewareActionLog)
		g.GET("/del/:id", c.Del)    // 删除请求
		g.POST("/add", c.Add)       // 添加请求
		g.POST("/update", c.Update) // 修改请求
		g.GET("/review", c.ReviewByAdmin)
	})
}
func (c cWalletTopUpApplication) Index(r *ghttp.Request) {
	var (
		ctx     = r.Context()
		reqPath = r.URL.Path
		file    = fmt.Sprintf("%s/index.html", c.FileDir)
		msg     = service.System.GetMsgFromSession(r)
		s       = model.Search{
			T1:           c.Table,
			T2:           "u_user t2 on t1.uid = t2.id",
			T3:           "u_wallet_change_type t3 on t1.change_type = t3.id",
			SearchFields: "t1.*,t2.uname uname,t3.title changeType,t3.class class", OrderBy: "t1.id desc", Fields: []model.Field{
				{Name: "id", Type: 1},
				{Name: "t1.ip", QueryName: "ip", Type: 2},
				{Name: "t1.trans_id", QueryName: "trans_id", Type: 2},
				{Name: "t1.status", QueryName: "status", Type: 1},
				{Name: "t1.description", QueryName: "description", Type: 2},
				{Name: "t1.change_type", QueryName: "change_type", Type: 1},
				{Name: "t2.uname", QueryName: "uname", Type: 2},
			}}
	)
	node, err := service.System.GetNodeInfo(ctx, reqPath)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := service.System.List(ctx, s, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	topUpOptions, err := service.Wallet.GetChangeTypeTopUpOptions(ctx)
	if err != nil {
		res.Err(err, r)
	}
	// 返回页面
	res.Tpl(file, g.Map{
		"node":         node,
		"list":         data,
		"page":         r.GetPage(total, s.Size).GetContent(3),
		"path":         reqPath, // 用于确定导航菜单
		"msg":          msg,
		"topUpOptions": topUpOptions,
	}, r)
}
func (c cWalletTopUpApplication) IndexEdit(r *ghttp.Request) {
	var (
		table = c.Table
		id    = xparam.ID(r)
		d     = g.Map{"msg": service.System.GetMsgFromSession(r)}
		f     = fmt.Sprint(c.FileDir, "/edit.html")
	)
	data, err := service.System.GetById(r.Context(), table, id, c.DBGroup)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data.Map() {
		r.SetForm(k, v)
	}
	res.Tpl(f, d, r)
}
func (c cWalletTopUpApplication) Add(r *ghttp.Request) {
	var (
		d = entity.WalletTopUpApplication{}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := service.System.Add(r.Context(), c.Table, &d, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/add?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletTopUpApplication) Del(r *ghttp.Request) {
	var (
		id    = r.Get("id")
		table = c.Table
	)
	res.OkSession("删除成功", r)
	if err := service.System.Del(r.Context(), table, id, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletTopUpApplication) Update(r *ghttp.Request) {
	var (
		d     = entity.WalletTopUpApplication{}
		table = c.Table
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("修改成功", r)
	adminId, err := service.System.GetAdminId(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	if err := service.System.Update(r.Context(), table, d.Id, g.Map{
		"description": d.Description,
		"aid":         adminId,
	}, c.DBGroup); err != nil {
		res.ErrSession(err, r)
	}
	path := fmt.Sprint(c.ReqPath, "/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap()))
	res.RedirectTo(path, r)
}
func (c cWalletTopUpApplication) ReviewByAdmin(r *ghttp.Request) {
	var (
		ctx           = r.Context()
		orderId       = r.Get("orderId").Uint64()
		operationType = r.Get("operationType").Int64()
		path          = fmt.Sprint(c.ReqPath, "?", xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("操作成功", r)
	session, err := service.Admin.GetInfoFromSession(r.Session)
	if err != nil {
		res.Err(err, r)
	}
	if err = service.Wallet.UpdateTopUpApplication(ctx, orderId, operationType, session.Admin.Id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
