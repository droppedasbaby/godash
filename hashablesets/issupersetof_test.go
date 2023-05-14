// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestIsSupersetOf(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			hashablesets.HashableSet[TestHashable],
			hashablesets.HashableSet[TestHashable]], bool]{
		{
			Name: "Empty sets are supersets of each other",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{},
			},
			Want: true,
		},
		{
			Name: "Empty set is a superset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{},
			},
			Want: true,
		},
		{
			Name: "Non-empty set is not a superset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
			},
			Want: false,
		},
		{
			Name: "Set is a superset of itself",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}}},
			Want: true,
		},
		{
			Name: "Non-overlapping sets are not supersets of each other",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{"b": TestHashable{"b"}}},
			Want: false,
		},
		{
			Name: "Is superset",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"},
				},
			},
			Want: true,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSupersetOf()", hashablesets.IsSupersetOf[TestHashable], testCases)
}
