package view

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gomarkdown/markdown"
)

func MD(data interface{}) string {
	return string(markdown.ToHTML([]byte(gconv.String(data)), nil, nil))
}
