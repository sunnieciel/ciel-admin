package logic

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	GoldChangeLog = lGoldChangeLog{}
)

type lGoldChangeLog struct {
}

func (l lGoldChangeLog) Add(ctx context.Context, tx *gdb.TX, transId string, t int, uid uint64, amount float64, balance float64, desc string) error {
	var (
		data = do.GoldChangeLog{
			TransId: transId,
			Uid:     uid,
			Type:    t,
			Amount:  amount,
			Balance: balance,
			Desc:    desc,
		}
	)
	if _, err := tx.Model(dao.GoldChangeLog.Table()).Insert(data); err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}
