package set

// Set is a set of values. Just a wrapper around a map where the keys are the set elements and the values are always
// true.
type Set[T comparable] map[T]bool

// FromSet returns a set from a slice.
func FromSet[C comparable](slice []C) (s Set[C]) {
	s = make(Set[C])
	for _, v := range slice {
		s[v] = true
	}
	return s
}

// ToSet returns a slice from a set.
func ToSet[C comparable](set Set[C]) (s []C) {
	s = make([]C, 0)
	for k := range set {
		s = append(s, k)
	}
	return s
}
