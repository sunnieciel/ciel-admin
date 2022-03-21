package _0_facade

import (
	"fmt"
	"testing"
)

/*
外观模式

API 为facade 模块的外观接口，大部分代码使用此接口简化对facade类的访问。

facade模块同时暴露了a和b 两个Module 的NewXXX和interface，其它代码如果需要使用细节功能时可以直接调用。
*/
type API interface {
	Say(name string) string
}
type aipImpl struct {
	a HiAPI
	b HelloAPI
}

func (a *aipImpl) Say(name string) string {
	testA := a.a.Say(name)
	testB := a.b.Say(name)
	return fmt.Sprintf("%s %s", testA, testB)
}

// hiAPI 接口
type HiAPI interface {
	Say(name string) string
}
type hiApi struct {
}

func (a *hiApi) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

// helloAPI 接口
type HelloAPI interface {
	Say(name string) string
}
type helloApi struct {
}

func (b *helloApi) Say(name string) string {
	return fmt.Sprintf("hello %s", name)
}
func TestSay(t *testing.T) {
	a := &aipImpl{
		a: &hiApi{},
		b: &helloApi{},
	}
	fmt.Println(a.Say("ciel"))
}
