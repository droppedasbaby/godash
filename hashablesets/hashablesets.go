package hashablesets

import "godash"

// HashableSet is a set of values. Just a wrapper around a map where the keys are the sets elements and the values are
// always true.
type HashableSet[H godash.Hashable] map[string]godash.Hashable

// FromSet returns a slice from a HashableSet.
func FromSet[H godash.Hashable](s HashableSet[H]) (vs []H) {
	vs = []H{}
	for _, v := range s {
		vh, ok := v.(H)
		if !ok {
			panic("element is not hashable")
		}
		vs = append(vs, vh)
	}
	return
}

// ToSet returns a HashableSet from a slice.
func ToSet[H godash.Hashable](vs []H) (s HashableSet[H]) {
	s = HashableSet[H]{}
	for _, v := range vs {
		s[v.Hash()] = v
	}
	return
}
