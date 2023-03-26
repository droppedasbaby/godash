package hashablesets

import "godash"

// Remove removes an element from the HashableSet.
// Modifies the set in place, does not return a new set.
func Remove[H godash.Hashable](s *HashableSet[H], e H) {
	delete(*s, e.Hash())
}
