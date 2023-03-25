package hashablemaps

import "godash"

// Add adds the key value pair to the HashableMap.
// Modifies the map in place, does not return a new HashableMap.
func Add[H godash.Hashable](m *HashableMap[H], v H) {
	(*m)[v.Hash()] = v
}
