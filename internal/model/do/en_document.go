// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// EnDocument is the golang structure of table e_en_document for DAO operations like Where/Data.
type EnDocument struct {
	g.Meta       `orm:"table:e_en_document, do:true"`
	Id           interface{} //
	Category1    interface{} //
	Category2    interface{} //
	Title        interface{} //
	En           interface{} //
	Desc         interface{} //
	Icon         interface{} //
	Poster       interface{} //
	WordNum      interface{} // 单词个数
	ParagraphNum interface{} // 段落个数
	Sort         interface{} //
	Status       interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
