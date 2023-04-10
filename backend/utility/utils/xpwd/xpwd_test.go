package xpwd

import "testing"

func TestGenPwd(t *testing.T) {
	pwd := GenPwd("1")
	println(pwd)
}
