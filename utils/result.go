package utils

import (
	"fmt"
	"testing"

	"github.com/GrewalAS/godash/slices"
)

type Result[V any] struct {
	Value       V
	Error       error
	ignoredErrs []error
}

// NewResult returns a new Result with the given value and error.
func NewResult[V any](value V, err error) *Result[V] {
	return &Result[V]{Value: value, Error: err, ignoredErrs: []error{}}
}

// NewResultWithIgnoredErrs returns a new Result with the given value and error.
func NewResultWithIgnoredErrs[V any](value V, err error, ignoredErrs ...error) *Result[V] {
	return &Result[V]{Value: value, Error: err, ignoredErrs: ignoredErrs}
}

// Destruct returns the value and error of the Result.
func (w *Result[V]) Destruct() (V, error) {
	return w.Value, w.Error
}

// Msg returns the error message of the Result.
func (w *Result[V]) Msg() string {
	return w.Error.Error()
}

// IsOk returns true if the Result is not nil.
func (w *Result[V]) IsOk() bool {
	return w.Error == nil
}

// IsErr returns true if the Result is nil.
func (w *Result[V]) IsErr() bool {
	return w.Error != nil && !slices.ContainsWith(w.ignoredErrs, w.Error, func(wErr error, err error) bool {
		return wErr.Error() == err.Error()
	})
}

// Unwrap returns the value of the Result or panics with the error.
func (w *Result[V]) Unwrap() V {
	if w.Error != nil {
		panic(w.Error)
	}
	return w.Value
}

// UnwrapOr returns the value of the Result or the provided default value.
func (w *Result[V]) UnwrapOr(defaultValue V) V {
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
func (w *Result[V]) AddInfoIfErr(info ...string) bool {
	if w.IsErr() {
		w.Error = AddInfoIfErr(w.Error, info...)
		return true
	}
	return false
}

// FailTestIfErr fails the test if the Result is an error.
func (w *Result[V]) FailTestIfErr(t *testing.T) {
	t.Helper()
	if w.IsErr() {
		t.Errorf("Result should not be an error: %v", w.Error)
	}
}

// FailTestIfOk fails the test if the Result is ok.
func (w *Result[V]) FailTestIfOk(t *testing.T) {
	t.Helper()
	if w.IsOk() {
		t.Errorf("Result should be an error: %v", w.Error)
	}
}

func (w *Result[V]) FailTestIfErrf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsErr() {
		t.Errorf(format, args...)
	}
}

func (w *Result[V]) FailTestIfOkf(t *testing.T, format string, args ...interface{}) {
	t.Helper()
	if w.IsOk() {
		t.Errorf(format, args...)
	}
}
