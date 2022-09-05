package xuser

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	UidKey = "userInfoKey"
)

func Uid(r *ghttp.Request) uint64 {
	return r.Get(UidKey).Uint64()
}
