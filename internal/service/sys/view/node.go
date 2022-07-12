package view

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

func nodeWeek(y, m, d interface{}) string {
	if y != nil && m != nil && d != nil {
		t := gtime.NewFromStr(fmt.Sprintf("%v-%v-%v", y, m, d))
		format := t.Format("w")
		week := format
		if week == "0" {
			week = "天"
		}
		switch gconv.Int(t.Format("w")) {
		case 0:
			d = "日"
		case 1:
			d = "一"
		case 2:
			d = "二"
		case 3:
			d = "三"
		case 4:
			d = "四"
		case 5:
			d = "五"
		case 6:
			d = "六"
		default:
		}
		return tag(format, fmt.Sprintf("星期%s", d))
	}
	return ""
}
func nodeTime(y, m, d interface{}) string {
	res := tag(y, fmt.Sprint(y, "年"))
	if m != nil {
		res += tag(m, fmt.Sprint(m, "月"))
	}
	if d != nil {
		res += tag(1, fmt.Sprint(d, "日"))
	}
	return res
}
func nodeLevel(i interface{}) string {
	content := ""
	i2 := gconv.Int(i)
	switch i2 {
	case 1:
		content = "普通"
	case 2:
		content = "稀有"
	case 3:
		content = "传承"
	case 4:
		content = "唯一"
	case 5:
		content = "史诗"
	case 6:
		content = "传说"
	case 7:
		content = "神话"
	default:
		content = "普通"
	}
	return tag(i, content)
}
