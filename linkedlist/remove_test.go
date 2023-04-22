package linkedlist_test

import (
	"testing"

	"godash/linkedlist"
	"godash/utils"
)

func TestRemoveInt(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int],
		[]int]{
		{
			Name: "Remove from a one-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1}),
				B: 0,
			},
			Want: []int{},
		},
		{
			Name: "Remove from a two-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1, 2}),
				B: 0,
			},
			Want: []int{2},
		},
		{
			Name: "Remove from a two-element list, tail",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1, 2}),
				B: 1,
			},
			Want: []int{1},
		},
		{
			Name: "Remove from a three-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 0,
			},
			Want: []int{2, 3},
		},
		{
			Name: "Remove from a three-element list, middle",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 1,
			},
			Want: []int{1, 3},
		},
		{
			Name: "Remove from a three-element list, tail",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[int], int]{
				A: linkedlist.FromSlice([]int{1, 2, 3}),
				B: 2,
			},
			Want: []int{1, 2},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(ll linkedlist.LinkedList[int], index int) []int {
		linkedlist.Remove(&ll, index)
		return linkedlist.ToSlice(ll)
	}, testCases)
}

func TestRemove_TestType(t *testing.T) {
	t.Parallel()

	type testType struct {
		ID int
	}

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int],
		[]testType]{
		{
			Name: "Remove from a one-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}}),
				B: 0,
			},
			Want: []testType{},
		},
		{
			Name: "Remove from a two-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}, {2}}),
				B: 0,
			},
			Want: []testType{{2}},
		},
		{
			Name: "Remove from a two-element list, tail",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}, {2}}),
				B: 1,
			},
			Want: []testType{{1}},
		},
		{
			Name: "Remove from a three-element list, head",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}, {2}, {3}}),
				B: 0,
			},
			Want: []testType{{2}, {3}},
		},
		{
			Name: "Remove from a three-element list, middle",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}, {2}, {3}}),
				B: 1,
			},
			Want: []testType{{1}, {3}},
		},
		{
			Name: "Remove from a three-element list, tail",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[testType], int]{
				A: linkedlist.FromSlice([]testType{{1}, {2}, {3}}),
				B: 2,
			},
			Want: []testType{{1}, {2}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(ll linkedlist.LinkedList[testType], index int) []testType {
		linkedlist.Remove(&ll, index)
		return linkedlist.ToSlice(ll)
	}, testCases)
}

func TestRemovePrevNextPointers(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[string], int],
		[]prevNext]{
		{
			Name: "remove from the middle",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[string], int]{
				A: linkedlist.FromSlice([]string{"a", "b", "c", "d", "e"}),
				B: 2,
			},
			Want: []prevNext{
				{Prev: nil, Next: pointerToString("b")},
				{Prev: pointerToString("a"), Next: pointerToString("d")},
				{Prev: pointerToString("b"), Next: pointerToString("e")},
				{Prev: pointerToString("d"), Next: nil},
			},
		},
		{
			Name: "remove from the beginning",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[string], int]{
				A: linkedlist.FromSlice([]string{"a", "b", "c"}),
				B: 0,
			},
			Want: []prevNext{
				{Prev: nil, Next: pointerToString("c")},
				{Prev: pointerToString("b"), Next: nil},
			},
		},
		{
			Name: "remove from the end",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[string], int]{
				A: linkedlist.FromSlice([]string{"a", "b", "c"}),
				B: 2,
			},
			Want: []prevNext{
				{Prev: nil, Next: pointerToString("b")},
				{Prev: pointerToString("a"), Next: nil},
			},
		},
		{
			Name: "remove single element",
			Args: utils.TwoArgumentTestCasesArgsType[linkedlist.LinkedList[string], int]{
				A: linkedlist.FromSlice([]string{"a"}),
				B: 0,
			},
			Want: []prevNext{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(ll linkedlist.LinkedList[string], index int) []prevNext {
		linkedlist.Remove(&ll, index)
		return getPrevNext(&ll)
	}, testCases)
}
