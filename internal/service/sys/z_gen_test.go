package sys

import (
	"ciel-admin/internal/model/bo"
	"ciel-admin/utility/utils/xpwd"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/olekukonko/tablewriter"
	"math"
	"strings"
	"testing"
)

func TestFields(t *testing.T) {
	fields, err := Fields(nil, "s_api")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(fields)
}
func TestGen(t *testing.T) {
	c := &bo.GenConf{}
	c.HtmlGroup = "sys"
	c.PageName = "Admin"
	c.T1 = "s_admin"
	c.UrlPrefix = "/admin/"
	c.Fields = make([]*bo.GenFiled, 0)
	fields, err := Fields(nil, c.T1)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range fields {
		f := bo.GenFiled{TableField: field}
		c.Fields = append(c.Fields, &f)
		if field.Name == "group" {
			f.SearchType = 1
		}
		if field.Name == "status" {
			f.SearchType = 1
			f.FieldType = "select"
		}
	}
	err = genHtml(nil, c)
	if err != nil {
		t.Fatal(err)
	}
}
func TestString(t *testing.T) {
	//data := [][]string{
	//	[]string{"A", "The Good", "500"},
	//	[]string{"B", "The Very very Bad Man", "288"},
	//	[]string{"C", "The Ugly", "120"},
	//	[]string{"D", "The Gopher", "800"},
	//}
	s := &strings.Builder{}
	table := tablewriter.NewWriter(s)
	table.SetHeader([]string{"步骤", "名称", "用时"})
	table.Render()
	fmt.Println(s.String())
	//for _, v := range data {
	//	table.Append(v)
	//	table.Render()
	//	fmt.Println(s.String())
	//}
}

func TestNum(t *testing.T) {
	glog.Debug(nil, math.Ceil(1.0))
}

func TestPwd(t *testing.T) {
	fmt.Println(xpwd.GenPwd("1"))
}
