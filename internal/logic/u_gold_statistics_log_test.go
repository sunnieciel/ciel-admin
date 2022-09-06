package logic

import (
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_lGoldStatisticsLog_GoldReport(t *testing.T) {
	g.Dump(GoldStatisticsLog.GoldReport(nil, "", ""))
}
