package linkedlist

// Remove removes the element at the specified index. Index 0 is the head of the list.
func Remove[T any](ll *LinkedList[T], index int) {
	if index < 0 || index >= ll.Length {
		panic("index out of range")
	}

	switch index {
	case 0:
		ll.Head = ll.Head.Next
	case ll.Length - 1:
		ll.Tail = ll.Tail.Prev
	default:
		prev := Get(ll, index-1)
		next := prev.Next.Next
		prev.Next = next
		next.Prev = prev
	}
	ll.Length--
}
