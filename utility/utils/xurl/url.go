package xurl

import (
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
)

func ToUrlParams(value interface{}) string {
	m := gconv.Map(value)
	res := ""
	for k, v := range m {
		res += fmt.Sprintf("%v=%v&", k, v)
	}
	return res
}
