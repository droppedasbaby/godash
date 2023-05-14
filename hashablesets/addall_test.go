package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestAddAll(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable],
		hashablesets.HashableSet[TestHashable],
	]{
		{
			Name: "Add multiple elements to an empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{},
				B: []TestHashable{{"a"}, {"b"}, {"c"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
		{
			Name: "Add multiple elements to a non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
				B: []TestHashable{{"b"}, {"c"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
		{
			Name: "Add duplicate and new elements",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: []TestHashable{{"b"}, {"c"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()",
		func(s hashablesets.HashableSet[TestHashable], es []TestHashable) hashablesets.HashableSet[TestHashable] {
			hashablesets.AddAll(&s, es...)
			return s
		}, testCases)
}
