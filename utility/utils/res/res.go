package res

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
)

func PageList(r *ghttp.Request, page string, total int, list interface{}, info interface{}) {
	if err := r.Response.WriteTpl(page, g.Map{
		"list":  list,
		"total": total,
		"page":  r.GetPage(total, r.GetQuery("size").Int()).GetContent(3),
		"c":     info,
	}); err != nil {
		glog.Error(r.Context(), err)
	}
	r.Exit()
}
func Page(r *ghttp.Request, page string, data ...interface{}) {
	d := g.Map{}
	if len(data) > 0 {
		d["data"] = data[0]
	}
	if err := r.Response.WriteTpl(page, d); err != nil {
		glog.Error(r.Context(), err)
	}
	r.Exit()
}
func GetPage(r *ghttp.Request) (page, size int) {
	page = r.GetQuery("page").Int()
	size = r.GetQuery("size").Int()
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	r.SetQuery("page", page)
	r.SetQuery("size", size)
	return
}
func Err(err error, r *ghttp.Request) {
	err = r.Response.WriteJsonExit(g.Map{
		"code": -1,
		"msg":  err.Error(),
	})
	if err != nil {
		glog.Error(r.Context(), err)
	}
}
func Ok(r *ghttp.Request) {
	err := r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "ok",
	})
	if err != nil {
		glog.Error(r.Context(), err)
		return
	}
}
func OkData(data interface{}, r *ghttp.Request) {
	_ = r.Response.WriteJsonExit(g.Map{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}
