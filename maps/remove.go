package maps

// Remove makes a copy of the map passed in and removes the keys passed in. Original map remains unmodified.
func Remove[K comparable, V any](m map[K]V, k ...K) (n map[K]V) {
	// TODO: Change to set function
	s := make(map[K]bool)
	n = make(map[K]V)

	for _, i := range k {
		s[i] = true
	}

	for k, v := range m {
		if _, ok := s[k]; !ok {
			n[k] = v
		}
	}

	return n
}
