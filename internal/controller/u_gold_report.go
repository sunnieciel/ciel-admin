package controller

import (
	"ciel-admin/internal/service/goldchangelog"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	GoldReport = cGoldReport{}
)

type cGoldReport struct {
}

func (c cGoldReport) RegisterRouter(g *ghttp.RouterGroup) {
	g.GET("/goldReport", c.GoldReport)
}

func (c cGoldReport) GoldReport(r *ghttp.Request) {
	var (
		d struct {
			Begin string
			End   string
		}
		ctx  = r.Context()
		path = r.URL.Path
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	report, err := goldchangelog.GoldReport(ctx, d.Begin, d.End)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl("/statistics/goldReport/index.html", g.Map{
		"report": report,
		"node":   node,
		"path":   path,
	}, r)
}
