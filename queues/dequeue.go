package queues

import "github.com/GrewalAS/godash/linkedlist"

// Dequeue removes a value from the front of a queue.
func Dequeue[T any](q *Queue[T]) *T {
	node := linkedlist.Get[T](q.ll, 0)
	linkedlist.Remove(q.ll, 0)
	return node.Value
}
