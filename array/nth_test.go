package array

import "testing"

func TestNth(t *testing.T) {
	t.Log(Nth(TestArr1, 10))
	t.Log(Nth(TestArr1, 2))
	t.Log(Nth(TestArr1, -3))
}
