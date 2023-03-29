package linkedlist

// Get returns the node at the specified index.
func Get[T any](ll *LinkedList[T], index int) *Node[T] {
	if index < 0 || index >= ll.Length {
		panic("index out of range")
	}

	n := ll.Head
	for i := 0; i < index; i++ {
		n = n.Next
	}
	return n
}
