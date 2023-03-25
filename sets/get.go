package sets

// Get returns the element in the sets if it exists. Mostly created for consistency with HashableSets.
// Returns the element if it exists, and a bool indicating whether it exists.
func Get[T comparable](s Set[T], e T) *T {
	if s[e] {
		return &e
	}
	return nil
}
