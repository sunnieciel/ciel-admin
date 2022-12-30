package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AllDictReq struct {
	g.Meta `tags:"系统" dc:"查询所有字典(每个字段说明请具体访问接口)"`
}
type DictReq struct {
	g.Meta `tags:"系统" dc:"查询单个字典信息"`
	Key    string `json:"key" in:"query"`
}
type BannersReq struct {
	g.Meta `tags:"系统" dc:"查询Banner图列表"`
}
type UploadFileReq struct {
	g.Meta `tags:"系统" dc:"上传图片"`
	Group  int `json:"group" dc:"分组:1头像,2图片,3动图,4音频,5文件" d:"2" in:"query"`
}
type UploadFileRes struct {
	DbName    string `json:"db_name" dc:"图片"`
	ImgPrefix string `json:"img_prefix" dc:"图片前缀"`
}
type BannerRes struct {
	Id        int        `json:"id" dc:"id"`
	Title     string     `json:"title" dc:"标题"`
	Image     string     `json:"image" dc:"图片地址"`
	Link      string     `json:"link" dc:"跳转链接"`
	CreatedAt gtime.Time `json:"created_at" dc:"创建时间"`
}
type DictRes struct {
	Value string `json:"value" dc:"值"`
}
type AllDictRes g.Map
