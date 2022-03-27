package oop

import (
	"ciel-admin/manifest/config"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

type ISystem interface {
	List(ctx context.Context, c *config.SearchConf) (count int, data gdb.List, err error)
	Add(ctx context.Context, table, data interface{}) error
	Update(ctx context.Context, table string, id, data interface{}) error
	Del(ctx context.Context, table, id interface{}) (err error)
	GetById(ctx context.Context, table, id interface{}) (gdb.Record, error)
	Init()
	GetMenuIcon(ctx context.Context, path string) (string, error)
}
