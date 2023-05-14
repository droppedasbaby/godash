package sets

// Union returns a new set with the elements that are in either set.
// Returns a new set, does not modify the original set.
func Union[T comparable](a, b Set[T]) (s Set[T]) {
	s = Set[T]{}
	for e := range a {
		s[e] = true
	}
	for e := range b {
		s[e] = true
	}
	return
}
