// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices_test

type testStruct struct {
	x int
	y int
}

type testHashable struct {
	Value string
}

func (t testHashable) Hash() string {
	return t.Value
}

func ptr[T any](t T) *T {
	return &t
}
