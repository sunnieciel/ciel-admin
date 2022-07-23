package logic

import "testing"

func TestCheckOrSave(t *testing.T) {
	err := checkGroupOrSave(nil, "haha")
	if err != nil {
		t.Fatal(err)
	}
}
