package array

// Pull
// func Pull[T any, V comparable](fn diffFunc[T, V], array []T, flag V) []T
//
// e.g. Pull(func(i int) int { return i }, []int{1, 2, 2, 2, 23, 3, 12, 4, 5}, 2) ==> [1 23 3 12 4 5]
// e.g. Pull(func(i TestStruct) string { return i.Name }, TestArr2, "name2") ==> [{name1} {name3} {name4} {name5}]
func Pull[T any, V comparable](fn diffFunc[T, V], array []T, flag V) []T {
	var res []T
	for i := range array {
		if fn(array[i]) != flag {
			res = append(res, array[i])
		}
	}
	return res
}
