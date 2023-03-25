package hashablemaps

import (
	"godash"
)

// Remove removes the key and associated value from the HashableMap.
func Remove[H godash.Hashable](m *HashableMap[H], key string) {
	delete(*m, key)
}
