package array

import "testing"

func TestIndexOf(t *testing.T) {
	t.Log(IndexOf([]string{"a", "b", "c", "d", "e"}, "a", func(i string) string { return i }, 0))
	t.Log(IndexOf([]string{"a", "b", "c", "d", "e"}, "a", func(i string) string { return i }, 3))
}
