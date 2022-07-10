package sys

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestDictGroup(t *testing.T) {
	group, err := DictApiGroup(nil)
	if err != nil {
		t.Fatal(err)
	}
	g.Log().Info(nil, group)
}
