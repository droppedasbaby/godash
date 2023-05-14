// Package maps provides functions for manipulating maps in Go.
package maps

// Remove removes the key and associated value from the map. Mostly created for consistency with HashableMap.
func Remove[K comparable, V any](m *map[K]V, k K) {
	delete(*m, k)
}
