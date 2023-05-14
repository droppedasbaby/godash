// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets

import "github.com/GrewalAS/godash"

// RemoveAll removes all elements passed as arguments from the HashableSet.
// Modifies the set in place, does not return a new set.
func RemoveAll[H godash.Hashable](s *HashableSet[H], elements ...H) {
	for _, e := range elements {
		delete(*s, e.Hash())
	}
}
