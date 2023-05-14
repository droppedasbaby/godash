package slices

import (
	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/sets"
)

// Intersection return the intersection of the two input slices. Standard union semantics apply.
// This is a fast implementation because it allows only comparable types.
func Intersection[T comparable](a []T, b []T) (c []T) {
	aSet := sets.ToSet(a)
	bSet := sets.ToSet(b)
	cSet := sets.Intersection(aSet, bSet)
	c = sets.FromSet(cSet)

	return
}

// IntersectionHashable returns the intersection of the two input slices. Standard union semantics apply.
// This is a slow implementation, but it allows for custom equality semantics.
// The predicate p is used to determine whether two elements are equal.
func IntersectionHashable[T godash.Hashable](a []T, b []T) (c []T) {
	aSet := hashablesets.ToSet(a)
	bSet := hashablesets.ToSet(b)
	cSet := hashablesets.Intersection(aSet, bSet)
	c = hashablesets.FromSet(cSet)

	return
}

// IntersectionWith returns the intersection of the two input slices. Standard union semantics apply.
// This is a slow implementation, but it allows for custom equality semantics.
// The predicate p is used to determine whether two elements are equal.
func IntersectionWith[T comparable](a []T, b []T, p func(T, T) bool) (c []T) {
	c = []T{}
	for _, e := range a {
		if ContainsWith(b, e, p) {
			c = append(c, e)
		}
	}

	return
}
