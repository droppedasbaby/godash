package sets_test

import (
	"sort"
	"testing"

	"godash/sets"
	"godash/utils"
)

func TestToSet(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]string], sets.Set[string]]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]string]{A: []string{}},
			Want: sets.Set[string]{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]string]{A: []string{"b", "a"}},
			Want: sets.Set[string]{"a": true, "b": true},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSet()", sets.ToSet[string], testCases)
}

func TestToSet_WithInts(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]int], sets.Set[int]]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{}},
			Want: sets.Set[int]{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{2, 1}},
			Want: sets.Set[int]{1: true, 2: true},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSetWithInts()", sets.ToSet[int], testCases)
}

func TestFromSet(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[sets.Set[string]], []string]{
		{
			Name: "Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[sets.Set[string]]{A: sets.Set[string]{}},
			Want: []string{},
		},
		{
			Name: "Non-Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[sets.Set[string]]{A: sets.Set[string]{"a": true, "b": true}},
			Want: []string{"a", "b"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSet()", func(s sets.Set[string]) []string {
		res := sets.FromSet(s)
		sort.Strings(res)
		return res
	}, testCases)
}

func TestFromSet_WithInts(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[sets.Set[int]], []int]{
		{
			Name: "Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[sets.Set[int]]{A: sets.Set[int]{}},
			Want: []int{},
		},
		{
			Name: "Non-Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[sets.Set[int]]{A: sets.Set[int]{1: true, 2: true}},
			Want: []int{1, 2},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSet()", func(s sets.Set[int]) []int {
		res := sets.FromSet(s)
		sort.Ints(res)
		return res
	}, testCases)
}
