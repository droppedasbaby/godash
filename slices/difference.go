package slices

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/sets"
)

// DifferenceWith returns the difference between two slices using the provided predicate.
func DifferenceWith[T any](a, b []T, p func(T, T) bool) (ret []T) {
	ret = []T{}
	for _, v := range a {
		if !ContainsWith(b, v, p) {
			ret = append(ret, v)
		}
	}
	return
}

// DifferenceHashable returns the difference between two slices using the provided hash function.
func DifferenceHashable[H godash.Hashable](a, b []H) []H {
	aSet := hashablesets.ToSet(a)
	bSet := hashablesets.ToSet(b)
	res := hashablesets.Difference(aSet, bSet)
	return hashablesets.FromSet(res)
}

// Difference returns the difference between two slices.
func Difference[T comparable](a, b []T) (ret []T) {
	aSet := sets.ToSet(a)
	bSet := sets.ToSet(b)
	res := sets.Difference(aSet, bSet)
	return sets.FromSet(res)
}
