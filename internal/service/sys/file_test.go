package sys

import "testing"

func TestRemoveFile(t *testing.T) {
	err := RemoveFile(nil, ".")
	if err != nil {
		panic(err)
	}
}
