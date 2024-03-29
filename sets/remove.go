package sets

// Remove removes an element from the set.
// Modifies the set in place, does not return a new set.
func Remove[K comparable](s *Set[K], k K) {
	delete(*s, k)
}
