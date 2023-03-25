package maps

import "godash"

// RemoveAll removes the keys with the associated values from the HashableMap.
func RemoveAll[H godash.Hashable](m *HashableMap[H], keys []string) {
	for _, k := range keys {
		delete(*m, k)
	}
}
