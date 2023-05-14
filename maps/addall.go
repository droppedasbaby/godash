// Package maps provides functions for manipulating maps in Go.
package maps

import "github.com/GrewalAS/godash"

// AddAll adds the key value pairs to the map.
// Modifies the map in place, does not return a new map.
func AddAll[K comparable, V any](m *map[K]V, ps ...godash.Pair[K, V]) {
	for _, i := range ps {
		(*m)[i.First] = i.Second
	}
}
