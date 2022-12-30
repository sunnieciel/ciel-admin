package model

import (
	"ciel-admin/internal/model/entity"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Admin struct {
	Admin *entity.Admin
	Menus []*Menu
}
type Menu struct {
	Id       int     `json:"id"        description:""`
	Pid      int     `json:"pid"       description:""`
	Name     string  `json:"name"      description:""`
	Path     string  `json:"path"      description:""`
	Type     int     `json:"type"      description:""`
	Sort     float64 `json:"sort"      description:""`
	Status   int     `json:"status"    description:""`
	Children []*Menu
}

type Search struct {
	T1           string
	T2           string
	T3           string
	T4           string
	T5           string
	T6           string
	Page, Size   int
	OrderBy      string
	SearchFields string
	Fields       []Field // 条件查询的字段
	Begin        string  // 查询时间的开始
	End          string  // 查询时间的结束
}

type Field struct {
	Name      string
	QueryName string
	Type      int // 0 no,1 = ,2 like,3 >, 4 <, 5>=,6 <=,7 != ,8 date,9 date begin
	Value     interface{}
}

// FilterConditions 过滤需要查询的字段
func (s *Search) FilterConditions(ctx context.Context) []Field {
	request := g.RequestFromCtx(ctx)
	data := make([]Field, 0)
	for _, field := range s.Fields {
		if field.QueryName == "" {
			field.QueryName = field.Name
		}
		query := request.GetQuery(field.QueryName)
		if !query.IsEmpty() {
			field.Value = query
		}
		data = append(data, field)
	}
	return data
}
