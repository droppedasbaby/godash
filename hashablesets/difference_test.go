package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestDifference(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]],
		hashablesets.HashableSet[TestHashable]]{
		{
			Name: "Difference between two non-empty sets with distinct elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{"c": TestHashable{"c"}, "d": TestHashable{"d"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
		{
			Name: "Difference between two non-empty sets with common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{"b": TestHashable{"b"}, "c": TestHashable{"c"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
		},
		{
			Name: "Difference between two non-empty sets with identical elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Difference between an empty set and a non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Difference between a non-empty set and an empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Difference()", hashablesets.Difference[TestHashable], testCases)
}
