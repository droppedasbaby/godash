package hashablesets

import "godash"

// AddAll adds a new element to the HashableSet.
// Modifies the set in place, does not return a new set.
func AddAll[H godash.Hashable](s *HashableSet[H], es ...H) {
	for _, e := range es {
		(*s)[e.Hash()] = e
	}
}
