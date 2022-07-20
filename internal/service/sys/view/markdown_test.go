package view

import (
	"github.com/gogf/gf/v2/os/glog"
	"testing"
)

func TestMarkdown(t *testing.T) {
	str := `
Name    | Age
--------|-----:
Bob     | 27
Alice   | 23`

	glog.Info(nil, MD(str))
}
