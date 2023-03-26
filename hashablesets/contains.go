package hashablesets

import "godash"

// Contains returns true if the HashableSet contains the element.
func Contains[H godash.Hashable](s HashableSet[H], e H) bool {
	_, ok := s[e.Hash()]
	return ok
}
