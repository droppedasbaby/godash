package queues

import "github.com/GrewalAS/godash/linkedlist"

// Enqueue adds a value to the end of a queue.
func Enqueue[T any](q *Queue[T], value *T) {
	node := &linkedlist.Node[T]{Prev: nil, Value: value, Next: nil}
	linkedlist.AddLast[T](q.ll, node)
}
