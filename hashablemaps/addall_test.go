// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestAddAll(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []TestHashable],
		hashablemaps.HashableMap[TestHashable]]{
		{
			Name: "Add multiple elements to empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{},
				B: []TestHashable{{"a"}, {"b"}, {"c"}},
			},
			Want: hashablemaps.HashableMap[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
		{
			Name: "Add multiple elements to non-empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}},
				B: []TestHashable{{"b"}, {"c"}},
			},
			Want: hashablemaps.HashableMap[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
		{
			Name: "Add duplicate and new elements to HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
				B: []TestHashable{{"b"}, {"c"}},
			},
			Want: hashablemaps.HashableMap[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()",
		func(m hashablemaps.HashableMap[TestHashable], v []TestHashable) hashablemaps.HashableMap[TestHashable] {
			hashablemaps.AddAll(&m, v...)
			return m
		}, testCases)
}
