package array

// Uniq
// func Uniq[T any, V comparable](array []T, fn diffFunc[T, V]) []T
//
//	e.g. array.Uniq(func(i string) string {
//		return i
//	},
//		[]string{"b","b", "c", "d"})
//
// ==> ["b", "c","d"]
func Uniq[T any, V comparable](fn diffFunc[T, V], array []T) []T {
	var sameMap = make(map[V]int)
	var res []T
	for i, t := range array {
		if sameMap[fn(t)] == 0 {
			sameMap[fn(t)] = i
		}
	}
	for _, v := range sameMap {
		res = append(res, array[v])
	}
	return res
}
