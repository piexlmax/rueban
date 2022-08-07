package array

// Remove
// func Remove[T any](fn loginFunc[T], array []T) []T
//
//	e.g. Remove(func(i TestStruct) bool {
//			return i.Name == "name2"
//		}, TestArr2) ==> [{name1} {name3} {name4} {name5}]
//
//	e.g. Remove(func(i int) bool {
//			return i < 3
//		}, []int{1, 2, 3, 4, 5}) ==> [3 4 5]
func Remove[T any](fn loginFunc[T], array []T) []T {
	var res []T
	for i := range array {
		if !fn(array[i]) {
			res = append(res, array[i])
		}
	}
	return res
}
