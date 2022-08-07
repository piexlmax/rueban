package array

// Diffrernce
// func Diffrernce[T any, V comparable](array []T, spArr []T, sp diffFunc[T, V]) []T
// array.Diffrernce(array ,diffArr,func(t T) V)
// the first parameter is the slice to be modified
// the second parameter is the slice that needs to be deleted
// the third parameter returns comparable data for the callback function to determine the reference value that needs to be removed.
// e.g.  Diffrernce([]int{1,2,3,4,5}, []int{2, 3},
//
//	func(i int) int {
//		return i
//	})
//
// ==> []int{1,4,5}
func Diffrernce[T any, V comparable](array []T, spArr []T, sp diffFunc[T, V]) []T {
	var diffSpArr []V

	for _, item := range spArr {
		diffSpArr = append(diffSpArr, sp(item))
	}

	var resArr []T
	for _, item := range array {
		flag := true
		for _, diff := range diffSpArr {
			if diff == sp(item) {
				flag = false
				break
			}
		}
		if flag {
			resArr = append(resArr, item)
		}
	}
	return resArr
}
