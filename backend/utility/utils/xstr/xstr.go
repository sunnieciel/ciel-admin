package xstr

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

func CheckEasyUname(uname string) error {
	num := len(uname)
	if num < 4 || num > 16 {
		return errors.New("用户名为4到16位字符")
	}
	return nil
}

func Like(str string) string {
	return fmt.Sprint("%", str, "%")
}
func LikePre(str string) string {
	return fmt.Sprint(str, "%")
}
func LikeSuffix(str string) string {
	return fmt.Sprint("%", str)
}
func ContainsAny(s string, substr ...string) bool {
	for _, item := range substr {
		if strings.Contains(s, item) {
			return true
		}
	}
	return false
}

func RandomUnameWithDate() int64 {
	return time.Now().Unix()
}
