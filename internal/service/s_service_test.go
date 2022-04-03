package service

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

func TestFetch(t *testing.T) {
	data, err := Rss().fetchXml(nil, "https://rsshub.app/douban/book/rank/nonfiction")
	if err != nil {
		panic(err)
	}
	g.Dump(data)
}
func TestGen_GetTableInfo(t *testing.T) {
	info, err := Gen().Fields(nil, "s_dict")
	if err != nil {
		panic(err)
	}
	g.Dump(info)
}
func TestListTables(t *testing.T) {
	ctx := context.TODO()
	tables, err := Gen().Tables(ctx)
	if err != nil {
		panic(err)
	}
	g.Dump(tables)
}
func TestGetProjectName(t *testing.T) {
	g.Dump(fmt.Sprintf("%s/internal/controller", gfile.MainPkgPath()))
}
func TestGen_GenController(t *testing.T) {
}
