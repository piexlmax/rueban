package array

import "testing"

func TestUniq(t *testing.T) {
	t.Log(Uniq(func(i string) string {
		return i
	},
		[]string{"b", "b", "c", "d"}))
	t.Log(Uniq(func(i TestStruct) string {
		return i.Name
	},
		TestArr2))
}
