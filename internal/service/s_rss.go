package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gxml"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/glog"
	"time"
)

type rss struct{}

var sRss = newRss()

func newRss() *rss {
	return &rss{}
}
func Rss() *rss {
	return sRss
}

func (s *rss) fetchXml(ctx context.Context, url string) (map[string]interface{}, error) {
	num := 0
	max := 5
retry:
	data, err := gclient.New().Timeout(time.Second*3).Get(ctx, url)
	if err != nil {
		num++
		if num > max {
			return nil, errors.New(fmt.Sprintf("获取RSS数据失败,已重试%d次,请稍后重试", max))
		}
		glog.Infof(ctx, "获取RSS失败,重试中...%d", num)
		goto retry
	}
	return gxml.DecodeWithoutRoot([]byte(data.ReadAllString()))
}

func (s *rss) Feftch(ctx context.Context, url string) (map[string]interface{}, error) {
	return s.fetchXml(ctx, url)
}
