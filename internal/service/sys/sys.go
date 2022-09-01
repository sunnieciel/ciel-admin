// Package sys 系统服务层
package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/logic"
	"ciel-admin/internal/model/bo"
	"ciel-admin/internal/model/entity"
	"ciel-admin/utility/utils/res"
	"ciel-admin/utility/utils/xstr"
	"ciel-admin/utility/utils/xtime"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func Init(ctx context.Context) {
	get, err := g.Cfg().Get(ctx, "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
	if err = logic.Dict.SetWhiteIps(ctx); err != nil {
		panic(err)
	}
}
func List(ctx context.Context, c bo.Search) (count int, data gdb.List, err error) {
	db := g.DB().Ctx(ctx).Model(c.T1 + " t1")
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
	if count, err = db.Count(); err != nil {
		return
	}
	var o = "t1.id desc"
	if c.OrderBy != "" {
		o = c.OrderBy
	}
	if c.SearchFields == "" {
		c.SearchFields = "t1.*"
	}
	all, err := db.Page(c.Page, c.Size).Fields(c.SearchFields).Order(o).All()
	if all.IsEmpty() {
		return
	}
	data = all.List()
	return
}
func Add(ctx context.Context, table, data interface{}) error {
	if _, err := g.DB().Ctx(ctx).Model(table).Insert(data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func Del(ctx context.Context, table, id interface{}) (err error) {
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}
func Update(ctx context.Context, table string, id, data interface{}) error {
	// 空值过滤
	_, err := g.DB().Model(table).Where("id", id).Update(data)
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
func GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	return one, nil
}

// NodeInfo 菜单信息
func NodeInfo(ctx context.Context, path string) (*entity.Menu, error) {
	m, err := dao.Menu.GetByPath(ctx, path)
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
func Icon(ctx context.Context, path string) (string, error) {
	menu, err := dao.Menu.GetByPath(ctx, path)
	if err != nil {
		return "", nil
	}
	icon := menu.Icon
	if err != nil {
		return "", nil
	}
	if icon == "" {
		return "", nil
	}
	if strings.HasPrefix(icon, "http") {
		return icon, nil
	}
	return consts.ImgPrefix + icon, err
}

func OperationLogClear(ctx context.Context) error {
	if _, err := dao.OperationLog.Ctx(ctx).Delete("id is not null"); err != nil {
		return err
	}
	return nil
}
func MsgFromSession(r *ghttp.Request) string {
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
