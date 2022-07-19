package sys

import (
	"ciel-admin/internal/model/bo"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"testing"
)

func TestParseTableFields(t *testing.T) {
	conf := bo.GenConf{T1: "s_admin_login_log"}
	parseTableFields(gctx.New(), &conf)
	glog.Info(nil, conf)
}
