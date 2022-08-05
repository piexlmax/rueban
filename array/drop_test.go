package array

import "testing"

func TestDrop(t *testing.T) {

	res := Drop(TestArr1, 0, 2)

	t.Log(res)
	res2 := Drop(TestArr2, 0, 1)
	t.Log(res2)

}
