// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets

import "github.com/GrewalAS/godash"

// Add adds a new element to the HashableSet.
// Modifies the set in place, does not return a new set.
func Add[H godash.Hashable](s *HashableSet[H], element H) {
	(*s)[element.Hash()] = element
}
