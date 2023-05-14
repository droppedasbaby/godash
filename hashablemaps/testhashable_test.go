// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

type TestHashable struct {
	Value string
}

func (t TestHashable) Hash() string {
	return t.Value
}
