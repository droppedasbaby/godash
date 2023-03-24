package sets

// AddAll adds a new element to the sets.
// Modifies the map in place, does not return a new map.
func AddAll[K comparable](s *Set[K], ks ...K) {
	for _, k := range ks {
		(*s)[k] = true
	}
}
