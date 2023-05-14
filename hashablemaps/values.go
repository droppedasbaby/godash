// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/maps"
)

// Values returns a slice of the values in the HashableMap.
func Values[H godash.Hashable](m HashableMap[H]) []H {
	return maps.Values(m)
}
