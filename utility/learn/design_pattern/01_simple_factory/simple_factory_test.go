package _1_simple_factory

import (
	"fmt"
	"testing"
)

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	}
	return nil
}

// type 1
type hiAPI struct{}

func (h *hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// type 2
type helloAPI struct{}

func (h *helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
func TestAPI(t *testing.T) {
	api := NewAPI(2)
	fmt.Println(api.Say("Tom"))
}
