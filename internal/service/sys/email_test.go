package sys

import "testing"

func TestSendEmail(t *testing.T) {
	err := SendEmail(nil, []string{"1211sciel@gmail.com"}, "hello", "hello")
	if err != nil {
		panic(err)
	}
}
