package array

import "testing"

func TestReverse(t *testing.T) {
	var array = []string{"a", "b", "c", "d", "e"}
	var res = Reverse(array)
	t.Log(res)
	if res[0] != "e" {
		t.Error("Reverse error")
	}
}
