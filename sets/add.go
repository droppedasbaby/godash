package sets

// Add adds a new element to the set.
// Modifies the set in place, does not return a new set.
func Add[K comparable](s *Set[K], k K) {
	(*s)[k] = true
}
