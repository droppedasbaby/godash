package linkedlist_test

import (
	"testing"

	"godash/linkedlist"
	"godash/utils"
)

func TestContains(t *testing.T) {
	t.Parallel()

	emptyLL := linkedlist.LinkedList[int]{}

	single := 0
	singleNodes := []*linkedlist.Node[int]{
		{Prev: nil, Value: &single, Next: nil},
	}
	singleLL := linkedlist.LinkedList[int]{}
	linkedlist.AddLast(&singleLL, singleNodes[0])

	multiple := []int{0, 1, 2}
	multipleNodes := []*linkedlist.Node[int]{
		{Prev: nil, Value: &multiple[0], Next: nil},
		{Prev: nil, Value: &multiple[1], Next: nil},
		{Prev: nil, Value: &multiple[2], Next: nil},
	}
	multipleLL := linkedlist.LinkedList[int]{}
	linkedlist.AddLast(&multipleLL, multipleNodes[0])
	linkedlist.AddLast(&multipleLL, multipleNodes[1])
	linkedlist.AddLast(&multipleLL, multipleNodes[2])

	nonExistent := 3

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int], bool]{
		{
			Name: "Empty linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int]{
				A: &emptyLL,
				B: &single,
			},
			Want: false,
		},
		{
			Name: "Single-element linked list, contains element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int]{
				A: &singleLL,
				B: &single,
			},
			Want: true,
		},
		{
			Name: "Single-element linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int]{
				A: &singleLL,
				B: &nonExistent,
			},
			Want: false,
		},
		{
			Name: "Multiple-element linked list, contains element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int]{
				A: &multipleLL,
				B: &multiple[1],
			},
			Want: true,
		},
		{
			Name: "Multiple-element linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[int], *int]{
				A: &multipleLL,
				B: &nonExistent,
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()",
		func(ll *linkedlist.LinkedList[int], e *int) bool {
			return linkedlist.Contains[int](ll, e)
		}, testCases)
}

func TestContains_CustomType(t *testing.T) {
	t.Parallel()

	type customType struct {
		ID   int
		Name string
	}

	customComparator := func(a, b *customType) bool {
		return a.ID == b.ID && a.Name == b.Name
	}

	e1 := customType{ID: 1, Name: "Alice"}
	e2 := customType{ID: 2, Name: "Bob"}
	e3 := customType{ID: 3, Name: "Charlie"}
	nonExistent := customType{ID: 4, Name: "David"}

	emptyLL := linkedlist.LinkedList[customType]{}

	single := []customType{e1}
	singleNodes := []*linkedlist.Node[customType]{
		{Prev: nil, Value: &single[0], Next: nil},
	}
	singleLL := linkedlist.LinkedList[customType]{}
	linkedlist.AddLast[customType](&singleLL, singleNodes[0])

	multiple := []customType{e1, e2, e3}
	multipleNodes := []*linkedlist.Node[customType]{
		{Prev: nil, Value: &multiple[0], Next: nil},
		{Prev: nil, Value: &multiple[1], Next: nil},
		{Prev: nil, Value: &multiple[2], Next: nil},
	}
	multipleLL := linkedlist.LinkedList[customType]{}
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[0])
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[1])
	linkedlist.AddLast[customType](&multipleLL, multipleNodes[2])

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType], bool]{
		{
			Name: "Empty linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType]{
				A: &emptyLL,
				B: &e1,
			},
			Want: false,
		},
		{
			Name: "Single-element linked list, contains element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType]{
				A: &singleLL,
				B: &e1,
			},
			Want: true,
		},
		{
			Name: "Single-element linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType]{
				A: &singleLL,
				B: &nonExistent,
			},
			Want: false,
		},
		{
			Name: "Multiple-element linked list, contains element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType]{
				A: &multipleLL,
				B: &e2,
			},
			Want: true,
		},
		{
			Name: "Multiple-element linked list, does not contain element",
			Args: utils.TwoArgumentTestCasesArgsType[*linkedlist.LinkedList[customType], *customType]{
				A: &multipleLL,
				B: &nonExistent,
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "ContainsWith()",
		func(ll *linkedlist.LinkedList[customType], e *customType) bool {
			return linkedlist.ContainsWith[customType](ll, e, customComparator)
		}, testCases)
}
