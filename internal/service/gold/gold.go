package gold

import (
	"ciel-admin/internal/logic"
	"context"
)

func UpdatePassByAdmin(ctx context.Context, pass string, uid uint64) error {
	return logic.Gold.UpdatePassByAdmin(ctx, pass, uid)
}

func TopUpByAdmin(ctx context.Context, t int, uid uint64, amount float64, desc string) error {
	return logic.Gold.TopUpByAdmin(ctx, t, uid, amount, desc)
}
func DeductByAdmin(ctx context.Context, t int, uid uint64, amount float64) error {
	return logic.Gold.DeductByAdmin(ctx, t, uid, amount)

}
