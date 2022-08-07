package array

// Filter
// func Filter[T any](fn loginFunc[T], array []T) []T
//
//	e.g. Filter(func(i TestStruct) bool {
//			return i.Name == "name2"
//		}, TestArr2) ==> [{name2}]
//
//	e.g. Filter(func(i int) bool {
//			return i < 3
//		}, []int{1, 2, 3, 4, 5}) ==> [1,2]
func Filter[T any](fn loginFunc[T], array []T) []T {
	var res []T
	for i := range array {
		if fn(array[i]) {
			res = append(res, array[i])
		}
	}
	return res
}
