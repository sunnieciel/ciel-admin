package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/service"
	"ciel-admin/utility/utils/res"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

type gen struct {
}

func (c gen) Path(r *ghttp.Request) {
	icon, err := service.System().Icon(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	res.Page(r, "/sys/gen.html", g.Map{"icon": icon})
}

func (c gen) Tables(r *ghttp.Request) {
	data, err := service.Gen().Tables(r.Context())
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}

func (c gen) Fields(r *ghttp.Request) {
	var d struct {
		Table string `v:"required#名表不能为空"`
	}
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	data, err := service.Gen().Fields(r.Context(), d.Table)
	res.OkData(data, r)
}

func (c gen) GenCode(r *ghttp.Request) {
	var d bo.GenCodeInfo
	err := r.Parse(&d)
	if err != nil {
		res.Err(err, r)
	}
	form := r.GetForm("fields")
	d.Fields = make([]*bo.Field, 0)
	glog.Info(r.Context(), form)
	for _, v := range form.Map() {
		stemp := v.(map[string]interface{})
		field := bo.Field{
			Name:        gconv.String(stemp["Name"]),
			Comment:     gconv.String(stemp["Comment"]),
			Type:        gconv.String(stemp["Type"]),
			SearchType:  gconv.String(stemp["SearchType"]),
			QueryField:  gconv.String(stemp["QueryField"]),
			Sort:        gconv.Int(stemp["sort"]),
			DetailsType: gconv.String(stemp["DetailsType"]),
		}
		d.Fields = append(d.Fields, &field)
	}

	err = service.Gen().GenCode(r.Context(), &d)
	if err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}

var Gen = &gen{}
