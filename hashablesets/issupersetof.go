package hashablesets

import "godash"

// IsSupersetOf checks if a is a superset of b. Returns true if a is a superset of b, false otherwise.
func IsSupersetOf[H godash.Hashable](a, b HashableSet[H]) bool {
	for k := range b {
		if _, ok := a[k]; !ok {
			return false
		}
	}
	return true
}
