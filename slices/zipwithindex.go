package slices

import "github.com/GrewalAS/godash"

// ZipWithIndex returns a slice of pairs of index and the matching element from the input slice.
func ZipWithIndex[T any](s []T) (pairs []godash.Pair[int, T]) {
	pairs = []godash.Pair[int, T]{}
	for i, e := range s {
		pairs = append(pairs, godash.Pair[int, T]{First: i, Second: e})
	}
	return
}
