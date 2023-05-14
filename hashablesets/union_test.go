// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			hashablesets.HashableSet[TestHashable],
			hashablesets.HashableSet[TestHashable]],
		hashablesets.HashableSet[TestHashable]]{
		{
			Name: "Union of two empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable],
				hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{},
			},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Union of an empty and non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable],
				hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
		},
		{
			Name: "Union of non-empty sets with no common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable],
				hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{"b": TestHashable{"b"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
			},
		},
		{
			Name: "Union of non-empty sets with common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable],
				hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"},
					"b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{
					"b": TestHashable{"b"},
					"c": TestHashable{"c"}}},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Union()", hashablesets.Union[TestHashable], testCases)
}
