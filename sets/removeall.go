package sets

// RemoveAll removes all elements passed as arguments from the sets.
// Modifies the map in place, does not return a new map.
func RemoveAll[K comparable](s *Set[K], keys ...K) {
	for _, k := range keys {
		delete(*s, k)
	}
}
