package res

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Err(err error, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code": -1,
		"msg":  err.Error(),
	})
}
func Ok(r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "ok",
	})
}

func OkMsg(msg string, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  msg,
	})
}
func OkData(data interface{}, r *ghttp.Request) {
	r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}
