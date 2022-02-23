package dao

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"testing"
)

func TestMenus(t *testing.T) {
	ctx := gctx.New()
	menus, err := RoleMenu.Menus(ctx, 1, -1)
	if err != nil {
		panic(err)
	}
	glog.Info(ctx, menus)
}
