package logic

import (
	"fmt"
	"testing"
)

func TestCheckOrSave(t *testing.T) {
	err := checkGroupOrSave(nil, "haha")
	if err != nil {
		t.Fatal(err)
	}
}

// {name:123,nae:313}
func TestMakeToJsonStr(t *testing.T) {
	str := `{"label":"用户id",searchType:"1",hide:1,disabled:1,required:1,options:"1:yes:tag-info,2:no:tag-danger"}`
	jsonStr := makeToJsonStr(str)
	fmt.Println(jsonStr)
}
