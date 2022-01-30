// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-01-30 15:53:54
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Dict is the golang structure of table s_dict for DAO operations like Where/Data.
type Dict struct {
	g.Meta    `orm:"table:s_dict, do:true"`
	Id        interface{} //
	K         interface{} //
	V         interface{} //
	Desc      interface{} //
	Group     interface{} //
	Status    interface{} //
	Type      interface{} // 0 文本，1 img
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
