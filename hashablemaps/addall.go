package maps

import "godash"

// AddAll adds a new element to the HashableMap.
// Modifies the map in place, does not return a new HashableMap.
func AddAll[H godash.Hashable](m *HashableMap[H], v ...H) {
	for _, h := range v {
		(*m)[h.Hash()] = h
	}
}
