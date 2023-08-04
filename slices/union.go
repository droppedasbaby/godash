package slices

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/sets"
)

// Union returns a slice containing the elements of the two input slices. Standard union semantics apply.
// This is a fast implementation because it allows only comparable types.
func Union[T comparable](a []T, b []T) (c []T) {
	aSet := sets.ToSet(a)
	bSet := sets.ToSet(b)
	cSet := sets.Union(aSet, bSet)
	c = sets.FromSet(cSet)

	return
}

// UnionHashable returns a slice containing the elements of the two input slices. Standard union semantics apply.
// This is a fast implementation because it allows only hashable types.
func UnionHashable[T godash.Hashable](a []T, b []T) (c []T) {
	aSet := hashablesets.ToSet(a)
	bSet := hashablesets.ToSet(b)
	cSet := hashablesets.Union(aSet, bSet)
	c = hashablesets.FromSet(cSet)

	return
}

// UnionWith returns a slice containing the elements of the two input slices. Standard union semantics apply.
// This is a slow implementation, but it allows for custom equality semantics.
// The predicate p is used to determine whether two elements are equal.
func UnionWith[T any](a []T, b []T, p func(T, T) bool) (c []T) {
	c = []T{}
	c = append(c, a...)
	for _, e := range b {
		if !ContainsWith(c, e, p) {
			c = append(c, e)
		}
	}

	return
}
