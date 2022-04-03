package bo

type GenCodeInfo struct {
	Table                  string `v:"required#表名不能为空"`
	StructName             string `v:"required#结构体名称"`
	Category               string `v:"required#请输入类别"`
	Desc                   string `v:"required#请输入描述"`
	Title                  string `v:"required#请输入标题"`
	T1, T2, T3, T4, T5, T6 string
	Fields                 []*Field
	OrderBy                string
	QueryField             string
}
type Field struct {
	Name        string
	Comment     string
	Type        string
	SearchType  string
	QueryField  string
	Title       string
	Sort        int
	DetailsType string // show no-show disabled
}
