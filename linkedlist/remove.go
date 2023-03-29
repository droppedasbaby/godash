package linkedlist

// Remove removes the element at the specified index. Index 0 is the head of the list.
func Remove[T any](ll *LinkedList[T], index int) {
	if index < 0 || index >= ll.Length {
		panic("index out of range")
	}

	if index == 0 {
		ll.Head = ll.Head.Next
	} else if index == ll.Length-1 {
		ll.Tail = ll.Tail.Prev
	} else {
		prev := Get(ll, index-1)
		next := prev.Next.Next
		prev.Next = next
		next.Prev = prev
	}
	ll.Length--
}
