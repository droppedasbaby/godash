package sets

// Intersection returns a new sets with the elements that are in both sets.
// Returns a new sets, does not modify the original sets.
func Intersection[K comparable](s1, s2 map[K]bool) (s Set[K]) {
	s = Set[K]{}
	for k := range s1 {
		if Contains(s2, k) {
			s[k] = true
		}
	}
	return s
}
