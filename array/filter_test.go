package array

import "testing"

func TestFilter(t *testing.T) {
	t.Log(Filter(func(i int) bool {
		return i < 3
	}, []int{1, 2, 3, 4, 5}))

	t.Log(Filter(func(i TestStruct) bool {
		return i.Name == "name2"
	}, TestArr2))
}
