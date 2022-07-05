package node

import (
	"ciel-admin/internal/model/entity"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

func NodeDesc(data interface{}) string {
	json := gjson.New(data)
	var d entity.Node
	_ = json.Scan(&d)
	year := year(d.Year)
	month := month(d.Month)
	day := day(d.Day)
	level := level(d.Level)
	g.Log().Notice(nil, d)
	return fmt.Sprint(year, month, day, level, d.MainThings)
}

func level(level int) interface{} {
	switch level {
	case 1:
		return "<span class='tag-info'>普通</span>"
	case 2:
		return "<span class='tag-success'>稀有</span>"
	case 3:
		return "<span class='tag-primary'>传承</span>"
	case 4:
		return "<span class='tag-warning'>唯一</span>"
	case 5:
		return "<span class='tag-danger'>史诗</span>"
	default:
		return ""
	}
}

func day(day int) interface{} {
	if day != 0 {
		return fmt.Sprint(day, "日")
	}
	return ""
}

func month(month int) interface{} {
	if month != 0 {
		return fmt.Sprint(month, "月")
	}
	return ""
}

func year(year string) interface{} {
	if year != "" {
		return fmt.Sprint(year, "年")
	}
	return ""
}
