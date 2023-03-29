package linkedlist

// Clear removes all elements from the linked list.
func Clear[T any](ll *LinkedList[T]) {
	ll.Head = nil
	ll.Tail = nil
	ll.Length = 0
}
