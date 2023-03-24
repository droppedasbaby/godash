package sets

// Contains returns true if the sets contains the element.
func Contains[K comparable](s Set[K], k K) bool {
	_, ok := s[k]
	return ok
}
