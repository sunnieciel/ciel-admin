package _6_option_model

import (
	"fmt"
	"testing"
)

type Option struct {
	A string
	B string
	C int
}

type OptionFunc func(*Option)

func WithA(a string) OptionFunc {
	return func(o *Option) {
		o.A = a
	}
}
func WithB(b string) func(o *Option) {
	return func(o *Option) {
		o.B = b
	}
}
func WithC(c int) func(o *Option) {
	return func(o *Option) {
		o.C = c
	}
}

var defaultOption = Option{
	A: "A",
	B: "B",
	C: 100,
}

func newOption(a, b string, c int) *Option {
	return &Option{
		A: a,
		B: b,
		C: c,
	}
}

func newOption2(opts ...OptionFunc) *Option {
	option := defaultOption
	for _, opt := range opts {
		opt(&option)
	}
	return &option
}

func TestOption(t *testing.T) {
	x := newOption("nazha", "小王子", 100)
	fmt.Println(x)
	x = newOption2()
	fmt.Println(x)
	x = newOption2(WithA("333"), WithB("小王子a"), WithC(200))
	fmt.Println(x)
}
