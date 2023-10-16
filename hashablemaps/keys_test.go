package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestKeys(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[
		hashablemaps.HashableMap[TestHashable], []string], bool]{
		{
			Name: "Empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{},
				B: []string{},
			},
			Want: true,
		},
		{
			Name: "Single element HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
				B: []string{"a"},
			},
			Want: true,
		},
		{
			Name: "Multiple elements HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{
					"a": {Value: "a"},
					"b": {Value: "b"},
					"c": {Value: "c"},
				},
				B: []string{"a", "b", "c"},
			},
			Want: true,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Keys()", func(h hashablemaps.HashableMap[TestHashable], ks []string) bool {
		return utils.TestingSlicesEqualWithoutOrder(hashablemaps.Keys[TestHashable](h), ks)
	}, testCases)
}
