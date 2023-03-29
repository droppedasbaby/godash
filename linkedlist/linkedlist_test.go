package linkedlist_test

import (
	"testing"

	"godash/linkedlist"
	"godash/utils"
)

func TestFromSlice(t *testing.T) {
	t.Parallel()

	single := []int{0}
	multiple := []int{0, 1, 2}

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]int], linkedlist.LinkedList[int]]{
		{
			Name: "From empty slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{}},
			Want: linkedlist.LinkedList[int]{Head: nil, Tail: nil},
		},
		{
			Name: "From single element slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: single},
			Want: linkedlist.LinkedList[int]{
				Head: &linkedlist.Node[int]{Prev: nil, Value: &single[0], Next: nil},
				Tail: &linkedlist.Node[int]{Prev: nil, Value: &single[0], Next: nil},
			},
		},
		{
			Name: "From multiple elements slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: multiple},
			Want: func() linkedlist.LinkedList[int] {
				nodes := []*linkedlist.Node[int]{
					{Value: &multiple[0]},
					{Value: &multiple[1]},
					{Value: &multiple[2]},
				}
				nodes[0].Next = nodes[1]
				nodes[1].Prev, nodes[1].Next = nodes[0], nodes[2]
				nodes[2].Prev = nodes[1]
				return linkedlist.LinkedList[int]{Head: nodes[0], Tail: nodes[2]}
			}(),
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSlice()", linkedlist.FromSlice[int], testCases)
}

func TestToSlice(t *testing.T) {
	t.Parallel()

	single := []int{0}
	multiple := []int{0, 1, 2}

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[linkedlist.LinkedList[int]], []int]{
		{
			Name: "To slice from empty linked list",
			Args: utils.SingleArgumentTestCasesArgsType[linkedlist.LinkedList[int]]{
				A: linkedlist.LinkedList[int]{Head: nil, Tail: nil},
			},
			Want: []int{},
		},
		{
			Name: "To slice from single element linked list",
			Args: utils.SingleArgumentTestCasesArgsType[linkedlist.LinkedList[int]]{
				A: func() linkedlist.LinkedList[int] {
					nodes := []*linkedlist.Node[int]{
						{Value: &single[0]},
					}
					return linkedlist.LinkedList[int]{Head: nodes[0], Tail: nodes[0]}
				}(),
			},
			Want: single,
		},
		{
			Name: "To slice from multiple elements linked list",
			Args: utils.SingleArgumentTestCasesArgsType[linkedlist.LinkedList[int]]{
				A: func() linkedlist.LinkedList[int] {
					nodes := []*linkedlist.Node[int]{
						{Value: &multiple[0]},
						{Value: &multiple[1]},
						{Value: &multiple[2]},
					}
					nodes[0].Next = nodes[1]
					nodes[1].Prev, nodes[1].Next = nodes[0], nodes[2]
					nodes[2].Prev = nodes[1]
					return linkedlist.LinkedList[int]{Head: nodes[0], Tail: nodes[2]}
				}(),
			},
			Want: multiple,
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSlice()", func(ll linkedlist.LinkedList[int]) []int {
		return linkedlist.ToSlice[int](ll)
	}, testCases)
}
