package array

import (
	"fmt"
	"testing"
)

func TestZip(t *testing.T) {
	fmt.Println(Zip([]string{"b", "b", "c", "d"}, []string{"a", "b", "c", "d"}, []string{"1", "2", "3", "4"}))
	fmt.Println(Zip([]string{"b", "c", "d"}, []string{"a", "b", "c", "d"}, []string{"1", "2", "3", "4"}))
}
