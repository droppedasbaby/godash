package sets_test

import (
	"testing"

	"godash/sets"
	"godash/utils"
)

func TestAddAll_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[sets.Set[int], []int], sets.Set[int]]{
		{
			Name: "Add multiple elements to empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], []int]{A: sets.Set[int]{}, B: []int{1, 2, 3}},
			Want: sets.Set[int]{1: true, 2: true, 3: true},
		},
		{
			Name: "Add multiple elements to non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], []int]{A: sets.Set[int]{1: true}, B: []int{2, 3}},
			Want: sets.Set[int]{1: true, 2: true, 3: true},
		},
		{
			Name: "Add duplicate and new elements",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], []int]{A: sets.Set[int]{1: true, 2: true}, B: []int{2, 3}},
			Want: sets.Set[int]{1: true, 2: true, 3: true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()", func(s sets.Set[int], ks []int) sets.Set[int] {
		sets.AddAll(&s, ks...)
		return s
	}, testCases)
}

func TestAddAll_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], []string], sets.Set[string]]{
		{
			Name: "Add multiple elements to empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{
				A: sets.Set[string]{}, B: []string{"a", "b", "c"},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true},
		},
		{
			Name: "Add multiple elements to non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{
				A: sets.Set[string]{"a": true}, B: []string{"b", "c"},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true},
		},
		{
			Name: "Add duplicate and new elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{
				A: sets.Set[string]{"a": true, "b": true}, B: []string{"b", "c"},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()", func(s sets.Set[string], ks []string) sets.Set[string] {
		sets.AddAll(&s, ks...)
		return s
	}, testCases)
}
