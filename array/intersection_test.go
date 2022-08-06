package array

import "testing"

type TestIntersectionStruct struct {
	Name string
	Age  int
}

var TArr = []TestIntersectionStruct{
	TestIntersectionStruct{
		Name: "qm",
		Age:  10,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  12,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  18,
	},
}

var TArr2 = []TestIntersectionStruct{
	TestIntersectionStruct{
		Name: "qm",
		Age:  10,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  12,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  18,
	},
}

var TArr3 = []TestIntersectionStruct{
	TestIntersectionStruct{
		Name: "qm",
		Age:  10,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  12,
	},
	TestIntersectionStruct{
		Name: "qm",
		Age:  18,
	},
}

func TestIntersection(t *testing.T) {
	t.Log(Intersection(func(item TestIntersectionStruct) int {
		return item.Age
	}, TArr, TArr2, TArr3))

	t.Log(Intersection(func(i string) string {
		return i
	},
		[]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "c", "f"}))
}
