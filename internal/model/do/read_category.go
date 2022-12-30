// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReadCategory is the golang structure of table r_read_category for DAO operations like Where/Data.
type ReadCategory struct {
	g.Meta      `orm:"table:r_read_category, do:true"`
	Id          interface{} //
	Name        interface{} //
	Desc        interface{} //
	Speak       interface{} //
	Icon        interface{} //
	Poster      interface{} //
	TotalNum    interface{} // 文章数量
	Sort        interface{} //
	SubCategory interface{} //
	CreatedAt   *gtime.Time //
}
