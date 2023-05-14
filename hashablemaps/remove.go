// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps

import (
	"github.com/GrewalAS/godash"
)

// Remove removes the key and associated value from the HashableMap.
func Remove[H godash.Hashable](m *HashableMap[H], key string) {
	delete(*m, key)
}
