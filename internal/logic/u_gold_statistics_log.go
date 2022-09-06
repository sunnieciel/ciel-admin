package logic

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"math"
	"time"
)

var (
	GoldStatisticsLog = lGoldStatisticsLog{}
)

type lGoldStatisticsLog struct {
}

func (l lGoldStatisticsLog) Add(ctx context.Context, tx *gdb.TX, t int, uid uint64, amount float64) error {
	todayLog, err := dao.GoldStatisticsLog.GetTodayLog(ctx, tx, uid)
	if err != nil {
		if err != consts.ErrDataNotFound {
			return err
		}
		data := g.Map{
			"uid":                 uid,
			"created_date":        time.Now(),
			fmt.Sprintf("t%d", t): math.Abs(amount),
		}
		if _, err = tx.Model(dao.GoldStatisticsLog.Table()).Insert(data); err != nil {
			return err
		}
		return nil
	}
	if _, err = tx.Model(dao.GoldStatisticsLog.Table()).
		WherePri(todayLog.Id).
		Increment(fmt.Sprintf("t%d", t), math.Abs(amount)); err != nil {
		return err
	}
	return nil
}

func (l lGoldStatisticsLog) Clear(ctx context.Context) error {
	_, err := dao.GoldStatisticsLog.Ctx(ctx).Delete("id is not null")
	if err != nil {
		g.Log().Error(ctx, err)
		return err
	}
	return nil
}

func (l lGoldStatisticsLog) GoldReport(ctx context.Context, begin string, end string) (gdb.Record, error) {
	if begin == "" {
		begin = gtime.Now().AddDate(0, -6, 0).StartOfDay().String()
	}
	db := dao.GoldStatisticsLog.Ctx(ctx).
		FieldSum("t1", "t1").
		FieldSum("t2", "t2").
		FieldSum("t3", "t3").
		FieldSum("t4", "t4").
		FieldSum("t5", "t5").
		WhereGTE("created_date", begin)
	if end != "" {
		db = db.WhereLTE("created_date", end)
	}
	one, err := db.One()
	if err != nil {
		return nil, err
	}
	return one, nil
}
