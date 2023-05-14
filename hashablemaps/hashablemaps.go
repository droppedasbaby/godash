package hashablemaps

import "github.com/GrewalAS/godash"

// HashableMap is a map of Hashable objects, where the key is the hash of the object.
type HashableMap[H godash.Hashable] map[string]H

// FromHashableMap returns an array of the key value pairs in the map.
func FromHashableMap[H godash.Hashable](m HashableMap[H]) (ps []godash.Pair[string, H]) {
	ps = []godash.Pair[string, H]{}
	for k, v := range m {
		ps = append(ps, godash.Pair[string, H]{First: k, Second: v})
	}
	return
}

// ToHashableMap returns a map of the key value pairs passed in.
func ToHashableMap[H godash.Hashable](hs []H) (m HashableMap[H]) {
	m = HashableMap[H]{}
	for _, h := range hs {
		m[h.Hash()] = h
	}
	return
}
