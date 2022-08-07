package array

import "testing"

func TestPull(t *testing.T) {
	t.Log(Pull(func(i int) int { return i }, []int{1, 2, 2, 2, 23, 3, 12, 4, 5}, 2))
	t.Log(Pull(func(i TestStruct) string { return i.Name }, TestArr2, "name2"))
}
