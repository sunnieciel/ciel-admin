package xparam

import "github.com/gogf/gf/v2/net/ghttp"

func ID(r *ghttp.Request) interface{} {
	query := r.GetQuery("id")
	if query.IsEmpty() {
		return r.Get("id")
	}
	return query
}
