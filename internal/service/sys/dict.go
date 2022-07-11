package sys

import (
	"ciel-admin/internal/dao"
	"ciel-admin/internal/service/sys/view"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

func DictGetByKey(ctx context.Context, key string) (string, error) {
	dict, err := dao.Dict.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return dict.V, nil
}
func DictApiGroup(ctx context.Context) (string, error) {
	d, err := dao.Dict.GetByKey(ctx, "api_group")
	if err != nil {
		return "", err
	}
	arr := make([]string, 0)
	for index, i := range gstr.Split(d.V, "\n") {
		if i != "" {
			i = gstr.TrimAll(i)
			arr = append(arr, fmt.Sprintf("%s:%s:%s", i, i, view.SwitchTagClass(index)))
		}
	}
	join := strings.Join(arr, ",")
	g.Log().Warning(nil, join)
	return join, nil
}
