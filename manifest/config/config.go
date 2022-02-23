package config

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Field struct {
	Field      string
	QueryFiled string
	Title      string
	Desc       string  // 描述
	Type       string  // text,select,number,showImg
	Search     bool    // 是否搜索
	Like       bool    // 是否模糊搜索
	Disabled   bool    // 是否禁用
	EditHidden bool    // 编辑是否隐藏
	Hidden     bool    // 展示是否隐藏
	Required   bool    // 是否必填
	Step       float64 // 当Type为number时，step为步长
	Items      []*Item
	Value      interface{} // 条件查询时从 request中获取的值
	ShowImg    *ShowImg    // type showImg
}
type ShowImg struct {
	ImgPrefix string
	Field     string
}
type Item struct {
	Text    string
	Value   interface{}
	BgColor string
}

type SearchConf struct {
	PageUrl      string
	PageTitle    string
	UrlPrefix    string
	NoEdit       bool
	NoDel        bool
	NoSearchShow bool
	T1           string
	T2           string
	T3           string
	T4           string
	T5           string
	Page, Size   int
	OrderBy      string
	FieldsEx     string
	SearchFields string
	Fields       []*Field // 条件查询的字段
}

// FilterConditions 过滤需要查询的字段
func (s *SearchConf) FilterConditions(ctx context.Context) []*Field {
	request := g.RequestFromCtx(ctx)
	data := make([]*Field, 0)
	for _, field := range s.Fields {
		if field.Search {
			if field.QueryFiled == "" {
				field.QueryFiled = field.Field
			}
			field.Value = request.GetQuery(field.QueryFiled)
			data = append(data, field)
		}
	}
	return data
}
