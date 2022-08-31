package sys

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	Uid = "userInfoKey"
)

func MsgFromSession(r *ghttp.Request) string {
	msg, err := r.Session.Get("msg")
	if err != nil {
		return ""
	}
	if !msg.IsEmpty() {
		_ = r.Session.Remove("msg")
	}
	return msg.String()
}
