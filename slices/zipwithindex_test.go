// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices_test

import (
	"testing"

	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
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
