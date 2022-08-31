package dict

import (
	"ciel-admin/internal/logic"
	"context"
)

func GetByKey(ctx context.Context, key string) (string, error) { return logic.Dict.GetByKey(ctx, key) }
func ApiGroup(ctx context.Context) (string, error)             { return logic.Dict.ApiGroup(ctx) }
func SetWhiteIps(ctx context.Context, v ...string) error       { return logic.Dict.SetWhiteIps(ctx, v...) }
