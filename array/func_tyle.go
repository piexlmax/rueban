package array

type diffFunc[T any, V comparable] func(t T) V
type loginFunc[T any] func(t T) bool
