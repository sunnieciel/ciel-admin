package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service/admin"
	"ciel-admin/internal/service/dict"
	"ciel-admin/internal/service/role"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	cRoleApi struct{ cBase }
)

var (
	RoleApi = &cRoleApi{cBase{Table: "s_role_api", ReqPath: "/admin/roleApi", FileDir: "/sys/roleApi"}}
)

func (c cRoleApi) Index(r *ghttp.Request) {
	var (
		ctx = r.Context()
		s   = bo.Search{
			T1: "s_role_api", T2: "s_role t2 on t1.rid = t2.id", T3: "s_api t3 on t1.aid = t3.id",
			OrderBy:      "t3.group,t3.type",
			SearchFields: "t1.*,t2.name r_name,t3.url url,t3.group,t3.method,t3.desc,t3.type", Fields: []bo.Field{
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
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	roleInfo, err := role.GetById(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	node.Name = "角色禁用API"
	node.Desc = fmt.Sprintf("如果你不希望<span class='color-red strong'>【%s】</span>角色访问某些api功能，可以点击添加按钮，将他们添加到下面的列表中", roleInfo.Name)
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), s)
	if err != nil {
		res.Err(err, r)
	}
	groups, err := dict.ApiGroup(ctx)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl(file, g.Map{
		"list":   data,
		"page":   r.GetPage(total, s.Size).GetContent(3),
		"node":   node,
		"msg":    sys.MsgFromSession(r),
		"path":   path,
		"groups": groups,
	}); err != nil {
		res.Err(err, r)
	}
}

func (c cRoleApi) AddIndex(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		rid  = r.Get("rid")
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	apis, err := role.NoApi(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	roleData, err := role.GetById(ctx, rid)
	if err != nil {
		res.Err(err, r)
	}
	d := g.Map{
		"msg":  sys.MsgFromSession(r),
		"apis": apis,
		"role": roleData,
	}
	_ = r.Response.WriteTpl(file, d)
}

func (c cRoleApi) Post(r *ghttp.Request) {
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
	if err := role.AddRoleApi(ctx, d.Rid, d.Aid); err != nil {
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
	if err := sys.Del(ctx, c.Table, id); err != nil {
		res.ErrSession(err, r)
	}
	g.Log().Infof(ctx, path)
	res.RedirectTo(path, r)
}

func (c cRoleApi) Clear(r *ghttp.Request) {
	var (
		ctx = r.Context()
		rid = r.Get("rid")
		t   = r.Get("type").Int()
	)
	err := role.ClearApi(ctx, rid, t)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

func (c cRoleApi) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/roleApi", func(g *ghttp.RouterGroup) {
		g.Middleware(admin.AuthMiddleware)
		g.GET("/", c.Index)
		g.GET("/add", c.AddIndex)
		g.Middleware(admin.LockMiddleware, admin.ActionMiddleware)
		g.GET("/del/:id", c.Del)
		g.GET("/clear/:rid", c.Clear)
		g.POST("/post", c.Post)
	})
}
