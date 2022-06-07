package bo

import "github.com/gogf/gf/v2/database/gdb"

type GenConf struct {
	Table      string
	StructName string `v:"required#结构体名称"`
	HtmlGroup  string
	PageName   string
	PageDesc   string
	AddBtn     int
	UpdateBtn  int
	DelBtn     int
	UrlPrefix  string

	T1, T2, T3, T4, T5, T6 string
	OrderBy                string
	QueryField             string
	Fields                 []*GenFiled
}
type GenFiled struct {
	*gdb.TableField
	Label     string //  label is empty, use name
	FieldType string // select number text date datetime
	EditHide  int    // 1 true
	NotShow   int    // 1 true  not show in table
	Comment   string // is comment is not empty ,add el-tag comment
	Options   []*FieldOption

	SelectType int // 0 no,1 = ,2 like,3 >, 4 <, 5>=,6 <=,7 !=
	QueryName  string
}
type FieldOption struct {
	Value interface{}
	Label string
	Type  string // primary info success warning danger
}
