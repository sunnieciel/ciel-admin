package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// switch environment 0 dev, 1 server
func Test_sSystem_SwitchEnvironment(t *testing.T) {
	g.Dump(Sys.SwitchEnvironment("/Users/ciel/work/9.brushing", 0))
}
