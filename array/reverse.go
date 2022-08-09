package array

// Reverse
// func Reverse[T any](array []T) []T
//
// e.g. array.Reverse([]int{1, 2, 3, 4, 5}) ==> [5 4 3 2 1]
func Reverse[T any](array []T) []T {
	var res []T
	for i := len(array) - 1; i >= 0; i-- {
		res = append(res, array[i])
	}
	return res
}
