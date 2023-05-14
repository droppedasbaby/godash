package maps

// Contains returns true if the map contains the key.
func Contains[K comparable, V any](m map[K]V, k K) bool {
	_, ok := m[k]
	return ok
}
