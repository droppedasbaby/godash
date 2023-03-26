package hashablesets

import "godash"

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
