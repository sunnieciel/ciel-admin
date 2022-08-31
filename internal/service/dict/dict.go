package dict

import (
	"ciel-admin/internal/consts"
	"ciel-admin/internal/dao"
	"ciel-admin/internal/service/view"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

func GetByKey(ctx context.Context, key string) (string, error) {
	dict, err := dao.Dict.GetByKey(ctx, key)
	if err != nil {
		return "", err
	}
	return dict.V, nil
}
func ApiGroup(ctx context.Context) (string, error) {
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
	return strings.Join(arr, ","), nil
}

func SetWhiteIps(ctx context.Context, v ...string) error {
	if len(v) == 0 {
		d, err := dao.Dict.GetByKey(ctx, "white_ips")
		if err != nil {
			return err
		}
		consts.WhiteIps = d.V
	} else {
		consts.WhiteIps = v[0]
	}
	return nil
}
