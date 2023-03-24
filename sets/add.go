package sets

// Add adds a new element to the sets.
// Modifies the map in place, does not return a new map.
func Add[K comparable](s *Set[K], k K) {
	(*s)[k] = true
}
