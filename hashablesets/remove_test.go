package hashablesets_test

import (
	"testing"

	"godash/hashablesets"
	"godash/utils"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable],
		hashablesets.HashableSet[TestHashable]]{
		{
			Name: "Remove element from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{}, B: TestHashable{"a"}},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Remove non-existent element from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: TestHashable{"c"},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
		{
			Name: "Remove existing element from set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: TestHashable{"a"},
			},
			Want: hashablesets.HashableSet[TestHashable]{"b": TestHashable{"b"}},
		},
		{
			Name: "Remove last element from set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}}, B: TestHashable{"a"}},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(
		s hashablesets.HashableSet[TestHashable], h TestHashable) hashablesets.HashableSet[TestHashable] {
		hashablesets.Remove[TestHashable](&s, h)
		return s
	}, testCases)
}
