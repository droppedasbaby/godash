package slices

import "godash/sets"

// Difference returns the difference between two slices.
func Difference[T comparable](a, b []T) (ret []T) {
	aSet := sets.ToSet(a)
	bSet := sets.ToSet(b)
	res := sets.Difference(aSet, bSet)
	ret = sets.FromSet(res)
	return
}
