package view

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func Balance(in interface{}) string {
	money := gconv.Float64(in)
	var (
		bronze   = `<img class="mlr-3" height="16" border="0" src="/resource/image/bronze@2x.png" alt="">`
		silver   = `<img class="mlr-3" height="16" border="0" src="/resource/image/silver@2x.png" alt="">`
		gold     = `<img class="mlr-3" height="16" border="0" src="/resource/image/gold@2x.png" alt="">`
		b, si, g string
		s        = fmt.Sprintf("%.f", money)
	)
	if money == 0 {
		return fmt.Sprint(0, bronze)
	}
	switch {
	case len(s) > 4:
		g = s[:len(s)-4]
		si = s[len(s)-4 : len(s)-2]
		b = s[len(s)-2:]
	case len(s) > 2:
		si = s[:len(s)-2]
		b = s[len(s)-2:]
	default:
		b = s
	}
	res := ""
	if g != "" && g != "-" {
		res += fmt.Sprint(g, gold)
	}
	if si != "" && si != "00" && si != "-" {
		if gstr.HasPrefix(si, "0") {
			si = si[1:]
		}
		res += fmt.Sprint(si, silver)
	}
	if b != "" && b != "00" {
		if gstr.HasPrefix(b, "0") {
			b = b[1:]
		}
		res += fmt.Sprint(b, bronze)
	}
	return res
}
