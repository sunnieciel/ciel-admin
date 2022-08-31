package controller

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/file"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xfile"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
)

type cFile struct{ cBase }

var File = &cFile{
	cBase{Table: "s_file", ReqPath: "/admin/file", FileDir: "/sys/file"}}

func (c cFile) Index(r *ghttp.Request) {
	var (
		s = bo.Search{
			T1: "s_file", OrderBy: "t1.id desc", SearchFields: "t1.*", Fields: []bo.Field{
				{Name: "url", Type: 2}, {Name: "group", Type: 1}, {Name: "status", Type: 1},
			},
		}
		ctx  = r.Context()
		file = fmt.Sprintf("%s/index.html", c.FileDir)
		path = r.URL.Path
	)
	node, err := sys.NodeInfo(ctx, path)
	if err != nil {
		res.Err(err, r)
	}
	s.Page, s.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), s)
	if err != nil {
		res.Err(err, r)
	}
	res.Tpl(file, g.Map{
		"list": data,
		"page": r.GetPage(total, s.Size).GetContent(3),
		"node": node,
		"msg":  sys.MsgFromSession(r),
		"path": path,
	}, r)
}
func (c cFile) AddIndex(r *ghttp.Request) {
	var (
		file = fmt.Sprintf("%s/add.html", c.FileDir)
	)
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cFile) EditIndex(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		table = c.Table
		id    = xparam.ID(r)
		file  = fmt.Sprintf("%s/edit.html", c.FileDir)
	)
	data, err := sys.GetById(ctx, table, id)
	if err != nil {
		res.Err(err, r)
	}
	for k, v := range data {
		r.SetForm(k, v)
	}
	res.Tpl(file, g.Map{"msg": sys.MsgFromSession(r)}, r)
}
func (c cFile) Post(r *ghttp.Request) {
	var (
		d     = entity.File{}
		ctx   = r.Context()
		table = c.Table
		path  = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	res.OkSession("添加成功", r)
	if err := sys.Add(ctx, table, &d); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cFile) Del(r *ghttp.Request) {
	var (
		ctx   = r.Context()
		p     = fmt.Sprintf("%s?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
		table = c.Table
		id    = xparam.ID(r)
	)
	f, err := file.GetById(ctx, id)
	if err != nil {
		res.Err(err, r)
	}
	path, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		res.Err(err, r)
	}
	filePath := gfile.Pwd() + path.String() + "/" + f.Url
	if err = xfile.Remove(ctx, filePath); err != nil {
		res.Err(err, r)
	}
	res.OkSession("删除成功", r)
	if err = sys.Del(ctx, table, id); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(p, r)
}
func (c cFile) Put(r *ghttp.Request) {
	var (
		d     = entity.File{}
		ctx   = r.Context()
		table = c.Table
		id    = r.GetForm("id").Uint()
		path  = fmt.Sprintf("%s/edit/%v?%s", c.ReqPath, id, xurl.ToUrlParams(r.GetQueryMap()))
	)
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	res.OkSession("修改成功", r)
	if err := sys.Update(ctx, table, id, m); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
func (c cFile) Upload(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = fmt.Sprintf("%s/add?%s", c.ReqPath, xurl.ToUrlParams(r.GetQueryMap()))
	)
	res.OkSession("上传成功", r)
	if err := file.Upload(ctx, r); err != nil {
		res.ErrSession(err, r)
	}
	res.RedirectTo(path, r)
}
