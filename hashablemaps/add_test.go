// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], TestHashable],
		hashablemaps.HashableMap[TestHashable]]{
		{
			Name: "Add single element to empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{}, B: TestHashable{"a"}},
			Want: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}},
		},
		{
			Name: "Add single element to non-empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}}, B: TestHashable{"b"}},
			Want: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
		},
		{
			Name: "Add duplicate element to HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], TestHashable]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}}, B: TestHashable{"a"}},
			Want: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Add()",
		func(m hashablemaps.HashableMap[TestHashable], v TestHashable) hashablemaps.HashableMap[TestHashable] {
			hashablemaps.Add(&m, v)
			return m
		}, testCases)
}
