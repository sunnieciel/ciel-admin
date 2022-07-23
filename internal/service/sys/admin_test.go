package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestGetAllAdminOptions(t *testing.T) {
	options, err := GetAllAdminOptions(nil)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(options)
}
