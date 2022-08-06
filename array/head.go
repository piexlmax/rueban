package array

// Head
// func Head[T any](array []T) T
//
//	e.g. array.Head([]string{"a", "b", "c", "d", "e"})
//
// ==> "a"
func Head[T any](array []T) T {
	if len(array) > 0 {
		return array[0]
	} else {
		var a T
		return a
	}
}
