package menu

import (
	"ciel-admin/internal/logic"
	"context"
)

func SetGroupSort(ctx context.Context, sort int, id uint64) error {
	return logic.Menu.SetGroupSort(ctx, sort, id)
}
