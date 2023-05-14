// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

type TestHashable struct {
	Value string
}

func (t TestHashable) Hash() string {
	return t.Value
}
