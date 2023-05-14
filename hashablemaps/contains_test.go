package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestContains(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string],
		bool]{
		{
			Name: "Contains returns true for existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}}, B: "a"},
			Want: true,
		},
		{
			Name: "Contains returns false for non-existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}}, B: "b"},
			Want: false,
		},
		{
			Name: "Contains returns false for empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{}, B: "a"},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", hashablemaps.Contains[TestHashable], testCases)
}
