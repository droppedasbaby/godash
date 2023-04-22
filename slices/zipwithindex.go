package slices

import "godash"

// ZipWithIndex returns a slice of pairs of index and the matching element from the input slice.
func ZipWithIndex[T any](s []T) (pairs []godash.Pair[int, T]) {
	pairs = []godash.Pair[int, T]{}
	for i, e := range s {
		pairs[i] = godash.Pair[int, T]{First: i, Second: e}
	}
	return
}