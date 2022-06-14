package sys

import (
	"ciel-admin/internal/service/internal/dao"
	"context"
)

func DictGetByKey(ctx context.Context, key string) (string, error) {
	dict, err := dao.Dict.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return dict.V, nil
}
