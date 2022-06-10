package bo

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/text/gstr"
	"strings"
)

type GenConf struct {
	StructName string `v:"required#结构体名称不能为空"`
	HtmlGroup  string `v:"required#html分组不能为空"`
	PageName   string `v:"required#页面名称不能为空"`
	PageDesc   string `v:"required#页面描述不能为空"`
	AddBtn     int
	UpdateBtn  int
	DelBtn     int
	UrlPrefix  string

	T1, T2, T3, T4, T5, T6 string
	OrderBy                string
	QueryField             string
	Fields                 []*GenFiled

	MenuLevel1 string `v:"required#菜单一级不能为空"`
	MenuLevel2 string `v:"required#菜单名不能为空"`
	MenuLogo   string
}

func (s *GenConf) SetUrlPrefix() error {
	if s.T1 == "" {
		return errors.New("表名称不能为空")
	}
	d := strings.Split(s.T1, "_")[1]
	s.UrlPrefix = fmt.Sprint("/", gstr.CaseCamelLower(d), "/")
	return nil
}

type GenFiled struct {
	*gdb.TableField
	Label     string //  label is empty, use name
	FieldType string // select number text date datetime
	EditHide  int    // 1 true
	NotShow   int    // 1 true  not show in table
	Comment   string // is comment is not empty ,add el-tag comment
	Options   []*FieldOption

	SearchType int // 0 no,1 = ,2 like,3 >, 4 <, 5>=,6 <=,7 !=
	QueryName  string
}
type FieldOption struct {
	Value interface{}
	Label string
	Type  string // primary info success warning danger
}
