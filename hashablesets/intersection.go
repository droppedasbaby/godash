package hashablesets

import "github.com/GrewalAS/godash"

// Intersection returns a new HashableSet with the elements that are in both sets.
// Returns a new sets, does not modify the original sets.
func Intersection[H godash.Hashable](a, b HashableSet[H]) (c HashableSet[H]) {
	c = HashableSet[H]{}
	for _, v := range a {
		e, ok := v.(H)
		if !ok {
			panic("element is not hashable")
		}
		if Contains[H](b, e) {
			c[v.Hash()] = v
		}
	}
	return c
}
