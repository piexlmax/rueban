package array

// Tail
// func Tail[T any](array []T) []T
// e.g. array.Tail([]int{1, 2, 3, 4, 5}) ==> [2 3 4 5]
func Tail[T any](array []T) []T {
	return array[1:]
}
