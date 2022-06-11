package sys

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/xstr"
	"ciel-admin/utility/utils/xtime"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

func List(ctx context.Context, c *config.Search) (count int, data gdb.List, err error) {
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
			switch item.SearchType {
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
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func Del(ctx context.Context, table, id interface{}) (err error) {
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}
func Update(ctx context.Context, table string, id, data interface{}) error {
	// 空值过滤
	_, err := g.DB().Model(table).Where("id", id).Data(data).Update()
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		glog.Error(ctx, err)
		return nil, err
	}
	return one, nil
}
func Icon(ctx context.Context, path string) (string, error) {
	menu, err := dao.Menu.GetByPath(ctx, path)
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
func Init() {
	get, err := g.Cfg().Get(gctx.New(), "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
}

func MenusLevel1(ctx context.Context) ([]gdb.Value, error) {
	return dao.Menu.Ctx(ctx).Array("name", "pid=-1")
}
