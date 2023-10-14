package search

// BisectLeftWith locates the insertion point for target t in slice s to maintain sorted order.
// The elements of s must be sorted in ascending order, but need not be unique.
// The predicate function p is used to compare the target with the elements of the slice.
func BisectLeftWith[T any](s []T, t T, p func(T, T) bool) (l int) {
	l, r := 0, len(s)
	for l < r {
		//nolint:gomnd // this "magic number" is foundational to BS
		m := l + (r-l)/2
		if p(s[m], t) {
			l = m + 1
		} else {
			r = m
		}
	}
	return
}

// BisectRightWith is similar to BisectLeftWith().
// Returns an insertion point which comes after (to the right of) any existing entries of t in s.
func BisectRightWith[T any](s []T, t T, p func(T, T) bool) (l int) {
	l, r := 0, len(s)
	for l < r {
		//nolint:gomnd // this "magic number" is foundational to BS
		m := l + (r-l)/2
		if p(t, s[m]) {
			r = m
		} else {
			l = m + 1
		}
	}
	return
}

// BisectWith locates the insertion point for target t in slice s to maintain sorted order.
// Behaves like BisectRightWith().
func BisectWith[T any](s []T, t T, p func(T, T) bool) int {
	return BisectRightWith(s, t, p)
}
