package maps

import "godash"

// FromMap returns an array of the key value pairs in the map.
func FromMap[K comparable, V any](m map[K]V) (ps []godash.Pair[K, V]) {
	ps = []godash.Pair[K, V]{}
	for k, v := range m {
		ps = append(ps, godash.Pair[K, V]{First: k, Second: v})
	}
	return
}

// ToMap returns a map of the key value pairs passed in.
func ToMap[K comparable, V any](ps ...godash.Pair[K, V]) (m map[K]V) {
	m = map[K]V{}
	for _, i := range ps {
		m[i.First] = i.Second
	}
	return
}
