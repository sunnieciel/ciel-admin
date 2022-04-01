package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestFetch(t *testing.T) {
	data, err := Rss().fetchXml(nil, "https://rsshub.app/douban/book/rank/nonfiction")
	if err != nil {
		panic(err)
	}
	g.Dump(data)
}
