package gen

import (
	"ciel-admin/internal/logic"
	"context"
)

func MenuLeve1(ctx context.Context) (string, error) {
	return logic.Gen.MenuLeve1(ctx)
}
func Tables(ctx context.Context) (string, error) {
	return logic.Gen.Tables(ctx)
}
func Gen(ctx context.Context, table string, group string, menu string, prefix string, apiGroup, htmlGroup string) error {
	return logic.Gen.Gen(ctx, table, group, menu, prefix, apiGroup, htmlGroup)
}
