package hashablesets_test

import (
	"testing"

	"godash/hashablesets"
	"godash/utils"
)

func TestIsSubsetOf(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]],
		bool]{
		{
			Name: "Empty set is a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{}, B: hashablesets.HashableSet[TestHashable]{}},
			Want: true,
		},
		{
			Name: "Empty set is a subset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
			},
			Want: true,
		},
		{
			Name: "Non-empty set is not a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{}},
			Want: false,
		},
		{
			Name: "Set is a subset of itself",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
			},
			Want: true,
		},
		{
			Name: "Non-empty set is a subset of a larger set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
			},
			Want: true,
		},
		{
			Name: "Non-empty set is not a subset of a smaller set",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}}},
			Want: false,
		},
		{
			Name: "Is subset",
			Args: utils.TwoArgumentTestCasesArgsType[
				hashablesets.HashableSet[TestHashable], hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}}},
			Want: true,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSubsetOf()", hashablesets.IsSubsetOf[TestHashable], testCases)
}
