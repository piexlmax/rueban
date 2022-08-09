package array

import "testing"

func TestTail(t *testing.T) {
	var array = []string{"a", "b", "c", "d", "e"}
	var res = Tail(array)
	t.Log(res)
	if res[0] != "b" {
		t.Error("Tail error")
	}
}
