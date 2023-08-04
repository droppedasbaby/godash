// Package godash provides a collection of utility functions for Go,
// inspired by JavaScript's lodash library.
package godash

import (
	"reflect"
)

// ComparableEquality Compares two values of type comparable and returns true if they are equal.
func ComparableEquality[A comparable](a, b A) bool {
	return a == b
}

// NumberEquality compares two numbers and returns true if they are equal.
func NumberEquality[N Number](a, b N) bool {
	return a == b
}

// NumberLessThan compares two numbers and returns true if a is less than b.
func NumberLessThan[N Number](a, b N) bool {
	return a < b
}

// NumberLessThanOrEqual compares two numbers and returns true if a is less than or equal to b.
func NumberLessThanOrEqual[N Number](a, b N) bool {
	return a <= b
}

// NumberGreaterThan compares two numbers and returns true if a is greater than b.
func NumberGreaterThan[N Number](a, b N) bool {
	return a > b
}

// NumberGreaterThanOrEqual compares two numbers and returns true if a is greater than or equal to b.
func NumberGreaterThanOrEqual[N Number](a, b N) bool {
	return a >= b
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
