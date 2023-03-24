package maps

// Add adds the key value pair to the map. Mostly created for consistency with HashableMap.
// Modifies the map in place, does not return a new map.
func Add[K comparable, V any](m *map[K]V, k K, v V) {
	(*m)[k] = v
}
