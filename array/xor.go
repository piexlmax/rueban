package array

// Xor
// Xor[T any, V comparable](fn diffFunc[T, V], arrays ...[]T) []T
//
//	func diffF(i int) int {
//		return i
//	}
//
//	e.g. array.Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})
//			==> []
//	e.g. array.Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4})
//			==> [5]
//	e.g. array.Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6})
//			==> [5 6]
//	e.g. array.Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6}, []int{1, 2, 3, 4, 7})
//			==> [5 6 7]

func Xor[T any, V comparable](fn diffFunc[T, V], arrays ...[]T) []T {
	var sameMap = make(map[V]int)
	var res []T
	var notStep T
	var waitUseArr = []T{notStep}
	for _, array := range arrays {
		waitUseArr = append(waitUseArr, array...)
	}

	for i, t := range waitUseArr {
		if i == 0 {
			continue
		}
		if sameMap[fn(t)] == 0 {
			sameMap[fn(t)] = i
		} else {
			sameMap[fn(t)] = -1
		}
	}
	for _, v := range sameMap {
		if v >= 0 {
			res = append(res, waitUseArr[v])
		}
	}
	return res
}
