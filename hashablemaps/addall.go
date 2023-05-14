// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps

import "github.com/GrewalAS/godash"

// AddAll adds a new element to the HashableMap.
// Modifies the map in place, does not return a new HashableMap.
func AddAll[H godash.Hashable](m *HashableMap[H], v ...H) {
	for _, h := range v {
		(*m)[h.Hash()] = h
	}
}
