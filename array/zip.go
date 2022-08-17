package array

import "fmt"

// Zip
// Zip[T any](arrays ...[]T)  [][]T
//
//	e.g. array.Zip([]string{"b","b", "c", "d"}, []string{"a","b","c","d"}, []string{"1","2","3","4"})
//
// ==> [["b", "a","1"],["b", "b","2"],["c", "c","3"],["d", "d","4"]]

func Zip[T any](arrays ...[]T) [][]T {
	var res [][]T
	for i := 0; i < len(arrays); i++ {
		if len(arrays[i]) != len(arrays[0]) {
			fmt.Println("array.Zip: arrays not equal length")
			return res
		}
	}
	for i := 0; i < len(arrays[0]); i++ {
		var row []T
		for j := 0; j < len(arrays); j++ {
			row = append(row, arrays[j][i])
		}
		res = append(res, row)
	}
	return res
}
