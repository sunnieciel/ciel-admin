package api

import "github.com/gogf/gf/v2/frame/g"

type Authorization struct {
	Token string `json:"Authorization" in:"header" `
}
type DefaultReq struct {
	Data   g.Map
	g.Meta `tags:"后台"`
}

type DefaultRes struct{}

type PageReq struct {
	Page int64 `json:"page" dc:"页面" d:"1" in:"query"`
	Size int64 `json:"size" dc:"页面大小" d:"15" in:"query"`
}
type PageRes struct {
	Page      int64 `json:"page" d:"1" dc:"当前页数"`
	TotalPage int64 `json:"total_page,omitempty" dc:"总页数"`
	Size      int64 `json:"size,omitempty" dc:"页面大小"`
	Total     int64 `json:"total" dc:"总条数"`
}
