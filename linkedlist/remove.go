// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
package linkedlist

// Remove removes the element at the specified index. Index 0 is the head of the list.
func Remove[T any](ll *LinkedList[T], index int) {
	if index < 0 || index >= ll.Length {
		panic("index out of range")
	}

	switch index {
	case 0:
		ll.Head = ll.Head.Next
		if ll.Head != nil {
			ll.Head.Prev = nil
		}
	case ll.Length - 1:
		ll.Tail = ll.Tail.Prev
		if ll.Tail != nil {
			ll.Tail.Next = nil
		}
	default:
		prev := Get(ll, index-1)
		curr := prev.Next
		next := prev.Next.Next
		prev.Next = next
		next.Prev = prev
		curr.Next = nil
		curr.Prev = nil
	}
	ll.Length--
}
