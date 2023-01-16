package godash

import (
	"reflect"
)

// ComparableEquality Compares two values of type comparable and returns true if they are equal.
func ComparableEquality[A comparable](a, b A) bool {
	return a == b
}

// NumberLessThan compares two numbers and returns true if a is less than b.
func NumberLessThan[N Number](a, b N) bool {
	return a < b
}

// HashableEquality compares two values of any type that implement the Hashable interface and returns true if they are
// equal.
func HashableEquality[A Hashable, B Hashable](a A, b B) bool {
	return a.Hash() == b.Hash()
}

// AnyEquality Compares two values of any type and returns true if they are equal using reflect.DeepEqual.
func AnyEquality[A any](a, b A) bool {
	return reflect.DeepEqual(a, b)
}
