package service

import (
	v1 "ciel-admin/api/v1"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/jaytaylor/html2text"
	"math"
)

var (
	Common = sCommon{}
)

type sCommon struct{}

func (s sCommon) MakeDesc(desc string) (string, error) {
	fromString, err := html2text.FromString(desc)
	if err != nil {
		return "", err
	}
	l := 33
	if len(fromString) < 33 {
		l = len(fromString)
	}
	return gstr.SubStrRune(fromString, 0, l) + "...", nil
}

func (s sCommon) GetPageInfo(page, size int, total int64) *v1.PageRes {
	if size == 0 {
		size = 10
	}
	totalPage := math.Ceil(float64(total) / float64(size)) //这里计算总页数时，要向上取整
	if totalPage <= 0 {
		totalPage = 1
	}
	return &v1.PageRes{
		TotalPage: int64(totalPage),
		Total:     total,
		Page:      int64(page),
		Size:      int64(size),
	}
}
