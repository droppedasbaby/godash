package sets_test

import (
	"testing"

	"godash/sets"
	"godash/utils"
)

func TestRemove_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			sets.Set[int], int], sets.Set[int]]{
		{
			Name: "Remove an element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{
				A: sets.Set[int]{1: true, 2: true, 3: true}, B: 2},
			Want: sets.Set[int]{1: true, 3: true},
		},
		{
			Name: "Remove a non-existent element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{
				A: sets.Set[int]{1: true, 2: true}, B: 3},
			Want: sets.Set[int]{1: true, 2: true},
		},
		{
			Name: "Remove an element from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{
				A: sets.Set[int]{}, B: 1,
			},
			Want: sets.Set[int]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(s sets.Set[int], k int) sets.Set[int] {
		sets.Remove[int](&s, k)
		return s
	}, testCases)
}

func TestRemove_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			sets.Set[string], string], sets.Set[string]]{
		{
			Name: "Remove an element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: "b",
			},
			Want: sets.Set[string]{"a": true, "c": true},
		},
		{
			Name: "Remove a non-existent element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{"a": true, "b": true}, B: "c",
			},
			Want: sets.Set[string]{"a": true, "b": true},
		},
		{
			Name: "Remove an element from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{}, B: "a",
			},
			Want: sets.Set[string]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(s sets.Set[string], k string) sets.Set[string] {
		sets.Remove[string](&s, k)
		return s
	}, testCases)
}
