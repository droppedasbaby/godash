package search

// BisectLeft returns the left most index where the target would be inserted in the sorted slice.
func BisectLeft[T any](slice []T, target T) int {
	return -1
}

// Bisect returns the index of the first occurrence of the target in the slice or -1 if not found.
func Search[T comparable](slice []T, target T) int {
	return -1
}

// BisectRight returns the right most index where the target would be inserted in the sorted slice.
func BisectRight[T comparable](slice []T, target T) int {
	return -1
}

// BisectLeftWith returns the left most index where the target would be inserted in the sorted slice.
// The predicate function p is used to compare the target with the elements of the slice.
func BisectLeftWith[T any](slice []T, target T, p func(T, T) bool) int {
	return -1
}

// SearchWith returns the index of the first occurrence of the target in the slice or -1 if not found.

func SearchWith[T any](slice []T, target T, p func(T, T) bool) int {
	return -1
}

// BisectRightWith returns the right most index where the target would be inserted in the sorted slice.
// The predicate function p is used to compare the target with the elements of the slice.
func BisectRightWith[T any](slice []T, target T, p func(T, T) bool) int {
	return -1
}
