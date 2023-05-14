// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
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
func ContainsWith[T any](ll *LinkedList[T], e *T, comp func(*T, *T) bool) bool {
	for curr := ll.Head; curr != nil; curr = curr.Next {
		if comp(curr.Value, e) {
			return true
		}
	}
	return false
}
