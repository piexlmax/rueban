package array

import "testing"

func TestFill(t *testing.T) {

	res := Fill(TestArr1, 4, 0, 2)

	t.Log(res)
	res2 := Fill(TestArr2, TestStruct{Name: "qm"}, 0, 3)
	t.Log(res2)

}
