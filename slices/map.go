// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices

// Map applies the function f to each element of the slice and returns a new slice with the results.
func Map[T, U any](s []T, f func(T) U) (res []U) {
	res = []U{}
	for _, v := range s {
		res = append(res, f(v))
	}
	return
}
