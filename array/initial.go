package array

// Initial
// Initial[T any](array []T) []T
//
//	e.g. array.Initial([]string{"a", "b", "c", "d", "e"})
//
// ==> ["a", "b", "c", "d"]
func Initial[T any](array []T) []T {
	if len(array) > 0 {
		return array[:len(array)-1]
	} else {
		var res []T
		return res
	}
}
