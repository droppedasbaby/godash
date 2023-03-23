package slices

// Filter returns a new slice containing only the elements of the given slice that satisfy the given predicate.
func Filter[T any](s []T, p func(T) bool) (res []T) {
	res = []T{}
	for _, v := range s {
		if p(v) {
			res = append(res, v)
		}
	}
	return
}
