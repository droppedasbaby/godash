package linkedlist_test

import (
	"testing"

	"godash/linkedlist"
	"godash/utils"
)

func TestClear(t *testing.T) {
	t.Parallel()

	single := []int{0}
	multiple := []int{0, 1, 2}

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[*linkedlist.LinkedList[int]], bool]{
		{
			Name: "Clear empty linked list",
			Args: utils.SingleArgumentTestCasesArgsType[*linkedlist.LinkedList[int]]{
				A: &linkedlist.LinkedList[int]{Head: nil, Tail: nil},
			},
			Want: true,
		},
		{
			Name: "Clear single-element linked list",
			Args: utils.SingleArgumentTestCasesArgsType[*linkedlist.LinkedList[int]]{
				A: func() *linkedlist.LinkedList[int] {
					node := &linkedlist.Node[int]{Prev: nil, Value: &single[0], Next: nil}
					return &linkedlist.LinkedList[int]{Head: node, Tail: node}
				}(),
			},
			Want: true,
		},
		{
			Name: "Clear multiple-element linked list",
			Args: utils.SingleArgumentTestCasesArgsType[*linkedlist.LinkedList[int]]{
				A: func() *linkedlist.LinkedList[int] {
					ll := &linkedlist.LinkedList[int]{}
					nodes := []*linkedlist.Node[int]{
						{Value: &multiple[0]},
						{Value: &multiple[1]},
						{Value: &multiple[2]},
					}
					nodes[0].Next = nodes[1]
					nodes[1].Prev = nodes[0]
					nodes[1].Next = nodes[2]
					nodes[2].Prev = nodes[1]
					ll.Head, ll.Tail = nodes[0], nodes[2]
					return ll
				}(),
			},
			Want: true,
		},
	}

	utils.RunSingleArgumentTestCases(t, "Clear()", func(ll *linkedlist.LinkedList[int]) bool {
		linkedlist.Clear[int](ll)
		return ll.Head == nil && ll.Tail == nil && ll.Length == 0
	}, testCases)
}

func TestClearConnectionBreaks(t *testing.T) {
	t.Parallel()

	values := []int{0, 1, 2}
	nodes := []*linkedlist.Node[int]{
		{Value: &values[0]},
		{Value: &values[1]},
		{Value: &values[2]},
	}

	ll := &linkedlist.LinkedList[int]{}
	for _, node := range nodes {
		linkedlist.AddLast[int](ll, node)
	}
	linkedlist.Clear[int](ll)

	if nodes[0].Next != nil || nodes[1].Prev != nil || nodes[1].Next != nil || nodes[2].Prev != nil {
		t.Errorf("Clear() did not break connections between nodes")
	}
}
