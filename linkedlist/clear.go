package linkedlist

// Clear removes all elements from the linked list. Breaks all connections between nodes and sets the head and tail to
// nil.
func Clear[T any](ll *LinkedList[T]) {
	// Remove all connections between nodes. Prevent memory leaks.
	for curr := ll.Head; curr != nil; {
		next := curr.Next
		curr.Next = nil
		curr.Prev = nil
		curr = next
	}
	ll.Head = nil
	ll.Tail = nil
	ll.Length = 0
}
