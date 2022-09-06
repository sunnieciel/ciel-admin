package goldchangelog

import (
	"ciel-admin/internal/logic"
	"context"
)

func Clear(ctx context.Context) error {
	return logic.GoldChangeLog.Clear(ctx)
}
