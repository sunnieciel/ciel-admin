package sys

import (
	"ciel-admin/internal/model/bo"
	"github.com/gogf/gf/v2/frame/g"
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
	c.Table = "s_admin"
	c.UrlPrefix = "/admin/"
	c.Fields = make([]*bo.GenFiled, 0)
	fields, err := Fields(nil, c.Table)
	if err != nil {
		t.Fatal(err)
	}
	for _, field := range fields {
		f := bo.GenFiled{TableField: field}
		c.Fields = append(c.Fields, &f)
		if field.Name == "group" {
			f.SelectType = 1
		}
		if field.Name == "status" {
			f.SelectType = 1
			f.FieldType = "select"
		}
	}
	err = genHtml(nil, c)
	if err != nil {
		t.Fatal(err)
	}
}
