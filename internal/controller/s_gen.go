package controller

import (
	"ciel-admin/internal/service/admin"
	"ciel-admin/internal/service/gen"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Gen = cGen{}
)

type cGen struct{}

func (c cGen) Index(r *ghttp.Request) {
	var (
		path = r.URL.Path
		ctx  = r.Context()
	)
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	tables, err := gen.Tables(ctx)
	if err != nil {
		res.ErrSession(err, r)
	}

	menuLeve1, err := gen.MenuLeve1(ctx)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl("/sys/gen/index.html", g.Map{
		"node":       node,
		"path":       path,
		"msg":        sys.MsgFromSession(r),
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
		}
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := gen.Gen(ctx, d.Table, d.Group, d.Menu, d.Prefix, d.ApiGroup, d.HtmlGroup); err != nil {
		res.Err(err, r)
	}
	res.OkMsg("生成成功", r)
}

func (c cGen) RegisterRouter(g *ghttp.RouterGroup) {
	g.Group("/gen", func(g *ghttp.RouterGroup) {
		g.Middleware(admin.AuthMiddleware)
		g.GET("/", c.Index)
		g.POST("/table", c.Gen)
	})
}
