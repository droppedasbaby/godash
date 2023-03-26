package linkedlist

// Node is a node in a doubly linked list.
type Node[T any] struct {
	Prev  *Node[T]
	Value T
	Next  *Node[T]
}

// LinkedList is a doubly linked list.
type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
}

// FromSlice returns a linked list from a slice/array.

// ToSlice returns a slice from a linked list.
