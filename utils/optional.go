// Package utils contains a collection of utility functions that can be
// used throughout the Go application.
package utils

import (
	"reflect"
)

// Optional is a wrapper around a value that can be nil.
type Optional[V any] struct {
	Value *V
}

// NewOptional returns a new Optional.
func NewOptional[V any](value V) Optional[V] {
	return Optional[V]{Value: &value}
}

// IsNil returns true if the Optional is nil.
func (o Optional[V]) IsNil() bool {
	return reflect.TypeOf(*o.Value) == reflect.TypeOf(nil)
}

// IsSome returns true if the Optional is not nil.
func (o Optional[V]) IsSome() bool {
	return !o.IsNil()
}

// Unwrap returns the value of the Optional or panics.
func (o Optional[V]) Unwrap() V {
	if o.IsNil() {
		panic("Optional is nil")
	}
	return *o.Value
}

// UnwrapOr returns the value of the Optional or the provided default value.
func (o Optional[V]) UnwrapOr(defaultValue V) V {
	if o.IsNil() {
		return defaultValue
	}
	return *o.Value
}

// UnwrapOrElse returns the value of the Optional or the result of the provided function.
func (o Optional[V]) UnwrapOrElse(f func() V) V {
	if o.IsNil() {
		return f()
	}
	return *o.Value
}
