package array

// Drop
// func Drop[T any](array []T, sub ...int) []T
// e.g. array.Drop(['a', 'b', 'c', 'd', 'e'], 2,3)
// ==> ['a', 'b', 'e']
func Drop[T any](array []T, sub ...int) []T {
	var res []T
	var dopMap = make(map[int]bool)
	for i := range sub {
		if sub[i] < len(array) {
			dopMap[sub[i]] = true
		}
	}
	for i := range array {
		if dopMap[i] {
			continue
		}
		res = append(res, array[i])
	}
	return res
}
