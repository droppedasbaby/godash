// Package hashablesets provides functionality for creating and manipulating
// set data structures in Go with hashable elements.
package hashablesets_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablesets"
	"github.com/GrewalAS/godash/utils"
)

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable],
		hashablesets.HashableSet[TestHashable]]{
		{
			Name: "Remove multiple elements from empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{}, B: []TestHashable{{"a"}, {"b"}}},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Remove non-existent elements from non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: []TestHashable{{"c"}, {"d"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
		{
			Name: "Remove existing elements from set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"},
					"b": TestHashable{"b"},
					"c": TestHashable{"c"},
				},
				B: []TestHashable{{"a"}, {"b"}}},
			Want: hashablesets.HashableSet[TestHashable]{"c": TestHashable{"c"}},
		},
		{
			Name: "Remove all elements from set",
			Args: utils.TwoArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable], []TestHashable]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: []TestHashable{{"a"}, {"b"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()",
		func(s hashablesets.HashableSet[TestHashable], es []TestHashable) hashablesets.HashableSet[TestHashable] {
			hashablesets.RemoveAll(&s, es...)
			return s
		}, testCases)
}
