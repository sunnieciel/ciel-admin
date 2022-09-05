package dao

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_goldStatisticsLogDao_GetTodayLog(t *testing.T) {
	g.DB().Transaction(nil, func(ctx context.Context, tx *gdb.TX) error {
		log, err := GoldStatisticsLog.GetTodayLog(ctx, tx, 4)
		if err != nil {
			t.Fatal(err)
		}
		g.Dump(log)
		return nil
	})
}
