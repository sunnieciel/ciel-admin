package controller

import (
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---home-------------------------------------------------------------------
type home struct{}

var Home = &home{}

func (c *home) IndexPage(r *ghttp.Request) {
	res.Page(r, "/index.html", g.Map{"icon": "/resource/image/v2ex.png"})
}
