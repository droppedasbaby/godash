package hashablesets

import "github.com/GrewalAS/godash"

// Difference returns a new HashableSet with the elements that are in the first set but not in the second.
// Returns a new set, does not modify the original sets.
func Difference[H godash.Hashable](a HashableSet[H], b HashableSet[H]) (c HashableSet[H]) {
	c = HashableSet[H]{}
	for _, v := range a {
		e, ok := v.(H)
		if !ok {
			panic("element is not hashable")
		}
		if !Contains[H](b, e) {
			c[v.Hash()] = v
		}
	}
	return
}
