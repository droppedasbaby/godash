// Package godash provides a collection of utility functions for Go,
// inspired by JavaScript's lodash library.
package godash

// Number is an interface that represents any numeric type in Go.
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | complex64 | complex128
}

// Pair is a pair of two values.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Hashable is an interface that represents any type that can be converted to a string using a hash function.
type Hashable interface {
	// Hash returns a string representation of the value.
	Hash() string
}

// Iterable is an interface that represents any type that can be iterated over.
type Iterable[T any] interface {
	[]T | string
}
