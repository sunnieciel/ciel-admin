package xparam

import "github.com/gogf/gf/v2/net/ghttp"

func ID(r *ghttp.Request) interface{} {
	query := r.GetQuery("id")
	if query.IsEmpty() {
		return r.Get("id")
	}
	return query
}

func TakeID(r *ghttp.Request) interface{} {
	id := r.GetQuery("id")
	if !id.IsEmpty() {
		return id.Int64()
	}
	id, err := r.Session.Get("id")
	if err != nil {
		return 0
	}
	if !id.IsEmpty() {
		_ = r.Session.Remove("id")
		return id.Int()
	}
	return 0
}
