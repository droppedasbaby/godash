package maps

// Values returns a slice of the values in the map.
func Values[K comparable, V any](m map[K]V) (vs []V) {
	vs = []V{}
	for _, v := range m {
		vs = append(vs, v)
	}
	return
}
