package slices

// Tail returns all elements of the slice except the head. This is the same as the slice from index 1 to the end.
func Tail[T any](s []T) []T {
	return s[1:]
}
