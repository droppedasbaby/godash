package search

// LinearSearch returns the index of the first occurrence of the target in the slice or -1 if not found.
// The elements of the slice and the target must be comparable.
func LinearSearch[T comparable](slice []T, target T) int {
	return LinearSearchWith[T](slice, target, func(a, b T) bool { return a == b })
}

// LinearSearchWith returns the index of the first occurrence of the target in the slice or -1 if not found.
// The predicate function p is used to compare the target with the elements of the slice.
func LinearSearchWith[T any](slice []T, target T, p func(T, T) bool) int {
	for i, v := range slice {
		if p(v, target) {
			return i
		}
	}
	return -1
}
