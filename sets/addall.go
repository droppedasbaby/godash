// Package sets provides functionality for creating and manipulating
// set data structures in Go.
package sets

// AddAll adds a new element to the set.
// Modifies the set in place, does not return a new set.
func AddAll[K comparable](s *Set[K], ks ...K) {
	for _, k := range ks {
		(*s)[k] = true
	}
}
