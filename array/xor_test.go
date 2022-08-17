package array

import (
	"fmt"
	"testing"
)

func diffF(i int) int {
	return i
}

func TestXor(t *testing.T) {

	fmt.Println(Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}))
	fmt.Println(Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4}))
	fmt.Println(Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6}))
	fmt.Println(Xor(diffF, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 6}, []int{1, 2, 3, 4, 7}))
}
