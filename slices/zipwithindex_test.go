package slices_test

import (
	"godash"
	"godash/slices"
	"godash/utils"
	"testing"
)

func TestZipWithIndex(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[]int], []godash.Pair[int, int]]{
		{
			Name: "ZipWithIndex with non-empty slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{
				A: []int{1, 2, 3},
			},
			Want: []godash.Pair[int, int]{
				{First: 0, Second: 1},
				{First: 1, Second: 2},
				{First: 2, Second: 3},
			},
		},
		{
			Name: "ZipWithIndex with empty slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]int]{
				A: []int{},
			},
			Want: []godash.Pair[int, int]{},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ZipWithIndex()", slices.ZipWithIndex[int], testCases)
}
