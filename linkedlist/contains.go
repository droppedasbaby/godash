package linkedlist

// Contains returns true if the linked list contains the element as a value in a node.
func Contains[T any](ll *LinkedList[T], e *T) bool {
	for curr := ll.Head; curr != nil; curr = curr.Next {
		if curr.Value == e {
			return true
		}
	}
	return false
}