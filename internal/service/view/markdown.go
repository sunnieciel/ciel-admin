package view

import (
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gomarkdown/markdown"
)

func MD(data string) string {
	data = gstr.Replace(data, "\r", "")
	return string(markdown.ToHTML([]byte(data), nil, nil))
}
