package sets

// Difference returns a new set with the elements that are in the first set but not in the second.
// Returns a new set, does not modify the original sets.
func Difference[K comparable](s1, s2 Set[K]) (s Set[K]) {
	s = Set[K]{}
	for k := range s1 {
		if !Contains(s2, k) {
			s[k] = true
		}
	}
	return s
}
