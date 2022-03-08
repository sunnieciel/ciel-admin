package controller

import (
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/net/ghttp"
)

// ---home-------------------------------------------------------------------
type home struct{}

func Home() *home                          { return &home{} }
func (c *home) IndexPage(r *ghttp.Request) { res.Page(r, "/index.html") }
