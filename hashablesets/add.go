package hashablesets

import "godash"

// Add adds a new element to the HashableSet.
// Modifies the set in place, does not return a new set.
func Add[H godash.Hashable](s *HashableSet[H], element H) {
	(*s)[element.Hash()] = element
}
