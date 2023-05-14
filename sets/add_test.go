package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestAdd_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[sets.Set[int], int], sets.Set[int]]{
		{
			Name: "Add to empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{}, B: 1},
			Want: sets.Set[int]{1: true},
		},
		{
			Name: "Add to non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{1: true}, B: 2},
			Want: sets.Set[int]{1: true, 2: true},
		},
		{
			Name: "Add duplicate element",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{1: true}, B: 1},
			Want: sets.Set[int]{1: true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Add()", func(s sets.Set[int], k int) sets.Set[int] {
		sets.Add(&s, k)
		return s
	}, testCases)
}

func TestAdd_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[sets.Set[string], string], sets.Set[string]]{
		{
			Name: "Add to empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[string], string]{A: sets.Set[string]{}, B: "a"},
			Want: sets.Set[string]{"a": true},
		},
		{
			Name: "Add to non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[string], string]{A: sets.Set[string]{"a": true}, B: "b"},
			Want: sets.Set[string]{"a": true, "b": true},
		},
		{
			Name: "Add duplicate element",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[string], string]{A: sets.Set[string]{"a": true}, B: "a"},
			Want: sets.Set[string]{"a": true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Add()", func(s sets.Set[string], k string) sets.Set[string] {
		sets.Add(&s, k)
		return s
	}, testCases)
}
