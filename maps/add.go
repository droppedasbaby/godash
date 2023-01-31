package maps

import (
	"godash"
)

// Add makes a copy of the map passed in and adds the pairs passed in. Original map remains unmodified.
func Add[K comparable, V any](m map[K]V, p ...godash.Pair[K, V]) (n map[K]V) {
	n = make(map[K]V)

	for k, v := range m {
		n[k] = v
	}

	for _, i := range p {
		n[i.First] = i.Second
	}

	return n
}
