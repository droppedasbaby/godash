package slices

// Foreach iterates over the given slice and calls the given function for each element.
// Should be used for side effects only since it doesn't return a value.
func Foreach[T any](s []T, f func(T)) {
	for _, v := range s {
		f(v)
	}
}
