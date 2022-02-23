package controller

import (
	"ciel-begin/internal/consts"
	"ciel-begin/internal/model/entity"
	"ciel-begin/internal/service"
	"ciel-begin/manifest/config"
	"ciel-begin/utility/utils/res"
	"ciel-begin/utility/utils/xparam"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
)

type file struct {
	*config.SearchConf
}

func File() *file {
	return &file{SearchConf: &config.SearchConf{
		PageTitle: "File", UrlPrefix: "/file", T1: "s_file",
		Fields: []*config.Field{
			{Field: "id", EditHidden: true},
			{Field: "img", Type: "showImg", ShowImg: &config.ShowImg{ImgPrefix: consts.ImgPrefix, Field: "url"}},
			{Field: "group", Type: "select", Items: []*config.Item{{Value: 1, Text: "icon"}, {Value: 2, Text: "img"}, {Value: 3, Text: "video"}, {Value: 4, Text: "audio"}, {Value: 5, Text: "file"}}},
			{Field: "status", EditHidden: true, Type: "select", Search: true, Items: []*config.Item{{Value: "1", Text: "NO"}, {Value: "2", Text: "OFF"}}, Required: true},
			{Field: "url", Type: "file"},
			{Field: "created_at", EditHidden: true},
			{Field: "updated_at", EditHidden: true},
		},
	}}
}
func (c *file) List(r *ghttp.Request) {
	page, size := res.GetPage(r)
	c.Page = page
	c.Size = size
	total, data, err := service.System().List(r.Context(), c.SearchConf)
	if err != nil {
		res.Err(err, r)
	}
	res.PageList(r, consts.DefaultPage, total, data, c)
}
func (c *file) GetById(r *ghttp.Request) {
	data, err := service.System().GetById(r.Context(), c.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	res.OkData(data, r)
}
func (c *file) Post(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Add(r.Context(), c.T1, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *file) Put(r *ghttp.Request) {
	d := entity.File{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	if err := service.System().Update(r.Context(), c.T1, d.Id, &d); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *file) Del(r *ghttp.Request) {
	f, err := service.File().GetById(r.Context(), xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	path, err := g.Cfg().Get(r.Context(), "server.rootFilePath")
	if err != nil {
		res.Err(err, r)
	}
	p := gfile.Pwd() + path.String() + "/" + f.Url
	if gfile.Exists(p) {
		_ = gfile.Remove(p)
	}
	if err := service.System().Del(r.Context(), c.T1, xparam.ID(r)); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
func (c *file) Upload(r *ghttp.Request) {
	if err := service.File().Upload(r.Context(), r); err != nil {
		res.Err(err, r)
	}
	res.Ok(r)
}
