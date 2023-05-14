package sets

// IsSupersetOf checks if a is a superset of b. Returns true if a is a superset of b, false otherwise.
func IsSupersetOf[T comparable](a, b Set[T]) bool {
	for v := range b {
		if !Contains(a, v) {
			return false
		}
	}
	return true
}
