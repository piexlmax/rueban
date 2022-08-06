package array

// Chunk
// func Chunk[T any](array []T, size int) (res [][]T)
// e.g. array.chunk([]string{"a", "b", "c", "d", "e"}, 2)
// ==> [["a", "b"],["c", "d"]["e"]]
func Chunk[T any](array []T, size int) (res [][]T) {
	for i := 0; i < len(array); i += size {
		start := i
		end := i + size
		if end > len(array) {
			end = len(array)
		}
		res = append(res, array[start:end])
	}
	return res
}
