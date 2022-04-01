package controller

import (
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sys struct{}

func (s sys) Path(r *ghttp.Request) {
	path := r.GetQuery("path")
	res.Page(r, path.String())
}

func (s sys) PathGithub(r *ghttp.Request) {
	res.Page(r, "/sys/rss/github.html", g.Map{"icon": "/resource/image/github.png"})
}

func (s sys) OsChina(r *ghttp.Request) {
	res.Page(r, "/sys/rss/oschina.html", g.Map{"icon": "/resource/image/github.png"})
}

func (s sys) Douban(r *ghttp.Request) {
	res.Page(r, "/sys/rss/douban.html", g.Map{"icon": "/resource/image/github.png"})
}

var Sys = &sys{}
