package hashablemaps_test

import (
	"sort"
	"testing"

	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/hashablemaps"
	"github.com/GrewalAS/godash/utils"
)

func TestFromHashableMap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]],
		[]godash.Pair[string, TestHashable],
	]{
		{
			Name: "Empty HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{},
			},
			Want: []godash.Pair[string, TestHashable]{},
		},
		{
			Name: "Single element HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}},
			},
			Want: []godash.Pair[string, TestHashable]{
				{First: "a", Second: TestHashable{"a"}},
			},
		},
		{
			Name: "Multiple elements HashableMap",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{
					"a": TestHashable{"a"},
					"b": TestHashable{"b"},
					"c": TestHashable{"c"},
				},
			},
			Want: []godash.Pair[string, TestHashable]{
				{First: "a", Second: TestHashable{"a"}},
				{First: "b", Second: TestHashable{"b"}},
				{First: "c", Second: TestHashable{"c"}},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromHashableMap()",
		func(m hashablemaps.HashableMap[TestHashable]) []godash.Pair[string, TestHashable] {
			ps := hashablemaps.FromHashableMap[TestHashable](m)
			sort.Slice(ps, func(i, j int) bool { return ps[i].Second.Value < ps[j].Second.Value })
			return ps
		}, testCases)
}

func TestToHashableMap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[]TestHashable], hashablemaps.HashableMap[TestHashable]]{
		{
			Name: "Empty Hashable slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{},
			},
			Want: hashablemaps.HashableMap[TestHashable]{},
		},
		{
			Name: "Single element Hashable slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{{Value: "a"}},
			},
			Want: hashablemaps.HashableMap[TestHashable]{"a": {Value: "a"}},
		},
		{
			Name: "Multiple elements Hashable slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{
					{Value: "a"},
					{Value: "b"},
					{Value: "c"},
				},
			},
			Want: hashablemaps.HashableMap[TestHashable]{
				"a": {Value: "a"},
				"b": {Value: "b"},
				"c": {Value: "c"},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToHashableMap()", hashablemaps.ToHashableMap[TestHashable], testCases)
}
