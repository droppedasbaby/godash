// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices_test

import (
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
)

func TestCompact(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[]*string], []*string]{
		{
			Name: "Compact with empty slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]*string]{
				A: []*string{},
			},
			Want: []*string{},
		},
		{
			Name: "Compact with all non-nil elements",
			Args: utils.SingleArgumentTestCasesArgsType[[]*string]{
				A: []*string{ptr("a"), ptr("b"), ptr("c")},
			},
			Want: []*string{ptr("a"), ptr("b"), ptr("c")},
		},
		{
			Name: "Compact with all nil elements",
			Args: utils.SingleArgumentTestCasesArgsType[[]*string]{
				A: []*string{nil, nil, nil},
			},
			Want: []*string{},
		},
		{
			Name: "Compact with mixed nil and non-nil elements",
			Args: utils.SingleArgumentTestCasesArgsType[[]*string]{
				A: []*string{nil, ptr("a"), nil, ptr("b"), nil, ptr("c"), nil},
			},
			Want: []*string{ptr("a"), ptr("b"), ptr("c")},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Compact()", slices.Compact[string], testCases)
}
