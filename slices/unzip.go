package slices

import "godash"

// Unzip returns two slices, the first containing the first elements of the input pairs,
// and the second containing the second elements of the input pairs.
func Unzip[T, U any](pairs []godash.Pair[T, U]) (a []T, b []U) {
	a = []T{}
	b = []U{}
	for _, p := range pairs {
		a = append(a, p.First)
		b = append(b, p.Second)
	}
	return
}
