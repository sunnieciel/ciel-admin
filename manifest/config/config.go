package config

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Field struct {
	Field      string
	QueryFiled string
	Value      interface{}
	Like       bool // 是否模糊搜索
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
		if field.QueryFiled == "" {
			field.QueryFiled = field.Field
		}
		field.Value = request.GetQuery(field.QueryFiled)
		data = append(data, field)
	}
	return data
}
