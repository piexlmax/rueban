package array

// Find
// func Find[T any](array []T, fn func(T) bool) ([]T, bool, []int)
//
//	e.g. array.Find([]string{"a", "b", "c", "d", "e"}, func(v string) bool {
//			return v != "c"
//		})
//
// ==> ["a", "b", "d", "e"], true, [0, 1, 3, 4]
func Find[T any](array []T, fn func(T) bool) ([]T, bool, []int) {
	res, isFind, indexs := find(array, fn, false)
	return res, isFind, indexs
}

// FindFirst
// func FindFirst[T any](array []T, fn func(T) bool) (T, bool, int)
//
//	e.g. array.FindFirst([]string{"a", "b", "c", "d", "e"}, func(v string) bool {
//			return v != "c"
//		})
//
// ==> ["a"], true, 0
func FindFirst[T any](array []T, fn func(T) bool) (T, bool, int) {
	res, isFind, indexs := find(array, fn, true)
	if isFind {
		return res[0], isFind, indexs[0]
	} else {
		var zero T
		return zero, isFind, -1
	}
}

func find[T any](array []T, fn func(T) bool, onlyOne bool) (res []T, isFind bool, indexs []int) {
	for i := range array {
		if fn(array[i]) {
			res = append(res, array[i])
			isFind = true
			indexs = append(indexs, i)
			if onlyOne {
				return res, isFind, indexs
			}
		}
	}
	return res, isFind, indexs
}
