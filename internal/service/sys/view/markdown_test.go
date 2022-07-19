package view

import (
	"github.com/gogf/gf/v2/os/glog"
	"testing"
)

func TestMarkdown(t *testing.T) {
	md := []byte(`
# hello
<div class='tag-info'>hello</div>

    golang

- ok
- ok2
`)
	html := MD(md)
	glog.Info(nil, html)
}
