package config

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Field struct {
	Field      string
	QueryField string
	Value      interface{}
	Like       bool // 是否模糊搜索
	GT         bool // >
	GTE        bool // >=
	LT         bool // <
	LTE        bool // <=
	In         bool // in
}

type SearchConf struct {
	PageUrl      string
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
		if field.QueryField == "" {
			field.QueryField = field.Field
		}
		field.Value = request.GetQuery(field.QueryField)
		data = append(data, field)
	}
	return data
}
