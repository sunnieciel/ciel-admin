// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Node is the golang structure of table f_node for DAO operations like Where/Data.
type Node struct {
	g.Meta     `orm:"table:f_node, do:true"`
	Id         interface{} //
	Year       interface{} //
	Month      interface{} //
	Day        interface{} //
	Uid        interface{} //
	Level      interface{} //
	Tag        interface{} //
	MainThings interface{} //
	OtherInfo  interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
