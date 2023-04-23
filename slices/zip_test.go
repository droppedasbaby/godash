package slices_test

import (
	"testing"

	"godash"
	"godash/slices"
	"godash/utils"
)

func TestZip(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[[]int, []string], []godash.Pair[int, string]]{
		{
			Name: "Zip with equal length slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []string]{
				A: []int{1, 2, 3},
				B: []string{"a", "b", "c"},
			},
			Want: []godash.Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
				{First: 3, Second: "c"},
			},
		},
		{
			Name: "Zip with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []string]{
				A: []int{},
				B: []string{},
			},
			Want: []godash.Pair[int, string]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Zip()", slices.Zip[int, string], testCases)
}
