package slices

import "godash"

// ContainsWith returns true if the slice contains the value using the provided predicate.
func ContainsWith[S, C any](s []S, c C, pred func(S, C) bool) bool {
	for _, v := range s {
		if pred(v, c) {
			return true
		}
	}
	return false
}

// ContainsHashable returns true if the slice contains the value using the provided hash function.
func ContainsHashable[H godash.Hashable](s []H, c H) bool {
	return ContainsWith(s, c.Hash(), func(a H, b string) bool {
		return a.Hash() == b
	})
}

// Contains returns true if the slice contains the value.
func Contains[T comparable](s []T, c T) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}
