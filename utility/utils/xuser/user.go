package xuser

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func Uid(r *ghttp.Request) uint64 {
	return r.Get("userInfoKey").Uint64()
}
