package hashablemaps

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/maps"
)

// Values returns a slice of the values in the HashableMap.
func Values[H godash.Hashable](m HashableMap[H]) []H {
	return maps.Values(m)
}
