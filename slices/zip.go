package slices

import "godash"

// Zip returns a slice of pairs of elements from the two input slices.
// If the two slices are not of equal length, the function panics.
func Zip[T, U any](a []T, b []U) (pairs []godash.Pair[T, U]) {
	if len(a) != len(b) {
		panic("length of slices must be equal")
	}

	pairs = []godash.Pair[T, U]{}
	for i := range a {
		pairs[i] = godash.Pair[T, U]{First: a[i], Second: b[i]}
	}
	return
}
