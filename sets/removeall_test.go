package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestRemoveAll_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[int], []int], sets.Set[int]]{
		{
			Name: "Remove multiple elements from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], []int]{A: sets.Set[int]{1: true, 2: true, 3: true, 4: true}, B: []int{1, 3}},
			Want: sets.Set[int]{2: true, 4: true},
		},
		{
			Name: "Remove non-existent elements from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], []int]{A: sets.Set[int]{1: true, 2: true, 3: true}, B: []int{4, 5}},
			Want: sets.Set[int]{1: true, 2: true, 3: true},
		},
		{
			Name: "Remove elements from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], []int]{A: sets.Set[int]{}, B: []int{1, 2}},
			Want: sets.Set[int]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()", func(s sets.Set[int], keys []int) sets.Set[int] {
		sets.RemoveAll(&s, keys...)
		return s
	}, testCases)
}

func TestRemoveAll_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], []string], sets.Set[string]]{
		{
			Name: "Remove multiple elements from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{
				A: sets.Set[string]{"a": true, "b": true, "c": true, "d": true}, B: []string{"a", "c"},
			},
			Want: sets.Set[string]{"b": true, "d": true},
		},
		{
			Name: "Remove non-existent elements from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: []string{"x", "y"},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true},
		},
		{
			Name: "Remove elements from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], []string]{A: sets.Set[string]{}, B: []string{"a", "b"}},
			Want: sets.Set[string]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()", func(s sets.Set[string], keys []string) sets.Set[string] {
		sets.RemoveAll(&s, keys...)
		return s
	}, testCases)
}
