package array

import (
	"testing"
)

func TestRemove(t *testing.T) {
	t.Log(Remove(func(i TestStruct) bool {
		return i.Name == "name2"
	}, TestArr2))

	t.Log(Remove(func(i int) bool {
		return i < 3
	}, []int{1, 2, 3, 4, 5}))
}
