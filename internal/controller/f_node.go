// =================================================================================

package controller

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/internal/service/sys"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xparam"
	"ciel-admin/utility/utils/xurl"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"time"
)

type cNode struct{ bo.Search }

var Node = &cNode{Search: bo.Search{
	T1: "f_node", T2: "s_admin t2 on t1.uid= t2.id", OrderBy: "t1.year desc,t1.month desc,t1.day desc", SearchFields: "t1.*,t2.uname",
	Fields: []bo.Field{
		{Name: "t1.year", SearchType: 1, QueryName: "node_year"},
		{Name: "t1.month", SearchType: 1, QueryName: "node_month"},
		{Name: "t1.day", SearchType: 1, QueryName: "node_day"},
		{Name: "t1.category", SearchType: 1, QueryName: "node_category"},
		{Name: "t2.uname", SearchType: 2, QueryName: "node_uname"}, {Name: "level", SearchType: 1, QueryName: "node_level"}, {Name: "tag", SearchType: 2, QueryName: "node_tag"}, {Name: "main_things", SearchType: 2, QueryName: "node_main_things"},
	},
}}

func (c cNode) Path(r *ghttp.Request) {
	node, err := sys.NodeInfo(r.Context(), r.URL.Path)
	if err != nil {
		res.Err(err, r)
	}
	c.Page, c.Size = res.GetPage(r)
	total, data, err := sys.List(r.Context(), c.Search)
	if err != nil {
		res.Err(err, r)
	}
	if err = r.Response.WriteTpl("/f/node/index.html", g.Map{
		"list":     data,
		"page":     r.GetPage(total, c.Size).GetContent(3),
		"node":     node,
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
	}); err != nil {
		res.Err(err, r)
	}
}

func (c cNode) PathAdd(r *ghttp.Request) {
	r.Response.WriteTpl("/f/node/add.html", g.Map{
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
	})
}
func (c cNode) Post(r *ghttp.Request) {
	data := entity.Node{}
	if err := r.Parse(&data); err != nil {
		res.Err(err, r)
	}
	y := fmt.Sprint(gtime.Now().Year())
	m := gtime.Now().Month()
	d := time.Now().Day()
	happenDate := r.Get("happen_date").String()
	if happenDate != "" {
		split := strings.Split(happenDate, ",")
		switch len(split) {
		case 1:
			y = gconv.String(split[0])
		case 2:
			y = gconv.String(split[0])
			m = gconv.Int(split[1])
		case 3:
			y = gconv.String(split[0])
			m = gconv.Int(split[1])
			d = gconv.Int(split[2])
		}
	}
	data.Year = y
	data.Month = m
	data.Day = d
	admin, _ := sys.GetAdmin(r)
	data.Uid = admin.Admin.Id
	msg := fmt.Sprintf(consts.MsgPrimary, "添加成功")
	if err := sys.Add(r.Context(), c.T1, &data); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/node/path/add?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cNode) Del(r *ghttp.Request) {
	id := r.Get("id")
	msg := fmt.Sprintf(consts.MsgPrimary, "删除成功")
	if err := sys.Del(r.Context(), c.T1, id); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/node/path?", xurl.ToUrlParams(r.GetQueryMap())))
}
func (c cNode) PathEdit(r *ghttp.Request) {
	data, err := sys.GetById(r.Context(), c.Search.T1, xparam.ID(r))
	if err != nil {
		res.Err(err, r)
	}
	_ = r.Session.Set("node_edit", data.Map())
	_ = r.Response.WriteTpl("/f/node/edit.html", g.Map{
		"msg":      sys.MsgFromSession(r),
		"category": getCategory(r),
	})
}
func (c cNode) Put(r *ghttp.Request) {
	d := entity.Node{}
	if err := r.Parse(&d); err != nil {
		res.Err(err, r)
	}
	m := gconv.Map(d)
	delete(m, "createdAt")
	msg := fmt.Sprintf(consts.MsgPrimary, "修改成功")
	if err := sys.Update(r.Context(), c.T1, d.Id, m); err != nil {
		msg = fmt.Sprintf(consts.MsgWarning, err.Error())
	}
	_ = r.Session.Set("msg", msg)
	r.Response.RedirectTo(fmt.Sprint("/node/path/edit/", d.Id, "?", xurl.ToUrlParams(r.GetQueryMap())))
}

func getCategory(r *ghttp.Request) []g.Map {
	key, err := sys.DictGetByKey(r.Context(), "node-category")
	if err != nil {
		res.Err(err, r)
	}
	category := make([]g.Map, 0)
	for _, i := range strings.Split(key, ",") {
		temp := strings.Split(i, "_")
		category = append(category, map[string]interface{}{
			"value": temp[0],
			"label": temp[1],
		})
	}
	return category
}
