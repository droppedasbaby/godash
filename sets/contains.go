// Package sets provides functionality for creating and manipulating
// set data structures in Go.
package sets

// Contains returns true if the set contains the element.
func Contains[K comparable](s Set[K], k K) bool {
	_, ok := s[k]
	return ok
}
