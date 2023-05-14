// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
package linkedlist

// Node is a node in a doubly linked list.
type Node[T any] struct {
	Prev  *Node[T]
	Value *T
	Next  *Node[T]
}

// LinkedList is a doubly linked list.
type LinkedList[T any] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

// FromSlice returns a linked list from a slice/array.
func FromSlice[T any](s []T) (ll LinkedList[T]) {
	ll = LinkedList[T]{Head: nil, Tail: nil}

	for _, v := range s {
		val := v
		AddLast(&ll, &Node[T]{Prev: nil, Value: &val, Next: nil})
	}
	return
}

// ToSlice returns a slice from a linked list.
func ToSlice[T any](ll LinkedList[T]) (s []T) {
	s = []T{}
	for n := ll.Head; n != nil; n = n.Next {
		s = append(s, *n.Value)
	}
	return
}
