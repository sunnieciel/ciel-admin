package _3_singleton

import (
	"github.com/gogf/gf/v2/container/garray"
	"sync"
	"testing"
)

type Singleton interface {
	foo()
}
type singleton struct{}

func (s singleton) foo() {}

var (
	instance *singleton
	one      sync.Once
)

func GetInstance() *singleton {
	one.Do(func() {
		instance = &singleton{}
	})
	return instance
}

const count = 100

func TestSingleton(t *testing.T) {
	arr := garray.New()
	for i := 0; i < count; i++ {
		arr.Append(GetInstance())
	}
	for _, item := range arr.Slice() {
		println(item)
	}
}
