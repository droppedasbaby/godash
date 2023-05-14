// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
package linkedlist

// AddFirst adds a node to the beginning of the linked list.
func AddFirst[T any](ll *LinkedList[T], n *Node[T]) {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Head.Prev = n
		n.Next = ll.Head
		ll.Head = n
	}
	ll.Length++
}

// AddLast adds a node to the end of the linked list.
func AddLast[T any](ll *LinkedList[T], n *Node[T]) {
	if ll.Head == nil {
		ll.Head = n
		ll.Tail = n
	} else {
		ll.Tail.Next = n
		n.Prev = ll.Tail
		ll.Tail = n
	}
	ll.Length++
}
