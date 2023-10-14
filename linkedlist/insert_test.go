package linkedlist_test

import (
	"testing"

	"github.com/GrewalAS/godash/linkedlist"
	"github.com/GrewalAS/godash/utils"
)

func TestInsert(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[int], int, *linkedlist.Node[int]],
		linkedlist.LinkedList[int]]{
		{
			Name: "Insert at head",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[int], int, *linkedlist.Node[int]]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 0,
				C: &linkedlist.Node[int]{Value: new(int), Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]int{0, 1, 2, 3}),
		},
		{
			Name: "Insert at tail",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[int], int, *linkedlist.Node[int]]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 3,
				C: &linkedlist.Node[int]{Value: new(int), Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]int{1, 2, 3, 0}),
		},
		{
			Name: "Insert at middle",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[int], int, *linkedlist.Node[int]]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 1,
				C: &linkedlist.Node[int]{Value: new(int), Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]int{1, 0, 2, 3}),
		},
		{
			Name: "Insert into empty list",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[int], int, *linkedlist.Node[int]]{
				A: linkedlist.LinkedList[int]{Head: nil, Tail: nil, Length: 0},
				B: 0,
				C: &linkedlist.Node[int]{Value: new(int), Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]int{0}),
		},
	}

	utils.RunThreeArgumentTestCases(t, "Insert()",
		func(ll linkedlist.LinkedList[int], index int, n *linkedlist.Node[int]) linkedlist.LinkedList[int] {
			linkedlist.Insert(&ll, index, n)
			return ll
		}, testCases)
}

func TestInsert_TestStruct(t *testing.T) {
	t.Parallel()

	type testStruct struct {
		ID   int
		Name string
	}

	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[testStruct], int, *linkedlist.Node[testStruct]],
		linkedlist.LinkedList[testStruct]]{
		{
			Name: "Insert at head",
			Args: utils.ThreeArgumentTestCasesArgsType[
				linkedlist.LinkedList[testStruct], int, *linkedlist.Node[testStruct]]{
				A: linkedlist.FromSlice([]testStruct{
					{1, "One"}, {2, "Two"}, {3, "Three"},
				}),
				B: 0,
				C: &linkedlist.Node[testStruct]{Value: &testStruct{0, "Zero"}, Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice(
				[]testStruct{
					{0, "Zero"}, {1, "One"}, {2, "Two"}, {3, "Three"},
				}),
		},
		{
			Name: "Insert at tail",
			Args: utils.ThreeArgumentTestCasesArgsType[
				linkedlist.LinkedList[testStruct], int, *linkedlist.Node[testStruct]]{
				A: linkedlist.FromSlice([]testStruct{{1, "One"}, {2, "Two"}, {3, "Three"}}),
				B: 3,
				C: &linkedlist.Node[testStruct]{Value: &testStruct{4, "Four"}, Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]testStruct{
				{1, "One"}, {2, "Two"}, {3, "Three"}, {4, "Four"},
			}),
		},
		{
			Name: "Insert at middle",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[testStruct], int, *linkedlist.Node[testStruct]]{
				A: linkedlist.FromSlice([]testStruct{{1, "One"}, {2, "Two"}, {3, "Three"}}),
				B: 1,
				C: &linkedlist.Node[testStruct]{Value: &testStruct{0, "Zero"}, Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]testStruct{
				{1, "One"}, {0, "Zero"}, {2, "Two"}, {3, "Three"},
			}),
		},
		{
			Name: "Insert into empty list",
			Args: utils.ThreeArgumentTestCasesArgsType[linkedlist.LinkedList[testStruct], int, *linkedlist.Node[testStruct]]{
				A: linkedlist.LinkedList[testStruct]{Head: nil, Tail: nil, Length: 0},
				B: 0,
				C: &linkedlist.Node[testStruct]{Value: &testStruct{0, "Zero"}, Prev: nil, Next: nil},
			},
			Want: linkedlist.FromSlice([]testStruct{{0, "Zero"}}),
		},
	}

	utils.RunThreeArgumentTestCases(t, "Insert()",
		func(
			ll linkedlist.LinkedList[testStruct],
			index int,
			n *linkedlist.Node[testStruct]) linkedlist.LinkedList[testStruct] {
			linkedlist.Insert(&ll, index, n)
			return ll
		}, testCases)
}
