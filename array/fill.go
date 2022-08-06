package array

// Fill
// func Fill[T any](array []T, value T, start int, end int) []T
// e.g. array.Fill([]string{"a", "b", "c", "d", "e"},"d", 1,5)
// ==> ["e","d","d","d","d"]
func Fill[T any](array []T, value T, start int, end int) []T {
	var res []T
	for i := 0; i < len(array); i++ {
		if i >= start && i < end {
			res = append(res, value)
		} else {
			res = append(res, array[i])
		}
	}
	return res
}
