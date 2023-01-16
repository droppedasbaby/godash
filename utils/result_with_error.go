package utils

import "testing"

type ResultWithError[V any] struct {
	Value V
	Error error
}

// NewResultWithError returns a new ResultWithError with the given value and error.
func NewResultWithError[V any](value V, err error) ResultWithError[V] {
	return ResultWithError[V]{Value: value, Error: err}
}

// IsOk returns true if the ResultWithError is not nil.
func (w ResultWithError[V]) IsOk() bool {
	return w.Error == nil
}

// IsErr returns true if the ResultWithError is nil.
func (w ResultWithError[V]) IsErr() bool {
	return w.Error != nil
}

// Unwrap returns the value of the ResultWithError or panics with the error.
func (w ResultWithError[V]) Unwrap() V {
	if w.Error != nil {
		panic(w.Error)
	}
	return w.Value
}

// UnwrapOr returns the value of the ResultWithError or the provided default value.
func (w ResultWithError[V]) UnwrapOr(defaultValue V) V {
	if w.Error != nil {
		return defaultValue
	}
	return w.Value
}

// UnwrapOrElse returns the value of the ResultWithError or the result of the provided function.
func (w ResultWithError[V]) UnwrapOrElse(f func() V) V {
	if w.Error != nil {
		return f()
	}
	return w.Value
}

// FailTestIfErr fails the test if the ResultWithError is an error.
func (w ResultWithError[V]) FailTestIfErr(t *testing.T) {
	t.Helper()
	if w.IsErr() {
		t.Errorf("ResultWithError should not be an error: %v", w.Error)
	}
}

// FailTestIfOk fails the test if the ResultWithError is ok.
func (w ResultWithError[V]) FailTestIfOk(t *testing.T) {
	t.Helper()
	if w.IsOk() {
		t.Errorf("ResultWithError should be an error: %v", w.Error)
	}
}

func (w ResultWithError[V]) FailTestIfErrf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsErr() {
		t.Errorf(format, args...)
	}
}

func (w ResultWithError[V]) FailTestIfOkf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsOk() {
		t.Errorf(format, args...)
	}
}
