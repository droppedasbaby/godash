package maps

// Keys returns a slice of the keys in the map.
func Keys[K comparable, V any](m map[K]V) (ks []K) {
	ks = make([]K, 0)
	for k := range m {
		ks = append(ks, k)
	}
	return
}
