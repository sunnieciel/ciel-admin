package xcmd

import (
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

func MustScan(format string) *gvar.Var {
	d := gcmd.Scan(format)
	for {
		if d == "" {
			d = gcmd.Scan("值不能为空！请重新输入:")
		} else {
			return gvar.New(d, true)
		}
	}
}
func MustChoose(msg string, data []*gvar.Var) gdb.Value {
	fmt.Printf("---------%s-------\n", msg)
	for i, datum := range data {
		fmt.Printf("%d %v\n", i, datum)
	}
	println("--------------------")
	for {
		d := gcmd.Scan("如果没有您的选项请自行输入:")
		if d == "" {
			d = gcmd.Scanf("值不能为空！请重新输入:")
		} else {
			num := gconv.Int(d)
			if gstr.IsNumeric(d) && num <= len(data) && num >= 0 { // 如果是数字则为选择
				return data[num]
			}
			return gvar.New(d)
		}
	}
}
