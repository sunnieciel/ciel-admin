package sys

import (
	"ciel-admin/internal/dao"
	"context"
	"github.com/gogf/gf/v2/text/gstr"
)

func DictGetByKey(ctx context.Context, key string) (string, error) {
	dict, err := dao.Dict.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return dict.V, nil
}
func DictApiGroup(ctx context.Context) ([]interface{}, error) {
	d, err := dao.Dict.GetByKey(ctx, "api_group")
	if err != nil {
		return nil, err
	}
	res := make([]interface{}, 0)
	for _, i := range gstr.Split(d.V, "\n") {
		if i != "" {
			all := gstr.TrimAll(i)
			res = append(res, map[string]interface{}{
				"label": all,
				"value": all,
			})
		}
	}
	return res, nil
}
