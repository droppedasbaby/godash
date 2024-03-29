package linkedlist

// Contains returns true if the linked list contains the element as a value in a node.
func Contains[T comparable](ll *LinkedList[T], e *T) bool {
	for curr := ll.Head; curr != nil; curr = curr.Next {
		if curr.Value == e {
			return true
		}
	}
	return false
}

// ContainsWith returns true if the linked list contains the element as a value in a node using the provided comparator.
func ContainsWith[T any](ll *LinkedList[T], e *T, p func(*T, *T) bool) bool {
	for curr := ll.Head; curr != nil; curr = curr.Next {
		if p(curr.Value, e) {
			return true
		}
	}
	return false
}
