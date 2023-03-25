package sets_test

import (
	"testing"

	"godash/sets"
	"godash/utils"
)

func TestGet_WithInt(t *testing.T) {
	t.Parallel()

	two := 2
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[int], int], *int]{
		{
			Name: "Get element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{A: sets.Set[int]{1: true, 2: true, 3: true}, B: 2},
			Want: &two,
		},
		{
			Name: "Get non-existent element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{A: sets.Set[int]{1: true, 2: true}, B: 3},
			Want: nil,
		},
		{
			Name: "Get an element from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], int]{A: sets.Set[int]{}, B: 1},
			Want: nil,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Get()", sets.Get[int], testCases)
}

func TestGet_WithString(t *testing.T) {
	t.Parallel()

	b := "b"
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], string], *string]{
		{
			Name: "Get element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{A: sets.Set[string]{"a": true, "b": true, "c": true}, B: "b"},
			Want: &b,
		},
		{
			Name: "Get non-existent element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{A: sets.Set[string]{"a": true, "b": true}, B: "c"},
			Want: nil,
		},
		{
			Name: "Get an element from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{A: sets.Set[string]{}, B: "a"},
			Want: nil,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Get()", sets.Get[string], testCases)
}
