package array

import (
	"testing"
)

func TestDiffrernce(t *testing.T) {
	res := Diffrernce(TestArr1, []int{2, 3},
		func(i int) int {
			return i
		})

	t.Log(res)

	res2 := Diffrernce(TestArr2, TestArr2[:3],
		func(i TestStruct) string {
			return i.Name
		})

	t.Log(res2)
}
