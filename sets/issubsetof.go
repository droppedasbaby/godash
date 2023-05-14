// Package sets provides functionality for creating and manipulating
// set data structures in Go.
package sets

// IsSubsetOf checks if a is a subset of b. Returns true if a is a subset of b, false otherwise.
func IsSubsetOf[T comparable](a, b Set[T]) bool {
	for v := range a {
		if !Contains(b, v) {
			return false
		}
	}
	return true
}
