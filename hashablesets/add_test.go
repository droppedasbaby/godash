// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable],
		hashablesets.HashableSet[TestHashable],
	]{
		{
			Name: "Add a new element to an empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: TestHashable{"a"},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
		},
		{
			Name: "Add a new element to a non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: TestHashable{"b"},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
		{
			Name: "Add a duplicate element to the set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: TestHashable{"a"},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Add()",
		func(s hashablesets.HashableSet[TestHashable], e TestHashable) hashablesets.HashableSet[TestHashable] {
			hashablesets.Add(&s, e)
			return s
		}, testCases)
}
