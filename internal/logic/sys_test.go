package logic

import (
	"fmt"
	"regexp"
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
func TestName(t *testing.T) {
	//	str := `{
	//　　　　"label" : {label} ,
	//　　　　 searchType : "hide_222" ,
	//　　　　"hide" : 333 disabled ,
	//　　　　 disabled : "required" ,
	//　　　　"required" : "options" ,
	//　　　　 options:1:yes:tag-info,2:no:tag-danger
	//　　}`
	str := `{"label":用户id,searchType:"1",options:"1:yes:tag-info,2:no:tag-danger",hide:1,disabled:1,required:1} `
	re := regexp.MustCompile(`\s*"?(options)"?\s*:\s*"?((?:[^,]*?:[^,]*?:[^,]*?,?)*)"?\s*([,}])|\s*"?(\w+)"?\s*:\s*"?(.*?)"?\s*([,}])`)
	fmt.Println(re.ReplaceAllString(str, `"$1$4":"$2$5"$3$6`))
	regexp.MustCompile("").ReplaceAllString("", ``)
}
