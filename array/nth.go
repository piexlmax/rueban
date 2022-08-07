package array

// Nth
// func Nth[T any](array []T, index int) T
//
//	e.g. Nth([]int{1,2,3,4,5}, 10)  ==> 5
//		 Nth([]int{1,2,3,4,5}, 2)   ==> 3
//		 Nth([]int{1,2,3,4,5}, -3)   ==> 3
func Nth[T any](array []T, index int) T {
	if index >= len(array) {
		return array[len(array)-1]
	} else {
		if index >= 0 {
			return array[index]
		} else {
			return array[len(array)+index]
		}
	}
}
