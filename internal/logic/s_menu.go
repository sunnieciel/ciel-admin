package logic

import (
	"ciel-admin/internal/dao"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
)

var Menu = lMenu{}

type lMenu struct {
}

// SetGroupSort
// 1.0  换成 2.0
func (m lMenu) SetGroupSort(ctx context.Context, sort int, id uint64) error {
	change := func(in float64) float64 {
		arr := strings.Split(fmt.Sprintf("%.2f", in), ".")
		resStr := fmt.Sprintf("%d.%s", sort, arr[1])
		return gconv.Float64(resStr)
	}
	pMenu, err := dao.Menu.GetById(ctx, id)
	if err != nil {
		return err
	}
	pMenu.Sort = change(pMenu.Sort)
	if _, err = dao.Menu.Ctx(ctx).Save(pMenu); err != nil {
		return err
	}
	arr, err := dao.Menu.ListByPid(ctx, pMenu.Id)
	if err != nil {
		return err
	}
	for _, i := range arr {
		i.Sort = change(i.Sort)
		if _, err = dao.Menu.Ctx(ctx).Save(i); err != nil {
			return err
		}
	}
	return nil
}
