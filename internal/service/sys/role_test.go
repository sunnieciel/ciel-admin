package sys

import (
	"fmt"
	"regexp"
	"testing"
)

func TestApi(t *testing.T) {
	str := "/admin/api/path/del/:id"
	findString := regexp.MustCompile(`.+/del/`).FindString(str)
	fmt.Println("result:", findString)
}
