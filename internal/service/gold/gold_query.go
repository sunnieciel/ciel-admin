package gold

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/model/entity"
	"context"
)

func GetById(ctx context.Context, uid uint64) (*entity.Gold, error) {
	return dao.Gold.GetByUid(ctx, uid)
}
