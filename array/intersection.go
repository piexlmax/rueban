package array

// Intersection
// func Intersection[T any, V comparable](fn diffFunc[T, V], array ...[]T) []T
//
//	e.g. array.Intersection(func(i string) string {
//		return i
//	},
//		[]string{"a", "b", "c"}, []string{"b", "c", "d"}, []string{"b", "c", "f"})
//
// ==> ["b", "c"]
func Intersection[T any, V comparable](fn diffFunc[T, V], array ...[]T) []T {
	var res []T
	if len(array) == 0 {
		return res
	}

	if len(array) == 1 {
		res = array[0]
		return res
	}

	var sameMap = make(map[V]map[string]int)
	for i, item := range array {
		for ii, item2 := range item {
			if sameMap[fn(item2)] == nil {
				sameMap[fn(item2)] = make(map[string]int)
			}
			if sameMap[fn(item2)]["count"] == i {
				sameMap[fn(item2)]["count"] = i + 1
				if i == 0 {
					sameMap[fn(item2)]["index"] = ii
				}
			}
		}
	}
	for _, v := range sameMap {
		if v["count"] == len(array) {
			res = append(res, array[0][v["index"]])
		}
	}
	return res
}
