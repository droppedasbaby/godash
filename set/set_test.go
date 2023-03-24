package set_test

import (
	"sort"
	"testing"

	"godash/set"
	"godash/utils"
)

func TestToSet(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]string], set.Set[string]]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]string]{A: []string{}},
			Want: set.Set[string]{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]string]{A: []string{"b", "a"}},
			Want: set.Set[string]{"a": true, "b": true},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSet()", set.ToSet[string], testCases)
}

func TestToSet_WithInts(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]int], set.Set[int]]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{}},
			Want: set.Set[int]{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{2, 1}},
			Want: set.Set[int]{1: true, 2: true},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSetWithInts()", set.ToSet[int], testCases)
}

func TestFromSet(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[set.Set[string]], []string]{
		{
			Name: "Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[set.Set[string]]{A: set.Set[string]{}},
			Want: []string{},
		},
		{
			Name: "Non-Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[set.Set[string]]{A: set.Set[string]{"a": true, "b": true}},
			Want: []string{"a", "b"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSet()", func(s set.Set[string]) []string {
		res := set.FromSet(s)
		sort.Strings(res)
		return res
	}, testCases)
}

func TestFromSet_WithInts(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[set.Set[int]], []int]{
		{
			Name: "Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[set.Set[int]]{A: set.Set[int]{}},
			Want: []int{},
		},
		{
			Name: "Non-Empty Set",
			Args: utils.SingleArgumentTestCasesArgsType[set.Set[int]]{A: set.Set[int]{1: true, 2: true}},
			Want: []int{1, 2},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSet()", func(s set.Set[int]) []int {
		res := set.FromSet(s)
		sort.Ints(res)
		return res
	}, testCases)
}
