package logic

import (
	v1 "ciel-admin/api/v1"
	"ciel-admin/internal/model"
	"ciel-admin/utility/utils/xhtml"
	"ciel-admin/utility/utils/xjwt"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/util/guid"

	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/do"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xcaptcha"
	"ciel-admin/utility/utils/xfile"
	"ciel-admin/utility/utils/xpwd"
	"ciel-admin/utility/utils/xredis"
	"ciel-admin/utility/utils/xstr"
	"ciel-admin/utility/utils/xtime"
	"context"
	"encoding/json"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"math"
	"net/http"
	"path"
	"regexp"
	"sort"
	"strings"
	"time"
)

var (
	Menu            = lMenu{}
	Role            = lRole{}
	System          = lSystem{}
	Admin           = lAdmin{}
	Dict            = lDict{}
	File            = lFile{}
	Gen             = lGen{}
	Ws              = lWs{}
	Session         = lSession{}
	AdminSessionKey = "adminInfo"
	User            = lUser{}
	Wallet          = lWallet{}
	users           = gmap.New(true)
	admins          = gmap.New(true)
)

type lMenu struct{}

func (l lMenu) GetById(ctx context.Context, id uint64) (*entity.Menu, error) {
	var data entity.Menu
	one, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lMenu) GetByPath(ctx context.Context, path string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).One("path", path)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lMenu) GetByName(ctx context.Context, name string) (*entity.Menu, error) {
	var data entity.Menu
	one, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).One("name", name)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, nil
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lMenu) ListByPid(ctx context.Context, id int) ([]*entity.Menu, error) {
	var data = make([]*entity.Menu, 0)
	err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Scan(&data, "pid", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lMenu) UpdateGroupSort(ctx context.Context, sort int, id uint64) error {
	change := func(in float64) float64 {
		arr := strings.Split(fmt.Sprintf("%.2f", in), ".")
		resStr := fmt.Sprintf("%d.%s", sort, arr[1])
		return gconv.Float64(resStr)
	}
	pMenu, err := l.GetById(ctx, id)
	if err != nil {
		return err
	}
	pMenu.Sort = change(pMenu.Sort)
	if _, err = g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Save(pMenu); err != nil {
		return err
	}
	arr, err := l.ListByPid(ctx, pMenu.Id)
	if err != nil {
		return err
	}
	for _, i := range arr {
		i.Sort = change(i.Sort)
		if _, err = g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Save(i); err != nil {
			return err
		}
	}
	return nil
}

type lRole struct{}

func (r lRole) AddMenu(ctx context.Context, rid int, ids []int) error {
	return g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for _, item := range ids {
			if _, err := tx.Ctx(ctx).Model(dao.RoleMenu.Table()).Replace(g.Map{
				"rid": rid,
				"mid": item,
			}); err != nil {
				return err
			}
		}
		return nil
	})
}
func (r lRole) AddApi(ctx context.Context, rid int, ids []int) error {
	for _, item := range ids {
		_, err := g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Replace(g.Map{
			"rid": rid,
			"aid": item,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (r lRole) GetById(ctx context.Context, id interface{}) (*entity.Role, error) {
	var data entity.Role
	one, err := g.DB("sys").Model(dao.Role.Table()).Ctx(ctx).One("id", id)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	err = one.Struct(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r lRole) GetRoleOptions(ctx context.Context) (string, error) {
	var (
		array = make([]string, 0)
	)
	all, err := g.DB("sys").Model(dao.Role.Table()).Ctx(ctx).All()
	if err != nil {
		return "", err
	}
	for index, m := range all {
		id := m["id"]
		name := m["name"]
		array = append(array, fmt.Sprintf(fmt.Sprintf("%v:%v:%s", id, name, xhtml.SwitchTagClass(index))))
	}
	return strings.Join(array, ","), nil
}
func (r lRole) ListRoleNoMenus(ctx context.Context, rid interface{}) (gdb.List, error) {
	array, err := g.DB("sys").Model(dao.RoleMenu.Table()).Ctx(ctx).Array("mid", "rid", rid)
	if err != nil {
		return nil, err
	}
	db := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx)
	if len(array) != 0 {
		db = db.WhereNotIn("id", array)
	}
	all, err := db.Order("sort").All()
	if err != nil {
		return nil, err
	}
	return all.List(), nil
}
func (r lRole) ListRoleNoApis(ctx context.Context, rid interface{}) (gdb.List, error) {
	array, err := g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Array("aid", "rid", rid)
	if err != nil {
		return nil, err
	}
	db := g.DB("sys").Model(dao.Api.Table()).Ctx(ctx)
	if len(array) != 0 {
		db = db.WhereNotIn("id", array)
	}
	all, err := db.Order("group").All()
	if err != nil {
		return nil, err
	}
	return all.List(), nil
}
func (r lRole) DelApis(ctx context.Context, rid interface{}, t int) error {
	// 做实际的清除工作
	if t == 0 {
		if _, err := g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Delete("rid", rid); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	all, err := g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).All("rid", rid)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if all.IsEmpty() {
		return nil
	}
	for _, i := range all {
		apiData, err := System.GetApiById(ctx, i["aid"].Int())
		if err != nil {
			return err
		}
		apiType := apiData.Type
		switch t {
		case 1: // 允许所有查询
			if apiType == 4 || apiType == 5 {
				if _, err = g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Delete("id", i["id"]); err != nil {
					return err
				}
			}
		case 2: // 允许所有添加操作
			if apiType == 1 {
				if _, err = g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Delete("id", i["id"]); err != nil {
					return err
				}
			}
		case 3: // 允许所有修改操作
			if apiType == 3 {
				if _, err = g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Delete("id", i["id"]); err != nil {
					return err
				}
			}
		case 4: // 允许所有删除操作
			if apiType == 2 {
				if _, err = g.DB("sys").Model(dao.RoleApi.Table()).Ctx(ctx).Delete("id", i["id"]); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
func (r lRole) DelMenus(ctx context.Context, rid interface{}) error {
	_, err := r.GetById(ctx, rid)
	if err != nil {
		return err
	}
	if _, err = g.DB("sys").Model(dao.RoleMenu.Table()).Ctx(ctx).Delete("rid", rid); err != nil {
		return err
	}
	return nil
}
func (r lRole) CheckRoleApi(ctx context.Context, rid int, uri string) bool {
	if strings.Contains(uri, "?") {
		uri = strings.Split(uri, "?")[0]
	}
	if uri == "/" {
		return true
	}
	s := fmt.Sprint(regexp.MustCompile(".+/del/").FindString(uri), ":id")
	if s != ":id" {
		uri = s
	}
	count, _ := g.DB("sys").Ctx(ctx).Model("s_role t1").
		LeftJoin("s_role_api t2 on t1.id = t2.rid").
		LeftJoin("s_api t3 on t2.aid = t3.id").
		Where("t3.url = ? and  t1.id = ?  ", uri, rid).
		Count()
	if count == 1 {
		return false
	}
	return true
}

type lDict struct{}

func (l lDict) GetByKeyString(ctx context.Context, key string) (string, error) {
	dict, err := l.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return dict.V, nil
}
func (l lDict) GetByKey(ctx context.Context, s string) (*entity.Dict, error) {
	var data entity.Dict
	one, err := g.DB("sys").Model(dao.Dict.Table()).Ctx(ctx).One("k", s)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lDict) UpdateWhiteIps(ctx context.Context, v ...string) error {
	if len(v) == 0 {
		d, err := l.GetByKey(ctx, "white_ips")
		if err != nil {
			return err
		}
		consts.WhiteIps = d.V
	} else {
		consts.WhiteIps = v[0]
	}
	return nil
}
func (l lDict) TakeApiGroupOptions(ctx context.Context) (string, error) {
	data, err := l.GetByKeyString(ctx, "api_group")
	if err != nil {
		return "", err
	}
	arr := make([]string, 0)
	for index, i := range gstr.Split(data, "\n") {
		if i != "" {
			i = gstr.TrimAll(i)
			arr = append(arr, fmt.Sprintf("%s:%s:%s", i, i, xhtml.SwitchTagClass(index)))
		}
	}
	return strings.Join(arr, ","), nil
}

type lSystem struct{}

func (l lSystem) Add(ctx context.Context, table interface{}, data interface{}, dbGroup ...string) error {
	group := ""
	if len(dbGroup) != 0 {
		group = dbGroup[0]
	}
	if _, err := g.DB(group).Ctx(ctx).Model(table).Insert(data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l lSystem) GetApiById(ctx context.Context, id interface{}) (*entity.Api, error) {
	var data entity.Api
	one, err := g.DB("sys").Model(dao.Api.Table()).Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lSystem) GetNodeInfo(ctx context.Context, path string) (*entity.Menu, error) {
	m, err := Menu.GetByPath(ctx, path)
	if err != nil {
		m = &entity.Menu{}
	}
	if m.Icon == "" {
		m.Icon = gstr.Replace(consts.ImgPrefix, "/upload", "") + "resource/image/golang.png"
	} else {
		if !strings.HasPrefix(m.Icon, "http") {
			m.Icon = consts.ImgPrefix + m.Icon
		}
	}
	if m.BgImg != "" {
		if !strings.HasPrefix(m.BgImg, "http") {
			m.BgImg = consts.ImgPrefix + m.BgImg
		}
	}
	if m.Desc == "" {
		m.Desc = "暂无相关说明"
	}
	return m, nil
}
func (l lSystem) GetMsgFromSession(r *ghttp.Request) string {
	msg, err := r.Session.Get("msg")
	if err != nil {
		return ""
	}
	if !msg.IsEmpty() {
		if err = r.Session.Remove("msg"); err != nil {
			res.Err(err, r)
		}
	}
	return msg.String()
}
func (l lSystem) GetById(ctx context.Context, table interface{}, id interface{}, dbGroup ...string) (gdb.Record, error) {
	group := ""
	if len(dbGroup) != 0 {
		group = dbGroup[0]
	}
	one, err := g.DB(group).Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return one, nil
}
func (l lSystem) List(ctx context.Context, c model.Search, dbGroup ...string) (int, gdb.List, error) {
	group := ""
	if len(dbGroup) != 0 {
		group = dbGroup[0]
	}
	db := g.DB(group).Ctx(ctx).Model(c.T1 + " t1")
	if c.T2 != "" {
		db = db.LeftJoin(c.T2)
	}
	if c.T3 != "" {
		db = db.LeftJoin(c.T3)
	}
	if c.T4 != "" {
		db = db.LeftJoin(c.T4)
	}
	if c.T5 != "" {
		db = db.LeftJoin(c.T5)
	}
	if c.T6 != "" {
		db = db.LeftJoin(c.T6)
	}
	conditions := c.FilterConditions(ctx)
	if len(conditions) > 0 {
		for _, item := range conditions {
			field := item.Name
			if g.IsEmpty(item.Value) {
				continue
			}
			if !strings.Contains(field, ".") {
				field = "t1." + field
			}
			switch item.Type {
			case 1:
				db = db.Where(field, item.Value)
			case 2: // like
				db = db.WhereLike(field, xstr.Like(gconv.String(item.Value)))
			case 3: // >
				db = db.WhereGT(field, item.Value)
			case 4: // <
				db = db.WhereLT(field, item.Value)
			case 5: // >=
				db = db.WhereGTE(field, item.Value)
			case 6: // <=
				db = db.WhereLTE(field, item.Value)
			case 7: // !=
				db = db.WhereNot(field, item.Value)
			case 8: //date
				if c.Begin != "" {
					db = db.Where(field, ">=", c.Begin)
				}
				if c.End != "" {
					db = db.Where(field, "<=", c.End)
				}
			case 9: // date start
				if c.Begin != "" {
					db = db.Where(field, ">=", xtime.BeginOfDateStr(c.Begin))
				}
				if c.End != "" {
					db = db.Where(field, "<=", xtime.EndOfDateStr(c.End))
				}
			default:
				continue
			}
		}
	}
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	var o = "t1.id desc"
	if c.OrderBy != "" {
		o = c.OrderBy
	}
	if c.SearchFields == "" {
		c.SearchFields = "t1.*"
	}
	all, err := db.Page(c.Page, c.Size).Fields(c.SearchFields).Order(o).All()
	if err != nil {
		return 0, nil, err
	}
	if all.IsEmpty() {
		return 0, nil, err
	}
	data := all.List()
	return int(count), data, nil
}
func (l lSystem) ListAllDict(ctx context.Context) (g.Map, error) {
	all, err := g.DB("sys").Ctx(ctx).Model(dao.Dict.Table()).All("`group`=2 and status=1")
	if err != nil {
		return nil, err
	}
	data := g.Map{}
	for _, i := range all {
		data[i["k"].String()] = g.Map{
			"value": i["v"].String(),
			"title": i["title"].String(),
			"desc":  i["desc"].String(),
		}
	}
	return data, nil
}
func (l lSystem) ListBanners(ctx context.Context) ([]*v1.BannerRes, error) {
	var data = make([]*v1.BannerRes, 0)
	if err := g.DB("sys").Model(dao.Banner.Table()).Ctx(ctx).Scan(&data, "status =1"); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	for _, i := range data {
		if !strings.HasPrefix(i.Image, "http") {
			i.Image = consts.ImgPrefix + i.Image
		}
	}
	return data, nil
}
func (l lSystem) Del(ctx context.Context, table interface{}, id interface{}, dbGroup ...string) error {
	group := ""
	if len(dbGroup) != 0 {
		group = dbGroup[0]
	}
	if _, err := g.DB(group).Ctx(ctx).Model(table).Delete("id", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lSystem) DelFun(ctx context.Context, dbGroup string, fun func(ctx context.Context, db gdb.DB) error) error {
	return fun(ctx, g.DB(dbGroup).Ctx(ctx))
}
func (l lSystem) Update(ctx context.Context, table, id, data interface{}, dbGroup ...string) error {
	group := ""
	if len(dbGroup) != 0 {
		group = dbGroup[0]
	}
	_, err := g.DB(group).Model(fmt.Sprint(table)).Where("id", id).Update(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lSystem) MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
func (l lSystem) MiddlewareWhiteIp(r *ghttp.Request) {
	if consts.WhiteIps != "" {
		if !gstr.Contains(consts.WhiteIps, r.GetClientIp()) {
			r.Response.WriteStatus(http.StatusForbidden, fmt.Sprintf("%l ip error", r.GetClientIp()))
			r.Exit()
		}
	}
	r.Middleware.Next()
}

type lAdmin struct{}

func (l lAdmin) Add(ctx context.Context, in entity.Admin) error {
	if in.Pwd == "" {
		return consts.ErrPassEmpty
	}
	if in.Nickname == "" {
		in.Nickname = in.Uname
	}
	if in.Email != "" {
		if err := g.Validator().Rules("email").Data(in.Email).Run(ctx); err != nil {
			return consts.ErrFormatEmail
		}
	}
	count, err := g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Count("uname", in.Uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	in.Pwd = xpwd.GenPwd(in.Pwd)
	if _, err = g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Insert(in); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lAdmin) AddMessage(ctx context.Context, uname string, title string, url string, t int) error {
	admin, err := l.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	return g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// add msg
		if _, err = tx.Model(dao.AdminMessage.Table()).Insert(entity.AdminMessage{
			Aid:   uint64(admin.Id),
			Type:  uint(t),
			Title: title,
			Url:   url,
		}); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		// add unread msg num
		if err = l.AddUnReadMsgNumTx(tx, ctx, admin.Id, 1); err != nil {
			return err
		}
		// notice user
		if err = Ws.NoticeAdmin(ctx, "", admin.Id); err != nil {
			return err
		}
		return nil
	})
}
func (l lAdmin) AddUnReadMsgNumTx(tx *gdb.TX, ctx context.Context, id int, i int) error {
	if _, err := tx.Model(dao.Admin.Table()).Where("id", id).Increment("unread_msg_num", i); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lAdmin) AddUnReadMsgNum(ctx context.Context, id int, i int) error {
	if _, err := g.DB("sys").Model(dao.Admin.Table()).Where("id", id).Increment("unread_msg_num", i); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lAdmin) AddUnreadMsg(ctx context.Context, message entity.AdminMessage) error {
	if _, err := g.DB("sys").Ctx(ctx).Model(dao.AdminMessage.Table()).Insert(message); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l lAdmin) GetByUname(ctx context.Context, uname string) (*entity.Admin, error) {
	var data entity.Admin
	one, err := g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).One("uname", uname)
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrUserNotFound
	}
	err = one.Struct(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lAdmin) ListNotifications(ctx context.Context, page int, size int, aid int) (int64, []*entity.AdminMessage, error) {
	var data []*entity.AdminMessage
	db := g.DB("sys").Model(dao.AdminMessage.Table()).Ctx(ctx).Where("aid", aid)
	count, err := db.Count()
	if err != nil {
		return 0, nil, err
	}
	if err = db.Page(page, size).OrderDesc("id").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return 0, nil, err
	}
	return count, data, nil
}
func (l lAdmin) ListMenus(ctx context.Context, rid int, pid int) ([]*model.Menu, error) {
	var d = make([]*model.Menu, 0)
	menus, err := l.doMenus(ctx, rid, pid)
	if err != nil {
		return nil, err
	}
	d = append(d, menus...)
	return d, err
}
func (l lAdmin) doMenus(ctx context.Context, rid, pid int) ([]*model.Menu, error) {
	var data []*model.Menu
	err := g.DB("sys").Ctx(ctx).Model(dao.RoleMenu.Table()+" t1").
		LeftJoin(dao.Menu.Table()+" t2 on t1.mid = t2.id").
		Fields("t2.*").
		Where("t1.rid = ? and t2.pid = ?", rid, pid).
		Order("t2.sort").
		Scan(&data)
	if err != nil {
		return nil, err
	}
	for _, item := range data {
		if item.Type == 2 {
			children, err := l.doMenus(ctx, rid, item.Id)
			if err != nil {
				return nil, err
			}
			item.Children = children
		}
	}
	return data, nil
}

func (l lAdmin) DelNotifications(ctx context.Context, aid int) error {
	// remove unread cache
	if _, err := gcache.Remove(ctx, fmt.Sprint(consts.AdminUnreadKey, aid)); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Update(g.Map{"unread_msg_num": 0}, "id", aid); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	ghttp.RequestFromCtx(ctx).Cookie.Set("unreadNum", "0")
	return nil
}
func (l lAdmin) DelOperationLogs(ctx context.Context) error {
	if _, err := g.DB("sys").Model(dao.OperationLog.Table()).Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}
	return nil
}
func (l lAdmin) DelLoginLogs(ctx context.Context) error {
	if _, err := g.DB("sys").Model(dao.AdminLoginLog.Table()).Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}
	return nil
}

func (l lAdmin) UpdatePwd(ctx context.Context, pwd string, pwd2 string) error {
	adminBo, err := Session.GetAdmin(ghttp.RequestFromCtx(ctx).Session)
	if err != nil {
		return err
	}
	u, err := l.GetByUname(ctx, adminBo.Admin.Uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(u.Pwd, pwd) {
		return errors.New("old password not match")
	}
	u.Pwd = xpwd.GenPwd(pwd2)
	err = Session.DelAdmin(ctx)
	if err != nil {
		return err
	}
	if _, err = g.DB("sys").Model(dao.Admin.Table()).Update(u, "id", u.Id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lAdmin) UpdateUname(ctx context.Context, id interface{}, uname interface{}) error {
	count, err := g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Count("uname", uname)
	if err != nil {
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if _, err = g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Update(g.Map{"uname": uname}, "id", id); err != nil {
		return err
	}
	return nil
}
func (l lAdmin) UpdatePwdWithoutOldPwd(ctx context.Context, id interface{}, pwd interface{}) error {
	_, err := g.DB("sys").Model(dao.Admin.Table()).Ctx(ctx).Update(g.Map{"pwd": xpwd.GenPwd(pwd.(string))}, "id", id)
	if err != nil {
		return err
	}
	return nil
}
func (l lAdmin) Login(ctx context.Context, id string, code string, uname string, pwd string, ip string) error {
	if !xcaptcha.Store.Verify(id, code, true) {
		return errors.New("验证码错误")
	}
	admin, err := l.GetByUname(ctx, uname)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(admin.Pwd, pwd) {
		return consts.ErrLogin
	}

	if admin.Status == 2 {
		return consts.ErrAuthNotEnough
	}
	menus, err := Admin.ListMenus(ctx, admin.Rid, -1)
	if err != nil {
		return err
	}
	adminInfo := model.Admin{Admin: admin, Menus: menus}
	if err = Session.SetAdmin(ctx, &adminInfo); err != nil {
		return err
	}
	if _, err = g.DB("sys").Model(dao.AdminLoginLog.Table()).Ctx(ctx).Insert(do.AdminLoginLog{Uid: admin.Id, Ip: ip}); err != nil {
		return err
	}
	return nil
}

func (l lAdmin) MiddlewareAuth(r *ghttp.Request) {
	user, err := Session.GetAdmin(r.Session)
	if err != nil || user == nil {
		r.Response.RedirectTo("/admin/login")
		return
	}
	if !Role.CheckRoleApi(r.Context(), user.Admin.Rid, r.RequestURI) {
		switch r.Method {
		case "GET", "DELETE", "POST":
			res.Err(consts.ErrAuthNotEnough, r)
		default:
			res.Err(fmt.Errorf("权限不足"), r)
		}
	}
	r.Middleware.Next()
}
func (l lAdmin) MiddlewareLock(r *ghttp.Request) {
	var uid uint64
	getAdmin, err := Session.GetAdmin(r.Session)
	if err != nil {
		res.Err(err, r)
	}
	uid = uint64(getAdmin.Admin.Id)
	if uid == 0 {
		err := errors.New("uid is empty")
		g.Log().Error(nil, err)
		res.Err(err, r)
	}
	lock, err := xredis.UserLock(uid)
	if err != nil {
		res.Err(err, r)
	}
	r.Middleware.Next()
	lock.Unlock()
}
func (l lAdmin) MiddlewareActionLog(r *ghttp.Request) {
	user, err := Session.GetAdmin(r.Session)
	if err != nil || user == nil {
		res.Err(fmt.Errorf("用户信息错误"), r)
		return
	}
	uid := user.Admin.Id
	content := ""
	method := r.Method
	ctx := r.Context()
	uri := r.Router.Uri
	ip := r.GetClientIp()
	begin := time.Now().UnixMilli()
	response := ""
	if uri == "/Admin/operationLog/clear" {
		r.Middleware.Next()
		return
	}

	switch method {
	case "GET":
		content = r.GetUrl()
	case "DELETE":
		content = fmt.Sprintf("删除记录ID %s", r.Get("id").String())
	case "POST", "PUT":
		content = fmt.Sprint(r.GetFormMap())
		if content == "" {
			content = r.Request.PostForm.Encode()
		}
		if content == "" {
			content = r.Request.Form.Encode()
		}
		if len(content) > 233 {
			content = fmt.Sprint(gstr.SubStrRune(content, 0, 233), "...")
		}
	}
	r.Middleware.Next()
	useTime := time.Now().UnixMilli() - begin
	data := g.Map{
		"uid":      uid,
		"content":  content,
		"method":   method,
		"uri":      uri,
		"response": response,
		"use_time": useTime,
		"ip":       ip,
	}
	_, err = g.DB("sys").Model(dao.OperationLog.Table()).Ctx(ctx).Insert(data)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}
func (l lAdmin) MiddlewareUnread(r *ghttp.Request, aid int) error {
	get := r.Cookie.Get(consts.AdminUnreadKey)
	if get.Uint() == 0 {
		// get unread num from cache
		num, err := gcache.GetOrSetFunc(r.Context(), fmt.Sprint(consts.AdminUnreadKey, aid), func(ctx context.Context) (value interface{}, err error) {
			return g.DB("sys").Ctx(ctx).Model(dao.Admin.Table()).Value("unread_msg_num", "id", aid)
		}, time.Second*20)
		if err != nil {
			g.Log().Error(r.Context(), err)
			return err
		}
		r.Cookie.Set(consts.AdminUnreadKey, num.String())
	}
	return nil
}

type lFile struct{}

func (l lFile) Uploads(ctx context.Context, r *ghttp.Request) error {
	files := r.GetUploadFiles("file")
	if len(files) == 0 {
		return errors.New("lFile can't be empty")
	}
	for _, file := range files {
		fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
		file.Filename = fileName
	}
	datePre := time.Now().Format("2006/01")
	group := r.Get("group").String()
	if group == "" || group == "undefined" {
		group = "1"
	}
	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	mixPath := fmt.Sprintf("%s/%s/%s/", rootPath, group, datePre)
	_, err = files.Save(mixPath)
	if err != nil {
		return err
	}
	for _, file := range files {
		dbName := fmt.Sprintf("%s/%s/%s", group, datePre, file.Filename)
		_, err = g.DB("sys").Model(dao.File.Table()).Ctx(ctx).Insert(entity.File{
			Url:    dbName,
			Group:  gconv.Int(group),
			Status: 1,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func (l lFile) Upload(ctx context.Context, group int) (*v1.UploadFileRes, error) {
	file := ghttp.RequestFromCtx(ctx).GetUploadFile("file")
	if file == nil {
		return nil, consts.ErrImgCannotBeEmpty
	}
	fileName := fmt.Sprint(grand.S(6), path.Ext(file.Filename))
	file.Filename = fileName
	datePre := time.Now().Format("2006/01")

	rootFilePath, err := g.Cfg().Get(ctx, "server.rootFilePath")
	if err != nil {
		return nil, err
	}
	rootPath := gfile.Pwd() + rootFilePath.String()
	mixPath := fmt.Sprintf("%s/%d/%s/", rootPath, group, datePre)
	_, err = file.Save(mixPath)
	if err != nil {
		return nil, err
	}
	dbName := fmt.Sprintf("%d/%s/%s", group, datePre, file.Filename)
	_, err = g.DB("sys").Model(dao.File.Table()).Ctx(ctx).Insert(entity.File{
		Url:    dbName,
		Group:  group,
		Status: 1,
	})
	return &v1.UploadFileRes{
		DbName:    dbName,
		ImgPrefix: consts.ImgPrefix,
	}, err
}
func (l lFile) GetById(ctx context.Context, id interface{}) (*entity.File, error) {
	var data entity.File
	one, err := g.DB("sys").Model(dao.File.Table()).Ctx(ctx).Where("id", id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lFile) GetRandomUserIcon(ctx context.Context) (string, error) {
	value, err := g.DB("sys").Model("s_file").OrderRandom().Value("url", "`group` = 1")
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	return value.String(), nil
}
func (l lFile) ListIcons(ctx context.Context) ([]string, error) {
	array, err := g.DB("sys").Model("s_file").Array("url", "`group`=1")
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	var r []string
	for _, i := range array {
		r = append(r, i.String())
	}
	return r, nil
}

type lGen struct{}

func (l lGen) Gen(ctx context.Context, table string, group string, menu string, prefix string, apiGroup string, htmlGroup string, dbGroup string) error {
	// 结构体名称
	structName := gstr.CaseCamelLower(gstr.Replace(table, prefix, ""))
	// 表所有的字段
	fields, err := l.ListTableFields(ctx, table, dbGroup)
	if err != nil {
		return err
	}
	// 生成菜单
	if err = l.doGenMenu(ctx, group, menu, table, prefix); err != nil {
		return err
	}
	// 生成api
	if err = l.genApi(ctx, structName, menu, apiGroup); err != nil {
		return err
	}
	// 生成控制层
	if err = l.genController(table, htmlGroup, structName, dbGroup); err != nil {
		return err
	}
	// 生成 router
	if err = l.genRouter(structName); err != nil {
		return err
	}
	// 生成 html index
	if err = l.genIndex(htmlGroup, structName, fields); err != nil {
		return err
	}
	// 生成 html add
	if err = l.genAdd(htmlGroup, structName, menu, fields); err != nil {
		return err
	}
	// 生成 html edit
	if err = l.genEdit(htmlGroup, structName, menu, fields); err != nil {
		return err
	}
	return nil
}
func (l lGen) ListTableFields(ctx context.Context, table string, dbGroup string) ([]*gdb.TableField, error) {
	var (
		arr = make([]*gdb.TableField, 0)
	)
	fields, err := g.DB(dbGroup).TableFields(ctx, table)
	if err != nil {
		return nil, err
	}
	for _, v := range fields {
		arr = append(arr, v)
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i].Index < arr[j].Index })
	return arr, nil
}
func (l lGen) MenuLeve1(ctx context.Context) (string, error) {
	var (
		arr []string
	)
	all, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Array("name", "pid=-1")
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	for _, i := range all {
		arr = append(arr, i.String())
	}
	return strings.Join(arr, ","), nil
}
func (l lGen) TakeTables(ctx context.Context, db string) (string, error) {
	var (
		str []string
	)
	if db == "" {
		db = "sys"
	}
	tables, err := g.DB(db).Tables(ctx)
	for index, i := range tables {
		str = append(str, fmt.Sprintf("%s:%s:%s", i, i, xhtml.SwitchTagClass(index)))
	}
	return strings.Join(str, ","), err
}
func (l lGen) genEdit(htmlGroup, structName, menu string, fields []*gdb.TableField) error {
	structNameLower := gstr.CaseCamelLower(structName)
	editTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.edit.html"))
	// pageName
	pageName := menu
	editTemp = gstr.Replace(editTemp, "[pageName]", pageName)
	// menu
	editTemp = gstr.Replace(editTemp, "menu", structNameLower)

	// tr
	tr := ""
	for index, i := range fields {
		if index == 0 {
			switch strings.ToLower(i.Name) {
			case "id":
				tr += fmt.Sprintf(`{{editTrReadonly "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`{{editTr "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			}
		} else {
			switch i.Name {
			case "status":
				tr += fmt.Sprintf(`                        {{editTrOptions "%s" "%s" .Config.options.status .Form.%s}}
`, i.Name, i.Name, i.Name)
			case "updated_at", "created_at":
				tr += fmt.Sprintf(`                        {{editTrReadonly "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`                        {{editTr "%s" "%s" .Form.%s}}
`, i.Name, i.Name, i.Name)
			}
		}
	}
	editTemp = gstr.Replace(editTemp, "[tr]", tr)
	// date
	date := gtime.Now()
	editTemp = gstr.Replace(editTemp, "[date]", date.String())
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", htmlGroup, "/", structNameLower, "/edit.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(editTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
func (l lGen) genAdd(htmlGroup, structName, pageName string, fields []*gdb.TableField) error {
	addTemp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/temp.add.html"))
	addTemp = gstr.Replace(addTemp, "[pageName]", pageName)
	// menu
	addTemp = gstr.Replace(addTemp, "menu", gstr.CaseCamelLower(structName))

	// tr
	tr := ""
	for index, i := range fields {
		if index == 0 {
			switch strings.ToLower(i.Name) {
			case "id", "created_at", "updated_at":
				continue
			default:
				tr += fmt.Sprintf(`{{editTr "%s" "%s" ""}}
`, i.Name, i.Name)
			}
		} else {
			switch strings.ToLower(i.Name) {
			case "created_at", "updated_at", "id":
				continue
			case "status":
				tr += fmt.Sprintf(`                        {{editTrOptions "%s" "%s" .Config.options.status 1}}
`, i.Name, i.Name)
			default:
				tr += fmt.Sprintf(`                        {{editTr "%s" "%s" ""}}
`, i.Name, i.Name)
			}
		}
	}
	addTemp = gstr.Replace(addTemp, "[tr]", tr)
	// date
	date := gtime.Now()
	addTemp = gstr.Replace(addTemp, "[date]", date.String())

	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", htmlGroup, "/", gstr.CaseCamelLower(structName), "/add.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(addTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
func (l lGen) genIndex(htmlGroup, structName string, fields []*gdb.TableField) error {
	indexTemp := gfile.GetContents(fmt.Sprintf("%s/resource/gen/temp.index.html", gfile.MainPkgPath()))
	group := htmlGroup
	structNameLower := gstr.CaseCamelLower(structName)
	// Menu
	caseCamel := gstr.CaseCamel(structName)
	indexTemp = gstr.Replace(indexTemp, "Menu", caseCamel)
	indexTemp = gstr.Replace(indexTemp, "menu", gstr.CaseCamelLower(structName))
	// th
	arr := make([]string, 0)
	for _, i := range fields {
		arr = append(arr, strings.ToUpper(i.Name))
	}
	arr = append(arr, "OPERATION")
	th := strings.Join(arr, ",")
	indexTemp = gstr.Replace(indexTemp, "[th]", th)
	// td
	td := ""
	for index, i := range fields {
		if index == 0 { // 如果是第一个
			td += fmt.Sprintf(`{{td "%s" .%s}}
`, i.Name, i.Name)
		} else {
			switch strings.ToLower(i.Name) {
			case "status":
				td += fmt.Sprintf(`                        {{tdChoose "%s" $.Config.options.status .%s}}
`, i.Name, i.Name)
			default:
				td += fmt.Sprintf(`                        {{td "%s" .%s}}
`, i.Name, i.Name)
			}
		}
	}
	indexTemp = gstr.Replace(indexTemp, "[td]", td)

	// date
	date := gtime.Now()
	indexTemp = gstr.Replace(indexTemp, "[date]", date.String())
	f, err := gfile.Create(fmt.Sprint(gfile.MainPkgPath(), "/resource/template/", group, "/", structNameLower, "/index.html"))
	if err != nil {
		return err
	}
	if _, err = f.WriteString(indexTemp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
func (l lGen) genRouter(name string) error {
	temp := gfile.GetContents(fmt.Sprint(gfile.MainPkgPath(), "/resource/gen/router.temp"))
	structName := gstr.CaseCamelLower(name)
	caseCamel := gstr.CaseCamel(structName)
	temp = gstr.Replace(temp, "menu", structName)
	temp = gstr.Replace(temp, "Menu", caseCamel)

	// sys_router
	sysRouterPath := fmt.Sprint(gfile.MainPkgPath(), "/internal/cmd/cmd.go")
	sysRouter := gfile.GetContents(sysRouterPath)
	if gstr.Contains(sysRouter, temp) {
		return nil
	}
	stat, err := gfile.Stat(sysRouterPath)
	if err != nil {
		return err
	}
	if err := gfile.Truncate(sysRouterPath, int(stat.Size()-2)); err != nil {
		return err
	}
	if err := gfile.PutContentsAppend(sysRouterPath, temp); err != nil {
		return err
	}
	return nil
}
func (l lGen) genController(table string, htmlGroup string, structName string, dbGroup string) error {
	pwd := gfile.MainPkgPath()
	line, err := xfile.ReadLine(fmt.Sprint(pwd, "/go.mod"), 1)
	if err != nil {
		return err
	}
	// mod
	mod := gstr.SplitAndTrim(line, " ")[1]
	temp := gfile.GetContents(fmt.Sprint(pwd, "/resource/gen/controller.temp"))
	temp = gstr.Replace(temp, "[mod]", mod)

	// dbGroup
	temp = gstr.Replace(temp, "[dbGroup]", dbGroup)

	// Menu
	caseCamel := gstr.CaseCamel(structName)
	temp = gstr.Replace(temp, "Menu", caseCamel)
	temp = gstr.Replace(temp, "menu", gstr.CaseCamelLower(structName))

	// group
	temp = gstr.Replace(temp, "[group]", htmlGroup)

	// table
	temp = gstr.Replace(temp, "[table]", table)

	// htmlGroup
	temp = gstr.Replace(temp, "[htmlGroup]", htmlGroup)

	// date
	date := gtime.Now()
	temp = gstr.Replace(temp, "[date]", date.String())

	// lFile
	filePath := fmt.Sprint(pwd, "/internal/controller/", table, ".go")
	f, err := gfile.Create(filePath)
	if err != nil {
		return err
	}
	if _, err = f.WriteString(temp); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}
func (l lGen) genApi(ctx context.Context, name, pageName, group string) error {
	// 检查在字典表中是否存在
	if err := l.checkGroupOrSave(ctx, group); err != nil {
		return err
	}
	if pageName == "" {
		pageName = name
	}
	name = gstr.CaseCamelLower(name)
	array := []*entity.Api{
		{Url: fmt.Sprintf("/%s", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s页面", pageName), Type: 5},
		{Url: fmt.Sprintf("/%s/add", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s添加页面", pageName), Type: 5},
		{Url: fmt.Sprintf("/%s/edit/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s修改页面", pageName), Type: 5},
		{Url: fmt.Sprintf("/%s/del/:id", name), Method: "1", Group: group, Desc: fmt.Sprintf("%s删除操作", pageName), Type: 2},
		{Url: fmt.Sprintf("/%s", name), Method: "2", Group: group, Desc: fmt.Sprintf("添加%s", pageName), Type: 1},
		{Url: fmt.Sprintf("/%s", name), Method: "2", Group: group, Desc: fmt.Sprintf("修改%s", pageName), Type: 3},
	}
	for _, i := range array {
		count, err := g.DB("sys").Model(dao.Api.Table()).Ctx(ctx).Count("url = ? and method = ?", i.Url, i.Method)
		if err != nil {
			return err
		}
		if count != 0 {
			continue
		}
		if _, err = g.DB("sys").Model(dao.Api.Table()).Ctx(ctx).Insert(i); err != nil {
			return err
		}
	}
	return nil
}
func (l lGen) doGenMenu(ctx context.Context, group, menu, table, prefix string) error {
	var (
		m1Sort, m2Sort = 0.0, 0.0
	)
	menu1, err := Menu.GetByName(ctx, group)
	if err != nil {
		if err == consts.ErrDataNotFound {
			g.Log().Debug(ctx, "一级菜单不存在")
			// 新增一级菜单
			maxSort, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Max("sort")
			if err != nil {
				return err
			}
			m1Sort = math.Ceil(maxSort)
			m2Sort = m1Sort + 0.1
			id, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).InsertAndGetId(&entity.Menu{
				Pid:    -1,
				Name:   group,
				Type:   2,
				Sort:   m1Sort,
				Status: 1,
			})
			if err != nil {
				return err
			}
			g.Log().Infof(ctx, "新增一级菜单,排序为%v", m1Sort)
			menu1 = &entity.Menu{Id: int(id)}
			goto here
		}
		return err
	} else {
		// select max sort from menu1'children
		childrenMaxSort, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Where("pid=?", menu1.Id).Max("sort")
		if err != nil {
			return err
		}
		if childrenMaxSort == 0 {
			m2Sort += menu1.Sort + 0.1
		} else {
			m2Sort += childrenMaxSort + 0.1
		}
		g.Log().Infof(ctx, "查询一级菜单，子菜单最大排序为%v", menu1.Sort)
	}
	if menu1.Type != 2 {
		return errors.New("一级菜单必须为分组菜单")
	}
here:
	// 查看菜单是否存在
	count, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Count("name", menu)
	if count != 0 {
		g.Log().Warningf(ctx, "%s 菜单已存在，就不创建啦", menu)
		return nil
	}
	// 新增二级菜单
	menuPath := fmt.Sprintf("/admin/%s", gstr.CaseCamelLower(gstr.Replace(table, prefix, "")))
	// count path
	g.Log().Debug(ctx, "检查二级菜单是否存在")
	pathCount, err := g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Where("path", menuPath).Count()
	if err != nil {
		return err
	}
	if pathCount > 0 {
		g.Log().Warning(ctx, "菜单路径已存在,未执行插入菜单操作")
		return nil
	}
	//menuLogo := xicon.GenIcon()
	if _, err = g.DB("sys").Model(dao.Menu.Table()).Ctx(ctx).Insert(&entity.Menu{
		Pid: menu1.Id,
		//Icon:   menuLogo,
		//BgImg:  menuLogo,
		Path:   menuPath,
		Sort:   m2Sort,
		Name:   menu,
		Status: 1,
		Type:   1,
	}); err != nil {
		return err
	}
	g.Log().Debugf(ctx, "新增二级菜单,排序为%v", m2Sort)
	return nil
}
func (l lGen) checkGroupOrSave(ctx context.Context, group string) error {
	d, err := Dict.GetByKey(ctx, "api_group")
	if err != nil {
		return err
	}
	for _, i := range gstr.Split(d.V, "\n") {
		i = gstr.TrimAll(i)
		if i == group {
			return nil
		}
	}
	d.V += fmt.Sprint("\n", group)
	if _, err = g.DB("sys").Model(dao.Dict.Table()).Ctx(ctx).Save(d); err != nil {
		return err
	}
	g.Log().Warningf(ctx, "%s 分组在词典表中不存在，已添加.", group)
	return nil
}

type lSession struct{}

func (s lSession) SetAdmin(ctx context.Context, b *model.Admin) error {
	return g.RequestFromCtx(ctx).Session.Set(AdminSessionKey, b)
}
func (lSession) GetAdmin(r *ghttp.Session) (*model.Admin, error) {
	get, err := r.Get(AdminSessionKey)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, errors.New("lAdmin info is nil")
	}
	var data *model.Admin
	err = get.Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s lSession) DelAdmin(ctx context.Context) error {
	return g.RequestFromCtx(ctx).Session.Remove(AdminSessionKey)
}

type lWs struct{}

func (l lWs) GetUserWs(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		res.Err(err, r)
	}
	uid := User.GetUidFromRequest(r)
	users.Set(uid, ws)
	l.printUserWs()
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			users.Remove(uid)
			l.printUserWs()
			return
		}
		g.Log().Info(gctx.New(), "ws:lUser msg ", messageType, msg)
	}
}
func (l lWs) GetAdminWs(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		res.Err(err, r)
	}
	adminBo, err := Session.GetAdmin(r.Session)
	if err != nil || adminBo == nil {
		res.Err(err, r)
		return
	}

	id := adminBo.Admin.Id
	admins.Set(id, ws)
	l.printAdminWs()
	for {
		messageType, msg, err := ws.ReadMessage()
		if err != nil {
			admins.Remove(id)
			l.printAdminWs()
			return
		}
		g.Log().Info(gctx.New(), "ws:lAdmin msg ", messageType, msg)
	}
}
func (l lWs) NoticeUsers(ctx context.Context, msg interface{}) error {
	if users.Size() == 0 {
		return nil
	}
	marshal, _ := json.Marshal(msg)
	for _, item := range users.Values() {
		if err := item.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) NoticeAdmin(ctx context.Context, msg interface{}, uid int) error {
	to := admins.Get(uid)
	if to != nil {
		marshal, _ := json.Marshal(msg)
		if err := to.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) NoticeAdmins(ctx context.Context, msg interface{}, info ...entity.AdminMessage) error {
	marshal, _ := json.Marshal(msg)
	for _, id := range admins.Keys() {
		if err := admins.Get(id).(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		if err := Admin.AddUnReadMsgNum(ctx, gconv.Int(id), 1); err != nil {
			return err
		}
		if len(info) != 0 {
			info[0].Aid = gconv.Uint64(id)
			if err := Admin.AddUnreadMsg(ctx, info[0]); err != nil {
				return err
			}
		}
	}
	return nil
}
func (l lWs) NoticeUser(ctx context.Context, uid int, msg interface{}) error {
	marshal, _ := json.Marshal(msg)
	item := users.Get(uid)
	if item != nil {
		if err := item.(*ghttp.WebSocket).WriteMessage(1, marshal); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	}
	return nil
}
func (l lWs) printUserWs() {
	g.Log().Infof(gctx.New(), "user连接个数%v %v", len(users.Map()), users.Keys())
}
func (l lWs) printAdminWs() {
	//g.Log().Infof(gctx.New(), "admin连接个数%v %v", len(admins.Map()), admins.Keys())
}

type lUser struct{}

func (l lUser) Add(ctx context.Context, input *v1.RegisterReq) (*v1.LoginRes, error) {
	ip := ghttp.RequestFromCtx(ctx).GetClientIp()
	if err := l.checkUnameAlreadyExist(ctx, input.Uname); err != nil {
		return nil, err
	}
	var (
		resVo    v1.LoginRes
		userData = entity.User{
			Uname:    input.Uname,
			Nickname: input.Uname,
			JoinIp:   ip,
			Pass:     xpwd.GenPwd(input.Pass),
			Status:   1,
		}
	)
	icon, err := File.GetRandomUserIcon(ctx)
	if err != nil {
		return nil, err
	}
	userData.Icon = icon
	if err = g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		uid, err := tx.Model(dao.User.Table()).InsertAndGetId(userData)
		if err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		userData.Id = uint64(uid)
		var wallet = entity.Wallet{}
		wallet.Uid = userData.Id
		if _, err = tx.Model(dao.Wallet.Table()).Insert(wallet); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		vo, err := l.makeLoginRes(&userData, &wallet)
		if err != nil {
			return err
		}
		resVo = *vo
		return nil
	}); err != nil {
		return nil, err
	}

	// 设置用户登陆信息
	l.addInfoToCookie(ctx, userData.Id, userData.Uname, userData.Icon, 0)
	return &resVo, nil
}
func (l lUser) addInfoToCookie(ctx context.Context, uid uint64, uname string, icon string, i float64) {
	r := ghttp.RequestFromCtx(ctx)
	r.Cookie.Set("uname", uname)
	r.Cookie.Set("icon", icon)
	r.Cookie.Set("uid", fmt.Sprint(uid))
	r.Cookie.Set("balance", fmt.Sprint(i))
}
func (l lUser) Login(ctx context.Context, in *v1.LoginReq) (*v1.LoginRes, error) {
	userData, err := l.GetByUname(ctx, in.Uname)
	if err != nil {
		return nil, consts.ErrUserDoesNotExist
	}
	if userData.PassErrorCount > 6 {
		return nil, consts.ErrPassErrorTooMany
	}
	if !xpwd.ComparePassword(userData.Pass, in.Pass) {
		userData.PassErrorCount++
		if userData.PassErrorCount >= 6 {
			userData.Status = consts.UserStatusLock
		}
		if _, err = g.DB("sys").Model(dao.User.Table()).Ctx(ctx).Save(userData); err != nil {
			return nil, err
		}
		return nil, consts.ErrLogin
	}
	if err = g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		var loginLog = entity.UserLoginLog{
			Uid: userData.Id,
			Ip:  g.RequestFromCtx(ctx).GetClientIp(),
		}
		if userData.PassErrorCount != 0 {
			if _, err = tx.Model(dao.User.Table()).WherePri(userData.Id).Update(do.User{PassErrorCount: 0}); err != nil {
				g.Log().Error(ctx, err)
				return err
			}
		}
		if _, err = tx.Model(dao.UserLoginLog.Table()).Insert(loginLog); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	wallet, err := Wallet.GetByUid(ctx, userData.Id)
	if err != nil {
		return nil, err
	}
	res, err := l.makeLoginRes(userData, wallet)
	if err != nil {
		return nil, err
	}
	l.addInfoToCookie(ctx, userData.Id, userData.Uname, userData.Icon, wallet.Balance)
	return res, nil
}

func (l lUser) GetUidFromCookie(ctx context.Context) uint64 {
	r := ghttp.RequestFromCtx(ctx)
	uid := r.Cookie.Get("uid")
	return gconv.Uint64(uid)
}
func (l lUser) GetById(ctx context.Context, id uint64) (*entity.User, error) {
	var data entity.User
	one, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).WherePri(id).One()
	if err != nil {
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lUser) GetByIdTx(ctx context.Context, tx *gdb.TX, id uint64) (*entity.User, error) {
	var data entity.User
	one, err := tx.Ctx(ctx).Model(dao.User.Table()).WherePri(id).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lUser) GetByUname(ctx context.Context, uname string) (*entity.User, error) {
	var data entity.User
	one, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).One("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrLogin
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lUser) GetInfo(ctx context.Context, uid uint64) (*v1.LoginRes, error) {
	user, err := l.GetById(ctx, uid)
	if err != nil {
		return nil, err
	}
	wallet, err := Wallet.GetByUid(ctx, uid)
	if err != nil {
		return nil, err
	}
	return l.makeLoginRes(user, wallet)
}
func (l lUser) getLoginVoWithTx(ctx context.Context, tx *gdb.TX, id uint64) (*v1.LoginRes, error) {
	userData, err := l.GetByIdTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	wallet, err := Wallet.GetByUidTx(ctx, tx, id)
	if err != nil {
		return nil, err
	}
	return l.makeLoginRes(userData, wallet)
}
func (l lUser) GetUidFromRequest(r *ghttp.Request) uint64 {
	return r.Get(consts.UidKey).Uint64()
}
func (l lUser) GetUidFromCtx(ctx context.Context) uint64 {
	return l.GetUidFromRequest(ghttp.RequestFromCtx(ctx))
}

func (l lUser) Del(ctx context.Context, id uint64) error {
	if _, err := g.DB("sys").Model("u_user").Delete("id", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_user_login_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet_change_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if _, err := g.DB("sys").Model("u_wallet_statistics_log").Delete("uid", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) DelLoinLogs(ctx context.Context) error {
	if _, err := g.DB("sys").Model(dao.UserLoginLog.Table()).Ctx(ctx).Delete("id is not null"); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) Logout(ctx context.Context) {
	r := ghttp.RequestFromCtx(ctx)
	r.Cookie.Remove("uname")
	r.Cookie.Remove("icon")
	r.Cookie.Remove("uid")
	r.Cookie.Remove("balance")
}

func (l lUser) UpdateUname(ctx context.Context, uname string, id uint64) error {
	count, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).Count("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	if err = g.Validator().Rules("password").Data(uname).Run(ctx); err != nil {
		return consts.ErrUnameFormat
	}
	if _, err = dao.User.Ctx(ctx).WherePri(id).Update(g.Map{"uname": uname}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) UpdatePass(ctx context.Context, pass string, id uint64) error {
	if err := g.Validator().Rules("password").Data(pass).Run(ctx); err != nil {
		return consts.ErrPassFormat
	}
	if _, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).WherePri(id).Update(g.Map{"pass": xpwd.GenPwd(pass)}); err != nil {
		g.Log().Error(ctx)
		return err
	}
	return nil
}
func (l lUser) UpdatePassByUser(ctx context.Context, in *v1.UpdatePassReq, id uint64) error {
	userData, err := l.GetById(ctx, id)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(userData.Pass, in.OldPass) {
		return consts.ErrOldPassNotMatch
	}
	data := do.User{Pass: xpwd.GenPwd(in.NewPass)}
	if _, err = g.DB("sys").Model(dao.User.Table()).Ctx(ctx).Update(data, "id", id); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) UpdateNickname(ctx context.Context, nickname string, uid uint64) error {
	if len(nickname) > 16 {
		return consts.ErrMaxLengthSixTy
	}
	if _, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).WherePri(uid).Update(do.User{Nickname: nickname}); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) UpdateIcon(ctx context.Context, icon string, uid uint64) error {
	if _, err := g.DB("sys").Model(dao.User.Table()).Ctx(ctx).WherePri(uid).Data(do.User{Icon: icon}).Update(); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lUser) checkUnameAlreadyExist(ctx context.Context, uname string) error {
	count, err := g.DB("sys").Model("u_user").Ctx(ctx).Count("uname", uname)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 0 {
		return consts.ErrUnameExist
	}
	return nil
}
func (l lUser) makeLoginRes(userData *entity.User, wallet *entity.Wallet) (*v1.LoginRes, error) {
	var res v1.LoginRes
	res.Uname = userData.Uname
	res.Nickname = userData.Nickname
	res.Email = userData.Email
	res.Phone = userData.Phone
	res.Summary = userData.Summary
	if strings.HasPrefix(userData.Icon, "http") {
		res.Icon = userData.Icon
	} else {
		res.Icon = consts.ImgPrefix + userData.Icon
	}
	res.WalletStatus = wallet.Status
	token, err := xjwt.GenToken(userData.Uname, userData.Id, 0)
	if err != nil {
		return nil, err
	}
	res.Token = token
	return &res, nil
}

type lWallet struct{}

func (l lWallet) AddChangeLog(ctx context.Context, tx *gdb.TX, transId string, t int, uid uint64, amount float64, balance float64, desc string) error {
	var (
		data = do.WalletChangeLog{
			TransId: transId,
			Uid:     uid,
			Type:    t,
			Amount:  amount,
			Balance: balance,
			Desc:    desc,
		}
	)
	if _, err := tx.Model(dao.WalletChangeLog.Table()).Insert(data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lWallet) AddStatisticsLog(ctx context.Context, tx *gdb.TX, t int, uid uint64, amount float64) error {
	todayLog, err := l.GetStatisticsTodayLog(ctx, tx, uid)
	if err != nil {
		if err != consts.ErrDataNotFound {
			return err
		}
		data := g.Map{
			"uid":                 uid,
			"created_date":        time.Now(),
			fmt.Sprintf("t%d", t): math.Abs(amount),
		}
		if _, err = tx.Model(dao.WalletStatisticsLog.Table()).Insert(data); err != nil {
			return err
		}
		return nil
	}
	if _, err = tx.Model(dao.WalletStatisticsLog.Table()).
		WherePri(todayLog.Id).
		Increment(fmt.Sprintf("t%d", t), math.Abs(amount)); err != nil {
		return err
	}
	return nil
}
func (l lWallet) AddTopUp(ctx context.Context, money float64, changeTypeId int, uid uint64) (err error) {
	// Check money
	if money < 10 || money > 10000 {
		return consts.ErrMinTopUpOrderMoney
	}
	// Check if the user has a pending order
	count, err := g.DB("sys").Model("u_wallet_top_up_application").Ctx(ctx).Count("uid = ? and status = 1", uid)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 0 {
		return consts.ErrTopUpOrderAlreadyHas
	}
	// Check changeTypeId is correct
	count, err = g.DB("sys").Model("u_wallet_change_type").Ctx(ctx).Count("id = ? and `type` = 1", changeTypeId)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	if count != 1 {
		return consts.ErrTopUpType
	}
	// Create TopUpOrder
	transId := fmt.Sprint("MR-", grand.S(13))
	order := entity.WalletTopUpApplication{
		Uid:        uid,
		TransId:    transId,
		ChangeType: uint(changeTypeId),
		Money:      money,
		Ip:         ghttp.RequestFromCtx(ctx).GetClientIp(),
		Status:     1,
		Aid:        0,
	}
	if _, err = g.DB("sys").Model("u_wallet_top_up_application").Ctx(ctx).Insert(order); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	userInfo, err := User.GetById(ctx, uid)
	if err != nil {
		return err
	}
	// Notice admins
	if err = Ws.NoticeAdmins(ctx, "", entity.AdminMessage{
		Type:  1,
		Title: fmt.Sprintf("%s创建了%.2f充值订单", userInfo.Uname, money),
		Url:   fmt.Sprintf("/admin/walletTopUpApplication?trans_id=%s", transId),
	}); err != nil {
		return err
	}
	return
}

func (l lWallet) GetChangeWallet(ctx context.Context, tx *gdb.TX, t int, uid uint64, amount float64) (*entity.Wallet, error) {
	wallet, err := l.GetByUidTx(ctx, tx, uid)
	if err != nil {
		return nil, err
	}
	wallet.Balance += amount
	if wallet.Balance < 0 {
		wallet.Balance = 0
	}
	var data = do.Wallet{Balance: wallet.Balance}
	if _, err = tx.Model(dao.Wallet.Table()).WherePri(wallet.Id).Data(data).Update(); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return wallet, nil
}
func (l lWallet) GetStatisticsLogReport(ctx context.Context, begin string, end string, uname string) (gdb.Record, error) {
	if begin == "" {
		begin = gtime.Now().AddDate(0, -6, 0).StartOfDay().String()
	}
	db := g.DB("sys").Model(dao.WalletStatisticsLog.Table()+" t1").Ctx(ctx).
		FieldSum("t1.t1", "t1").
		FieldSum("t1.t2", "t2").
		FieldSum("t1.t3", "t3").
		FieldSum("t1.t4", "t4").
		FieldSum("t1.t5", "t5").
		WhereGTE("t1.created_date", begin)
	if end != "" {
		db = db.WhereLTE("t1.created_date", end)
	}
	if uname != "" {
		db = db.LeftJoin("u_user t2 on t1.uid = t2.id").Where("t2.uname", uname)
	}
	one, err := db.One()
	if err != nil {
		return nil, err
	}
	return one, nil
}
func (l lWallet) GetByUidTx(ctx context.Context, tx *gdb.TX, id uint64) (*entity.Wallet, error) {
	var data entity.Wallet
	one, err := tx.Ctx(ctx).Model(dao.Wallet.Table()).One("uid", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		g.Log().Errorf(ctx, "%d 钱包信息不存在", id)
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lWallet) GetStatisticsTodayLog(ctx context.Context, tx *gdb.TX, uid uint64) (*entity.WalletStatisticsLog, error) {
	var data entity.WalletStatisticsLog
	one, err := tx.Model(dao.WalletStatisticsLog.Table()).Where("uid = ? and created_date>=?", uid, gtime.Date()).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lWallet) GetByUid(ctx context.Context, uid uint64) (*entity.Wallet, error) {
	var data entity.Wallet
	one, err := g.DB("sys").Model(dao.Wallet.Table()).Ctx(ctx).Where("uid", uid).One()
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lWallet) GetChangeTypeOptions(ctx context.Context) (string, error) {
	changeTypes, err := l.listChangeTypesByType(ctx)
	if err != nil {
		return "", err
	}
	var arr []string
	for _, i := range changeTypes {
		arr = append(arr, fmt.Sprintf(`%d:%s:%s`, i.Id, i.Title, i.Class))
	}
	return strings.Join(arr, ","), nil
}
func (l lWallet) GetChangeTypeTopUpOptions(ctx context.Context) (string, error) {
	changeTypes, err := l.listChangeTypesByType(ctx, 1)
	if err != nil {
		return "", err
	}
	var arr []string
	for _, i := range changeTypes {
		arr = append(arr, fmt.Sprintf("%d:%s:%s", i.Id, i.Title, i.Class))
	}
	return strings.Join(arr, ","), nil
}
func (l lWallet) GetChangeTypeDeductOptions(ctx context.Context) (string, error) {
	changeTypes, err := l.listChangeTypesByType(ctx, 2)
	if err != nil {
		return "", err
	}
	var arr []string
	for _, i := range changeTypes {
		arr = append(arr, fmt.Sprintf("%d:%s:%s", i.Id, i.Title, i.Class))
	}
	return strings.Join(arr, ","), nil
}
func (l lWallet) GetInfo(ctx context.Context, uid uint64) (*v1.WalletInfoRes, error) {
	var data v1.WalletInfoRes
	one, err := g.DB("sys").Ctx(ctx).Model(dao.Wallet.Table()).One("uid", uid)
	if err != nil {
		return nil, err
	}
	if err = one.Struct(&data); err != nil {
		return nil, err
	}
	return &data, nil
}
func (l lWallet) GetTopUpApplication(ctx context.Context, id uint64) (*entity.WalletTopUpApplication, error) {
	var data entity.WalletTopUpApplication
	one, err := g.DB("sys").Ctx(ctx).Model(dao.WalletTopUpApplication.Table()).One("id", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	if one.IsEmpty() {
		return nil, consts.ErrDataNotFound
	}
	if err = one.Struct(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return &data, nil
}
func (l lWallet) ListTopUpByUid(ctx context.Context, page int, size int, status int, uid uint64) (int64, []*model.TopUpItem, error) {
	var data = make([]*model.TopUpItem, 0)
	db := g.DB("sys").Model(dao.WalletTopUpApplication.Table()).Ctx(ctx).Where("uid", uid)
	if status != 0 {
		db.Where("status", status)
	}
	total, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return 0, nil, err
	}
	if err = db.Page(page, size).OrderDesc("id").Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return 0, nil, err
	}
	return total, data, nil
}
func (l lWallet) ListChangeTypes(ctx context.Context) ([]*v1.ListChangeTypesRes, error) {
	var data = make([]*v1.ListChangeTypesRes, 0)
	if err := g.DB("sys").Model(dao.WalletChangeType.Table()).Ctx(ctx).Where("status", 1).Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lWallet) ListChangeLogs(ctx context.Context, page int, size int, t int, uid uint64) (int64, []*model.ChangeLogItem, error) {
	// This is the data we need
	var data = make([]*model.ChangeLogItem, 0)
	// Get query db
	db := g.DB("sys").Ctx(ctx).Model(dao.WalletChangeLog.Table()).Where("uid", uid)
	if t != 0 {
		db = db.Where("type", t)
	}
	// Query total number
	total, err := db.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		return 0, nil, err
	}
	// Paging query data
	if err = db.Page(page, size).OrderDesc("id").Scan(&data); err != nil {
		return 0, nil, err
	}
	return total, data, nil
}
func (l lWallet) listChangeTypesByType(ctx context.Context, t ...int) ([]*entity.WalletChangeType, error) {
	var data []*entity.WalletChangeType
	db := g.DB("sys").Model(dao.WalletChangeType.Table())
	if len(t) != 0 {
		db = db.Where("type", t[0])
	}
	if err := db.Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lWallet) listChangeTypesByCountStatus(ctx context.Context, t ...int) ([]*entity.WalletChangeType, error) {
	var data []*entity.WalletChangeType
	db := g.DB("sys").Model(dao.WalletChangeType.Table())
	if len(t) != 0 {
		db = db.Where("count_status", t[0])
	}
	if err := db.Scan(&data); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return data, nil
}
func (l lWallet) ListTopUpCategory(ctx context.Context) ([]*v1.TopUpCategoryRes, error) {
	var res []*v1.TopUpCategoryRes
	if err := g.DB("sys").Model("u_wallet_change_type").Scan(&res, "status=1 and type = 1"); err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return res, nil
}

func (l lWallet) DelChangeLogs(ctx context.Context) error {
	if _, err := g.DB("sys").Model(dao.WalletChangeLog.Table()).Ctx(ctx).Delete("id is not null"); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lWallet) DelStatisticsLogs(ctx context.Context) error {
	_, err := g.DB("sys").Model(dao.WalletStatisticsLog.Table()).Ctx(ctx).Delete("id is not null")
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l lWallet) UpdatePassByAdmin(ctx context.Context, pass string, uid uint64) error {
	if _, err := g.DB("sys").Model(dao.Wallet.Table()).Ctx(ctx).Update(do.Wallet{Pass: xpwd.GenPwd(pass)}, "uid", uid); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func (l lWallet) UpdateTopUpByAdmin(ctx context.Context, t int, uid uint64, amount float64, desc string) error {
	if err := g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		// 修改用户金币
		walletInfo, err := l.GetChangeWallet(ctx, tx, t, uid, amount)
		if err != nil {
			return err
		}
		// 创建账变记录
		transId := guid.S()
		if desc == "" {
			desc = "人工充值"
		}
		if err = l.AddChangeLog(ctx, tx, transId, t, uid, amount, walletInfo.Balance, desc); err != nil {
			return err
		}
		// 创建账变统计
		if err = l.AddStatisticsLog(ctx, tx, t, uid, amount); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
func (l lWallet) UpdateDeductByAdmin(ctx context.Context, t int, uid uint64, amount float64) error {
	if err := g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		walletInfo, err := l.GetChangeWallet(ctx, tx, t, uid, amount)
		if err != nil {
			return err
		}
		// 创建账变记录
		transId := guid.S()
		if err = l.AddChangeLog(ctx, tx, transId, t, uid, amount, walletInfo.Balance, "人工扣除"); err != nil {
			return err
		}
		// 创建账变统计
		if err = l.AddStatisticsLog(ctx, tx, t, uid, amount); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
func (l lWallet) UpdateChangeWalletKit(ctx context.Context, tx *gdb.TX, changeType int, uid uint64, money float64, transId string, desc string) error {
	// add user money
	wallet, err := l.GetChangeWallet(ctx, tx, changeType, uid, money)
	if err != nil {
		return err
	}
	// add change log
	if err = l.AddChangeLog(ctx, tx, transId, changeType, uid, money, wallet.Balance, desc); err != nil {
		return err
	}
	// add statistics log
	if err = l.AddStatisticsLog(ctx, tx, changeType, uid, money); err != nil {
		return err
	}
	return nil
}
func (l lWallet) UpdateSetPass(ctx context.Context, pass string, uid uint64) error {
	if err := l.checkPassFormat(pass); err != nil {
		return err
	}
	count, err := g.DB("sys").Ctx(ctx).Model("u_wallet").Count("uid = ? and status = 0", uid)
	if err != nil {
		return err
	}
	if count != 1 {
		return consts.ErrUseWalletPassAlreadySet
	}
	if _, err := g.DB("sys").Ctx(ctx).Model("u_wallet").Update(g.Map{
		"pass":   xpwd.GenPwd(pass),
		"status": 1,
	}, "uid", uid); err != nil {
		return err
	}
	return nil
}
func (l lWallet) UpdatePass(ctx context.Context, oldPass string, newPass string, uid uint64) error {
	wallet, err := l.GetByUid(ctx, uid)
	if err != nil {
		return err
	}
	if !xpwd.ComparePassword(wallet.Pass, oldPass) {
		return consts.ErrOldPassNotMatch
	}
	if err = l.checkPassFormat(newPass); err != nil {
		return err
	}
	if _, err = g.DB("sys").Ctx(ctx).Model("u_wallet").Update(g.Map{"pass": xpwd.GenPwd(newPass)}, "uid", uid); err != nil {
		return err
	}
	return nil
}
func (l lWallet) UpdateTopUpApplication(ctx context.Context, id uint64, operationType int64, aid int) error {
	application, err := l.GetTopUpApplication(ctx, id)
	if err != nil {
		return err
	}
	switch operationType {
	case 2: // fail
		application.Status = consts.ApplicationStatusFail
		application.Aid = uint64(aid)
		if _, err := g.DB("sys").Model(dao.WalletTopUpApplication.Table()).Save(application); err != nil {
			g.Log().Error(ctx, err)
			return err
		}
	case 1: // ok
		if err = g.DB("sys").Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
			// update application
			if _, err := tx.Model(dao.WalletTopUpApplication.Table()).Update(g.Map{
				"status": consts.ApplicationStatusSuccess,
				"aid":    aid,
			}, "id", id); err != nil {
				g.Log().Error(ctx, err)
				return err
			}
			// change user wallet
			if err = l.UpdateChangeWalletKit(ctx, tx, int(application.ChangeType), application.Uid, application.Money, application.TransId, ""); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

func (l lWallet) TakeStatisticsLogFieldsNeedToBeCountedOptionsIntoStr(ctx context.Context) (string, error) {
	changeTypes, err := l.listChangeTypesByCountStatus(ctx, 1)
	if err != nil {
		g.Log().Error(ctx, err)
		return "", err
	}
	var (
		th = []string{"ID", "用户名", "日期"}
	)
	for _, i := range changeTypes {
		th = append(th, i.Title)
	}
	th = append(th, "OPERATION")

	return strings.Join(th, ","), nil
}
func (l lWallet) TakeStatisticsLogFieldsNeedToBeCountedOptionsIntoArray(ctx context.Context) ([]string, error) {
	changeTypes, err := l.listChangeTypesByCountStatus(ctx, 1)
	if err != nil {
		return nil, err
	}
	var array []string
	for _, i := range changeTypes {
		array = append(array, i.Title)
	}
	return array, nil
}

func (l lWallet) checkPassFormat(pass string) error {
	if !gstr.IsNumeric(pass) {
		return consts.ErrFormatNotNumber
	}
	if len(pass) != 6 {
		return consts.ErrFormatKeepLengthSix
	}
	return nil
}
