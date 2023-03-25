package sets

// Get returns the element in the set if it exists. Mostly created for consistency with HashableSets.
// Returns a pointer to the element, or nil if the element is not in the set.
func Get[T comparable](s Set[T], e T) *T {
	if s[e] {
		return &e
	}
	return nil
}
