package linkedlist

// Insert inserts the element at the specified index. Index 0 is the head of the list.
func Insert[T any](ll *LinkedList[T], index int, n *Node[T]) {
	if index < 0 || index > ll.Length {
		panic("index out of range")
	}

	if index == 0 {
		AddFirst(ll, n)
	} else if index == ll.Length {
		AddLast(ll, n)
	} else {
		prev := Get(ll, index-1)
		next := prev.Next
		prev.Next = n
		n.Prev = prev
		n.Next = next
		next.Prev = n
		ll.Length++
	}
}