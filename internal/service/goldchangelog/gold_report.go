package goldchangelog

import (
	"ciel-admin/internal/logic"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
)

func GoldReport(ctx context.Context, begin, end string) (gdb.Record, error) {
	return logic.GoldStatisticsLog.GoldReport(ctx, begin, end)
}
