package queues

import "github.com/GrewalAS/godash/linkedlist"

// Peek returns the first element of the queue without removing it.
func Peek[T any](q *Queue[T]) *T {
	return linkedlist.Get(q.ll, 0).Value
}
