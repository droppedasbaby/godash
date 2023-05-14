// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestRemove(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string],
		hashablemaps.HashableMap[TestHashable],
	]{
		{
			Name: "Remove from empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{},
				B: "a",
			},
			Want: hashablemaps.HashableMap[TestHashable]{},
		},
		{
			Name: "Remove existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
				B: "a",
			},
			Want: hashablemaps.HashableMap[TestHashable]{},
		},
		{
			Name: "Remove existing key with multiple keys",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}, "b": {Value: "b"}},
				B: "a",
			},
			Want: hashablemaps.HashableMap[TestHashable]{"b": {Value: "b"}},
		},
		{
			Name: "Remove non-existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
				B: "b",
			},
			Want: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()",
		func(m hashablemaps.HashableMap[TestHashable], key string) hashablemaps.HashableMap[TestHashable] {
			hashablemaps.Remove(&m, key)
			return m
		}, testCases)
}
