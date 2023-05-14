// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestContains(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable],
		bool,
	]{
		{
			Name: "Contains returns true for existing element",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: TestHashable{"a"},
			},
			Want: true,
		},
		{
			Name: "Contains returns false for non-existing element",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: TestHashable{"b"},
			},
			Want: false,
		},
		{
			Name: "Contains returns false for empty HashableSet",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: TestHashable{"a"},
			},
			Want: false,
		},
		{
			Name: "Contains returns true for existing element in a set with multiple elements",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"},
					"b": TestHashable{"b"},
					"c": TestHashable{"c"},
				},
				B: TestHashable{"b"},
			},
			Want: true,
		},
		{
			Name: "Contains returns false for non-existing element in a set with multiple elements",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"},
					"b": TestHashable{"b"},
					"c": TestHashable{"c"},
				},
				B: TestHashable{"d"},
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", hashablesets.Contains[TestHashable], testCases)
}
