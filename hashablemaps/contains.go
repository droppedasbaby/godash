package hashablemaps

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/maps"
)

// Contains returns true if the hashmap contains the key.
func Contains[H godash.Hashable](m HashableMap[H], key string) bool {
	return maps.Contains(m, key)
}
