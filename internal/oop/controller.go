package oop

import "github.com/gogf/gf/v2/net/ghttp"

type IControler interface {
	List(*ghttp.Request)
	GetById(*ghttp.Request)
	Post(*ghttp.Request)
	Put(*ghttp.Request)
	Del(*ghttp.Request)
	Path(*ghttp.Request)
}
