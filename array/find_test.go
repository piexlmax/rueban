package array

import (
	"strings"
	"testing"
)

func TestFind(t *testing.T) {
	res, isFind, indexs := Find(TestArr2, func(v TestStruct) bool {
		return strings.Index(v.Name, "name4") != -1
	})
	t.Log(res, isFind, indexs)

	res2, isFind, index := FindFirst(TestArr2, func(v TestStruct) bool {
		return strings.Index(v.Name, "nam2e4") != -1
	})
	t.Log(res2, isFind, index)
}
