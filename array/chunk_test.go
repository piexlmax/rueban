package array

import (
	"testing"
)

type TestStruct struct {
	Name string
}

var TestArr1 = []int{1, 2, 3, 4, 5}
var TestArr2 = []TestStruct{
	TestStruct{
		Name: "name1",
	},
	TestStruct{
		Name: "name2",
	},
	TestStruct{
		Name: "name3",
	},
	TestStruct{
		Name: "name4",
	},
	TestStruct{
		Name: "name5",
	},
}

func TestChunk(t *testing.T) {
	res := Chunk(TestArr1, 2)
	if res[2][0] == 5 {
		t.Log(res)
		t.Log("done")
	} else {
		t.Error("error")
	}

	res2 := Chunk[TestStruct](TestArr2, 2)

	if res2[2][0].Name == "name5" {
		t.Log(res2)
		t.Log("done")
	} else {
		t.Error("error")
	}
}
