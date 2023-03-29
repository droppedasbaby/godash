package slices

// Contains returns true if the slice contains the value.
func Contains[T comparable](s []T, c T) bool {
	for _, v := range s {
		if v == c {
			return true
		}
	}
	return false
}
