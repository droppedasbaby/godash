// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices

// Head returns the first element in a slice, this is the same as the 0th element.
func Head[T any](s []T) T {
	return s[0]
}
