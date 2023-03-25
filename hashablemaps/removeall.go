package hashablemaps

import "godash"

// RemoveAll removes the keys with the associated values from the HashableMap.
func RemoveAll[H godash.Hashable](m *HashableMap[H], ks ...string) {
	for _, k := range ks {
		delete(*m, k)
	}
}
