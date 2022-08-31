package sys

import (
	"ciel-admin/utility/utils/xtrans"
	"fmt"
	"testing"
)

func TestLanguage(t *testing.T) {
	fmt.Println(xtrans.T("zh-CN", "back"))
}
