package slices

// Compact removes all nil values from the slice.
func Compact[T any](s []*T) (ret []*T) {
	ret = []*T{}
	for _, v := range s {
		if v != nil {
			ret = append(ret, v)
		}
	}
	return
}
