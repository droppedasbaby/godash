package linkedlist_test

import (
	"testing"

	"github.com/GrewalAS/godash/linkedlist"
	"github.com/GrewalAS/godash/utils"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type customType struct {
		ID   int
		Name string
	}

	e1 := customType{ID: 1, Name: "Alice"}
	e2 := customType{ID: 2, Name: "Bob"}
	e3 := customType{ID: 3, Name: "Charlie"}

	single := []customType{e1}
	singleNodes := []*linkedlist.Node[customType]{
		{Prev: nil, Value: &single[0], Next: nil},
	}
	singleLL := linkedlist.LinkedList[customType]{Head: nil, Tail: nil, Length: 0}
	linkedlist.AddLast[customType](&singleLL, singleNodes[0])

	multiple := []customType{e1, e2, e3}
	multipleNodes := []*linkedlist.Node[customType]{
		{Prev: nil, Value: &multiple[0], Next: nil},
		{Prev: nil, Value: &multiple[1], Next: nil},
		{Prev: nil, Value: &multiple[2], Next: nil},
	}
	multipleLL := linkedlist.LinkedList[customType]{Head: nil, Tail: nil, Length: 0}
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[0])
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[1])
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[2])

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], int], *linkedlist.Node[customType]]{
		{
			Name: "Single-element linked list, get element at index 0",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], int]{
				A: &singleLL,
				B: 0,
			},
			Want: singleNodes[0],
		},
		{
			Name: "Multiple-element linked list, get element at index 0",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], int]{
				A: &multipleLL,
				B: 0,
			},
			Want: multipleNodes[0],
		},
		{
			Name: "Multiple-element linked list, get element at index 1",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], int]{
				A: &multipleLL,
				B: 1,
			},
			Want: multipleNodes[1],
		},
		{
			Name: "Multiple-element linked list, get element at index 2",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], int]{
				A: &multipleLL,
				B: 2,
			},
			Want: multipleNodes[2],
		},
	}

	utils.RunTwoArgumentTestCases(t, "Get()",
		func(ll *linkedlist.LinkedList[customType], index int) *linkedlist.Node[customType] {
			return linkedlist.Get[customType](ll, index)
		}, testCases)
}
