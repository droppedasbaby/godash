package maps

// RemoveAll removes the keys with the associated values from the map.
func RemoveAll[K comparable, V any](m *map[K]V, ks ...K) {
	for _, k := range ks {
		delete(*m, k)
	}
}
