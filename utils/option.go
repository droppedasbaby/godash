package utils

import (
	"reflect"
)

// Option is a wrapper around a value that can be nil.
type Option[V any] struct {
	Value *V
}

// NewOption returns a new Option.
func NewOption[V any](value V) Option[V] {
	return Option[V]{Value: &value}
}

// IsNil returns true if the Option is nil.
func (o Option[V]) IsNil() bool {
	return reflect.TypeOf(*o.Value) == reflect.TypeOf(nil)
}

// IsSome returns true if the Option is not nil.
func (o Option[V]) IsSome() bool {
	return !o.IsNil()
}

// Unwrap returns the value of the Option or panics.
func (o Option[V]) Unwrap() V {
	if o.IsNil() {
		panic("Option is nil")
	}
	return *o.Value
}

// UnwrapOr returns the value of the Option or the provided default value.
func (o Option[V]) UnwrapOr(defaultValue V) V {
	if o.IsNil() {
		return defaultValue
	}
	return *o.Value
}

// UnwrapOrElse returns the value of the Option or the result of the provided function.
func (o Option[V]) UnwrapOrElse(f func() V) V {
	if o.IsNil() {
		return f()
	}
	return *o.Value
}
