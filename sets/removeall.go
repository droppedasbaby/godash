package sets

// RemoveAll removes all elements passed as arguments from the set.
// Modifies the set in place, does not return a new set.
func RemoveAll[K comparable](s *Set[K], keys ...K) {
	for _, k := range keys {
		delete(*s, k)
	}
}
