package set

// Set is a set of values. Just a wrapper around a map where the keys are the set elements and the values are always
// true.
type Set[T comparable] map[T]bool

// ToSet returns a set from a slice.
func ToSet[C comparable](slice []C) (s Set[C]) {
	s = Set[C]{}
	for _, v := range slice {
		s[v] = true
	}
	return s
}

// FromSet returns a slice from a set.
func FromSet[C comparable](set Set[C]) (s []C) {
	s = make([]C, 0)
	for k := range set {
		s = append(s, k)
	}
	return s
}
