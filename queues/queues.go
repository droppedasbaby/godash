package queues

import "github.com/GrewalAS/godash/linkedlist"

type Queue[T any] struct {
	ll *linkedlist.LinkedList[T]
}
