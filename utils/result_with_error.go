package utils

import (
	"fmt"
	"testing"

	"godash/slices"
)

type ResultWithError[V any] struct {
	Value       V
	Error       error
	ignoredErrs []error
}

// NewResultWithError returns a new ResultWithError with the given value and error.
func NewResultWithError[V any](value V, err error) *ResultWithError[V] {
	return &ResultWithError[V]{Value: value, Error: err}
}

// NewResultWithErrorWithIgnoredErrs returns a new ResultWithError with the given value and error.
func NewResultWithErrorWithIgnoredErrs[V any](value V, err error, ignoredErrs ...error) *ResultWithError[V] {
	return &ResultWithError[V]{Value: value, Error: err, ignoredErrs: ignoredErrs}
}

// Destruct returns the value and error of the ResultWithError.
func (w *ResultWithError[V]) Destruct() (V, error) {
	return w.Value, w.Error
}

// Msg returns the error message of the ResultWithError.
func (w *ResultWithError[V]) Msg() string {
	return w.Error.Error()
}

// IsOk returns true if the ResultWithError is not nil.
func (w *ResultWithError[V]) IsOk() bool {
	return w.Error == nil
}

// IsErr returns true if the ResultWithError is nil.
func (w *ResultWithError[V]) IsErr() bool {
	return w.Error != nil && !slices.ContainsWith(w.ignoredErrs, w.Error, func(wErr error, err error) bool {
		return wErr.Error() == err.Error()
	})
}

// Unwrap returns the value of the ResultWithError or panics with the error.
func (w *ResultWithError[V]) Unwrap() V {
	if w.Error != nil {
		panic(w.Error)
	}
	return w.Value
}

// UnwrapOr returns the value of the ResultWithError or the provided default value.
func (w *ResultWithError[V]) UnwrapOr(defaultValue V) V {
	if w.Error != nil {
		return defaultValue
	}
	return w.Value
}

// AddInfoIfErr adds info to the error that is passed in.
func AddInfoIfErr(err error, info ...string) error {
	ret := err
	for i := len(info) - 1; i >= 0; i-- {
		ret = fmt.Errorf("%s: %w", info[i], ret)
	}
	return ret
}

// AddInfoIfErr adds info to the error if Error is not nil. Returns true if the result contained an error.
func (w *ResultWithError[V]) AddInfoIfErr(info ...string) bool {
	if w.IsErr() {
		w.Error = AddInfoIfErr(w.Error, info...)
		return true
	}
	return false
}

// FailTestIfErr fails the test if the ResultWithError is an error.
func (w *ResultWithError[V]) FailTestIfErr(t *testing.T) {
	t.Helper()
	if w.IsErr() {
		t.Errorf("ResultWithError should not be an error: %v", w.Error)
	}
}

// FailTestIfOk fails the test if the ResultWithError is ok.
func (w *ResultWithError[V]) FailTestIfOk(t *testing.T) {
	t.Helper()
	if w.IsOk() {
		t.Errorf("ResultWithError should be an error: %v", w.Error)
	}
}

func (w *ResultWithError[V]) FailTestIfErrf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsErr() {
		t.Errorf(format, args...)
	}
}

func (w *ResultWithError[V]) FailTestIfOkf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsOk() {
		t.Errorf(format, args...)
	}
}
