package array

// IndexOf
// func IndexOf[T any, V comparable](array []T, value T, sp diffFunc, fromIndex int) int
//
//	e.g. array.IndexOf([]string{"a", "b", "c", "d", "e"}, "a", func(i string) string { return i }, 3)
//
// ==> -1
//
//	e.g. array.IndexOf([]string{"a", "b", "c", "d", "e"}, "a", func(i string) string { return i }, 0)
//
// ==> 0
func IndexOf[T any, V comparable](array []T, value T, sp diffFunc[T, V], fromIndex int) int {
	var index int = -1
	for i := fromIndex; i < len(array); i++ {
		if sp(value) == sp(array[i]) {
			index = i - fromIndex
			return index
		}
	}
	return index
}
