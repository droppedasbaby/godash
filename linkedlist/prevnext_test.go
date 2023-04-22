package linkedlist_test

import "godash/linkedlist"

type prevNext struct {
	Prev, Next *string
}

func pointerToString(s string) *string {
	return &s
}

func GetPrevNext(ll *linkedlist.LinkedList[string]) (result []prevNext) {
	pointers := []prevNext{}
	node := ll.Head
	i := 0
	for node != nil {
		if node.Prev == nil && node.Next == nil {
			pointers = append(pointers, prevNext{})
		} else if node.Prev == nil && node.Next != nil {
			pointers = append(pointers, prevNext{Prev: nil, Next: node.Next.Value})
		} else if node.Prev != nil && node.Next == nil {
			pointers = append(pointers, prevNext{Prev: node.Prev.Value, Next: nil})
		} else {
			pointers = append(pointers, prevNext{Prev: node.Prev.Value, Next: node.Next.Value})
		}
		node = node.Next
		i++
	}
	return pointers
}
