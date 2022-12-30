package v1

type Authorization struct {
	Token string `json:"Authorization" in:"header" `
}

type DefaultRes struct{}

type PageReq struct {
	Page int `json:"page" dc:"页面" d:"1" in:"query"`
	Size int `json:"size" dc:"页面大小" d:"10" in:"query"`
}
type PageRes struct {
	Page      int64 `json:"page,omitempty" dc:"当前页数"`
	TotalPage int64 `json:"total_page,omitempty" dc:"总页数"`
	Size      int64 `json:"size,omitempty" dc:"页面大小"`
	Total     int64 `json:"total" dc:"总条数"`
}
