package service

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/oop"
	"ciel-admin/internal/service/internal/dao"
	"ciel-admin/manifest/config"
	"ciel-admin/utility/utils/xstr"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

//  ---sSystem ------------------------------------------------------------
type sSystem struct{}

func (s *sSystem) Icon(ctx context.Context, path string) (string, error) {
	menu, err := dao.Menu.GetByPath(ctx, path)
	if err != nil {
		return "", err
	}
	if menu.Icon == "" {
		return "", err
	}
	return consts.ImgPrefix + menu.Icon, err
}

var insSystem = newSystem()

func newSystem() *sSystem {
	return &sSystem{}
}

func System() oop.ISystem { return insSystem }
func (s *sSystem) List(ctx context.Context, c *config.SearchConf) (count int, data gdb.List, err error) {
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

	conditions := c.FilterConditions(ctx)
	if len(conditions) > 0 {
		for _, item := range conditions {
			field := item.Field
			if g.IsEmpty(item.Value) {
				continue
			}
			if !strings.Contains(field, ".") {
				field = "t1." + field
			}
			if item.Like {
				db = db.WhereLike(field, xstr.Like(gconv.String(item.Value)))
			} else {
				db = db.Where(field, item.Value)
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
func (s *sSystem) Add(ctx context.Context, table, data interface{}) error {
	_, err := g.DB().Ctx(ctx).Model(table).Insert(data)
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func (s *sSystem) Del(ctx context.Context, table, id interface{}) (err error) {
	if _, err = g.DB().Ctx(ctx).Model(table).Delete("id", id); err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}
func (s *sSystem) Update(ctx context.Context, table string, id, data interface{}) error {
	// 空值过滤
	_, err := g.DB().Model(table).Where("id", id).Data(data).Update()
	if err != nil {
		glog.Error(ctx, err)
		return err
	}
	return nil
}
func (s *sSystem) GetById(ctx context.Context, table, id interface{}) (gdb.Record, error) {
	one, err := g.DB().Ctx(ctx).Model(table).One("id", id)
	if err != nil {
		glog.Error(ctx, err)
		return nil, err
	}
	return one, nil
}
func (s *sSystem) Init() {
	get, err := g.Cfg().Get(gctx.New(), "server.imgPrefix")
	if err != nil {
		panic(err)
	}
	consts.ImgPrefix = get.String()
}
