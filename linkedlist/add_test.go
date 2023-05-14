// Package linkedlist provides functionality for creating and manipulating
// doubly linked lists in Go.
package linkedlist_test

import (
	"testing"

	"github.com/GrewalAS/godash/linkedlist"
	"github.com/GrewalAS/godash/utils"
)

func TestAddFirst(t *testing.T) {
	t.Parallel()

	single := []int{0}
	double := []int{0, 1}
	multiple := []int{0, 1, 2}

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]], []int]{
		{
			Name: "Add first node to empty linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: &linkedlist.LinkedList[int]{Head: nil, Tail: nil},
				B: &linkedlist.Node[int]{Prev: nil, Value: &single[0], Next: nil},
			},
			Want: single,
		},
		{
			Name: "Add first node to single element linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: func() *linkedlist.LinkedList[int] {
					nodes := []*linkedlist.Node[int]{
						{Value: &double[1]},
					}
					return &linkedlist.LinkedList[int]{Head: nodes[0], Tail: nodes[0]}
				}(),
				B: &linkedlist.Node[int]{Prev: nil, Value: &double[0], Next: nil},
			},
			Want: double,
		},
		{
			Name: "Add first node to multiple elements linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: func() *linkedlist.LinkedList[int] {
					ll := &linkedlist.LinkedList[int]{}
					nodes := []*linkedlist.Node[int]{
						{Value: &multiple[1]},
						{Value: &multiple[2]},
					}
					nodes[0].Next = nodes[1]
					nodes[1].Prev = nodes[0]
					ll.Head, ll.Tail = nodes[0], nodes[1]
					return ll
				}(),
				B: &linkedlist.Node[int]{Prev: nil, Value: &multiple[0], Next: nil},
			},
			Want: multiple,
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddFirst()",
		func(ll *linkedlist.LinkedList[int], n *linkedlist.Node[int]) []int {
			linkedlist.AddFirst[int](ll, n)
			return linkedlist.ToSlice[int](*ll)
		}, testCases)
}

func TestAddLast(t *testing.T) {
	t.Parallel()

	single := []int{0}
	double := []int{0, 1}
	multiple := []int{0, 1, 2}

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]], []int]{
		{
			Name: "Add last node to empty linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: &linkedlist.LinkedList[int]{Head: nil, Tail: nil},
				B: &linkedlist.Node[int]{Prev: nil, Value: &single[0], Next: nil},
			},
			Want: single,
		},
		{
			Name: "Add last node to single element linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: func() *linkedlist.LinkedList[int] {
					nodes := []*linkedlist.Node[int]{
						{Value: &double[0]},
					}
					return &linkedlist.LinkedList[int]{Head: nodes[0], Tail: nodes[0]}
				}(),
				B: &linkedlist.Node[int]{Prev: nil, Value: &double[1], Next: nil},
			},
			Want: double,
		},
		{
			Name: "Add last node to multiple elements linked list",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *linkedlist.Node[int]]{
				A: func() *linkedlist.LinkedList[int] {
					ll := &linkedlist.LinkedList[int]{}
					nodes := []*linkedlist.Node[int]{
						{Value: &multiple[0]},
						{Value: &multiple[1]},
					}
					nodes[0].Next = nodes[1]
					nodes[1].Prev = nodes[0]
					ll.Head, ll.Tail = nodes[0], nodes[1]
					return ll
				}(),
				B: &linkedlist.Node[int]{Prev: nil, Value: &multiple[2], Next: nil},
			},
			Want: multiple,
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddLast()",
		func(ll *linkedlist.LinkedList[int], n *linkedlist.Node[int]) []int {
			linkedlist.AddLast[int](ll, n)
			return linkedlist.ToSlice[int](*ll)
		}, testCases)
}
