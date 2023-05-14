// Package hashablemaps offers tools for handling and manipulating
// map data structures in Go with hashable keys.
package hashablemaps_test

import (
	"testing"

	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestRemoveAll(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string],
		hashablemaps.HashableMap[TestHashable],
	]{
		{
			Name: "Remove from empty HashableMap",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{},
				B: []string{"a"},
			},
			Want: hashablemaps.HashableMap[TestHashable]{},
		},
		{
			Name: "Remove existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
				B: []string{"a"},
			},
			Want: hashablemaps.HashableMap[TestHashable]{},
		},
		{
			Name: "Remove existing keys with multiple keys",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}, "b": {Value: "b"}, "c": {Value: "c"}},
				B: []string{"a", "c"},
			},
			Want: hashablemaps.HashableMap[TestHashable]{"b": {Value: "b"}},
		},
		{
			Name: "Remove non-existing key",
			Args: utils.TwoArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable], []string]{
				A: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
				B: []string{"b"},
			},
			Want: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()",
		func(m hashablemaps.HashableMap[TestHashable], keys []string) hashablemaps.HashableMap[TestHashable] {
			hashablemaps.RemoveAll(&m, keys...)
			return m
		}, testCases)
}
