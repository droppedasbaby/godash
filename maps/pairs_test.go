package maps_test

import (
	"sort"
	"testing"

	"godash"
	"godash/maps"
	"godash/utils"
)

func TestFromMap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[map[string]int], []godash.Pair[string, int]]{
		{
			Name: "Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]int]{A: map[string]int{}},
			Want: []godash.Pair[string, int]{},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]int]{A: map[string]int{"a": 1}},
			Want: []godash.Pair[string, int]{
				{First: "a", Second: 1},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromMap()", func(m map[string]int) []godash.Pair[string, int] {
		pairs := maps.FromMap(m)
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].First < pairs[j].First })
		return pairs
	}, testCases)
}

func TestFromMap_WithIntKeys(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[map[int]string], []godash.Pair[int, string]]{
		{
			Name: "Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{A: map[int]string{}},
			Want: []godash.Pair[int, string]{},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{A: map[int]string{1: "a", 2: "b"}},
			Want: []godash.Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromMap()", func(m map[int]string) []godash.Pair[int, string] {
		pairs := maps.FromMap(m)
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].First < pairs[j].First })
		return pairs
	}, testCases)
}

func TestToMap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[]godash.Pair[string, int]], map[string]int]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[string, int]]{A: []godash.Pair[string, int]{}},
			Want: map[string]int{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[string, int]]{
				A: []godash.Pair[string, int]{
					{First: "a", Second: 1},
					{First: "b", Second: 2},
				}},
			Want: map[string]int{"a": 1, "b": 2},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToMap()", func(pairs []godash.Pair[string, int]) map[string]int {
		return maps.ToMap(pairs...)
	}, testCases)
}

func TestToMap_WithIntKeys(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]], map[int]string]{
		{
			Name: "Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]]{A: []godash.Pair[int, string]{}},
			Want: map[int]string{},
		},
		{
			Name: "Non-Empty Slice",
			Args: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]]{A: []godash.Pair[int, string]{
				{First: 1, Second: "a"},
				{First: 2, Second: "b"},
			}},
			Want: map[int]string{1: "a", 2: "b"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToMap()", func(pairs []godash.Pair[int, string]) map[int]string {
		return maps.ToMap(pairs...)
	}, testCases)
}
