package hashablesets

import "github.com/GrewalAS/godash"

// Union returns a new HashableSet with the elements that are in either set.
// Returns a new set, does not modify the original set.
func Union[H godash.Hashable](a, b HashableSet[H]) (c HashableSet[H]) {
	c = HashableSet[H]{}
	for k, v := range a {
		c[k] = v
	}
	for k, v := range b {
		c[k] = v
	}
	return
}
