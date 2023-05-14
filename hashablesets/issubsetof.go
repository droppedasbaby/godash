// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets

import "github.com/GrewalAS/godash"

// IsSubsetOf checks if a is a subset of b. Returns true if a is a subset of b, false otherwise.
func IsSubsetOf[H godash.Hashable](a, b HashableSet[H]) bool {
	for k := range a {
		_, ok := b[k]
		if !ok {
			return false
		}
	}
	return true
}
