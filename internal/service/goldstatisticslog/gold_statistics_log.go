package goldstatisticslog

import (
	"ciel-admin/internal/logic"
	"context"
)

func Clear(ctx context.Context) error { return logic.GoldStatisticsLog.Clear(ctx) }
