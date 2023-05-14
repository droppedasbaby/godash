package linkedlist_test

import "github.com/GrewalAS/godash/linkedlist"

type prevNext struct {
	Prev, Next *string
}

func pointerToString(s string) *string {
	return &s
}

func getPrevNext(ll *linkedlist.LinkedList[string]) (result []prevNext) {
	pointers := []prevNext{}
	node := ll.Head
	i := 0
	for node != nil {
		switch {
		case node.Prev == nil && node.Next == nil:
			pointers = append(pointers, prevNext{Prev: nil, Next: nil})
		case node.Prev == nil && node.Next != nil:
			pointers = append(pointers, prevNext{Prev: nil, Next: node.Next.Value})
		case node.Prev != nil && node.Next == nil:
			pointers = append(pointers, prevNext{Prev: node.Prev.Value, Next: nil})
		default:
			pointers = append(pointers, prevNext{Prev: node.Prev.Value, Next: node.Next.Value})
		}
		node = node.Next
		i++
	}
	return pointers
}
