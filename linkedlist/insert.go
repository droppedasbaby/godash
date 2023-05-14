// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
package linkedlist

// Insert inserts the element at the specified index. Index 0 is the head of the list.
func Insert[T any](ll *LinkedList[T], index int, n *Node[T]) {
	if index < 0 || index > ll.Length {
		panic("index out of range")
	}

	switch index {
	case 0:
		AddFirst(ll, n)
	case ll.Length:
		AddLast(ll, n)
	default:
		prev := Get(ll, index-1)
		next := prev.Next
		prev.Next = n
		n.Prev = prev
		n.Next = next
		next.Prev = n
		ll.Length++
	}
}
