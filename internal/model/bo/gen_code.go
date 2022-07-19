package bo

import (
	"github.com/gogf/gf/v2/database/gdb"
)

type GenConf struct {
	GenType    int // 生成页面类型  0 curd 1静态页面
	StructName string
	HtmlGroup  string `v:"required#html分组不能为空"`
	PageName   string `v:"required#页面名称不能为空"`
	PageDesc   string `v:"required#页面描述不能为空"`

	MenuLevel1 string `v:"required#菜单一级不能为空"`
	MenuLevel2 string `v:"required#菜单名不能为空"`
	MenuLogo   string
	AddBtn     int
	UpdateBtn  int
	DelBtn     int
	ShowStatus int // 展示状态  0是 1否

	T1, T2, T3, T4, T5, T6 string
	OrderBy                string
	QueryField             string
	Fields                 []*GenFiled
}

type GenFiled struct {
	*gdb.TableField
	Label      string //  字段的名称 为空将使用默认字段
	FieldType  string // 字段的html类型 text,textarea,markdown,select,number,date,datetime
	SearchType int    // 字段的搜索类型 0 不查询(默认),1 = ,2 like,3 >, 4 <, 5>=,6 <=,7 !=
	EditHide   int    // 编辑时是否隐藏 1 隐藏,默认不隐藏
	AddHide    int    // 添加时是否隐藏 1 隐藏,默认不隐藏
	Hide       int    // 列表该字段是否隐藏 1隐藏,默认不隐藏
	Disabled   int    // 编辑时是是否禁用  1是的,默认否
	Required   int    // 添加是是否必须  1是,默认否
	Comment    string // 字段说明描述，不为空则 在添加和编辑时会展示该描述
	Options    string // 当 FieldType为例 `select`时 进行的选项 格式 "值:名称:类名,值:名称:类名"  eg "1:文本:tag-info,2:图片:tag-warning,3:富文本:tag-primary,4:文件:tag-danger,5:其他:tag-purple"
	QueryName  string // 不用设置，生成时会如果有额外的关联查询程序会提示进行填写
}
