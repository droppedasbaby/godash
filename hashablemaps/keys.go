package hashablemaps

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/maps"
)

// Keys returns a slice of the keys in the HashableMap.
func Keys[H godash.Hashable](m HashableMap[H]) []string {
	return maps.Keys(m)
}
