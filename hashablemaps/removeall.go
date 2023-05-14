// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps

import "github.com/GrewalAS/godash"

// RemoveAll removes the keys with the associated values from the HashableMap.
func RemoveAll[H godash.Hashable](m *HashableMap[H], ks ...string) {
	for _, k := range ks {
		delete(*m, k)
	}
}
