// Package sets provides functionality for creating and manipulating
// set data structures in Go.
package sets

// Set is a set of values. Just a wrapper around a map where the keys are the sets elements and the values are always
// true.
type Set[T comparable] map[T]bool

// ToSet returns a sets from a slice.
func ToSet[C comparable](slice []C) (s Set[C]) {
	s = Set[C]{}
	for _, v := range slice {
		s[v] = true
	}
	return s
}

// FromSet returns a slice from a sets.
func FromSet[C comparable](set Set[C]) (s []C) {
	s = []C{}
	for k := range set {
		s = append(s, k)
	}
	return s
}
