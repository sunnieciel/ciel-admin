package controller

import (
	"ciel-admin/internal/service"
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/net/ghttp"
)

type rss struct {
}

func (c rss) Fetch(r *ghttp.Request) {
	//timeout, cancelFunc := context.WithTimeout(r.Context(), time.Second*15)
	//defer cancelFunc()
	data, err := service.Rss().Feftch(r.Context(), r.GetQuery("url").String())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

var Rss = &rss{}
