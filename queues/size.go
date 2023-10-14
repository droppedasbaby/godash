package queues

// Size returns the number of elements in the queue.
func Size[T any](q *Queue[T]) int {
	return q.ll.Length
}
