package sets

// Union returns a new sets with the elements that are in either sets.
// Returns a new sets, does not modify the original sets.
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
