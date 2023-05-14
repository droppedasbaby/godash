// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices

// Reduce applies the function f to each element of the slice and returns a single value.
func Reduce[T, U any](s []T, f func(U, T) U, init U) (res U) {
	res = init
	for _, v := range s {
		res = f(res, v)
	}
	return
}
