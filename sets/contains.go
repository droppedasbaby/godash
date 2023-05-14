package sets

// Contains returns true if the set contains the element.
func Contains[K comparable](s Set[K], k K) bool {
	_, ok := s[k]
	return ok
}
