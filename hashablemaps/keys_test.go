package hashablemaps_test

import (
	"testing"

	"godash/hashablemaps"
	"godash/utils"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]],
		[]string,
	]{
		{
			Name: "Empty HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{},
			},
			Want: []string{},
		},
		{
			Name: "Single element HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
			},
			Want: []string{"a"},
		},
		{
			Name: "Multiple elements HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{
					"a": {Value: "a"},
					"b": {Value: "b"},
					"c": {Value: "c"},
				},
			},
			Want: []string{"a", "b", "c"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Keys()", hashablemaps.Keys[TestHashable], testCases)
}
