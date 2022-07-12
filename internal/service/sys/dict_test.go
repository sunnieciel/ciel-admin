package sys

import (
	"fmt"
	"github.com/kayon/iploc"
	"testing"
)

func TestIP(t *testing.T) {
	loc, err := iploc.Open("qqwry.dat")
	if err != nil {
		panic(err)
	}
	detail := loc.Find("8.8.8") // 补全为8.8.0.8, 参考 ping 工具
	fmt.Printf("IP:%s; 网段:%s - %s; %s\n", detail.IP, detail.Start, detail.End, detail)

	detail2 := loc.Find("8.8.3.1")
	fmt.Printf("%t %t\n", detail.In(detail2.IP.String()), detail.String() == detail2.String())
	// output
	// IP:8.8.0.8; 网段: 8.7.245.0 - 8.8.3.255; 美国 科罗拉多州布隆菲尔德市Level 3通信股份有限公司
	// true true
	detail = loc.Find("1.24.41.0")
	fmt.Println(detail.String())
	fmt.Println(detail.Country, detail.Province, detail.City, detail.County)
}
