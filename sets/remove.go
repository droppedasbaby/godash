package sets

// Remove removes an element from the sets.
// Modifies the map in place, does not return a new map.
func Remove[K comparable](s *Set[K], k K) {
	delete(*s, k)
}
