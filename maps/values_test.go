package maps_test

import (
	"sort"
	"testing"

	"github.com/GrewalAS/godash/maps"
	"github.com/GrewalAS/godash/utils"
)

func TestValues(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[map[string]int], []int]{
		{
			Name: "Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]int]{A: map[string]int{}},
			Want: []int{},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]int]{A: map[string]int{"a": 1, "b": 2}},
			Want: []int{1, 2},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Values()", func(m map[string]int) []int {
		vals := maps.Values[string, int](m)
		sort.Ints(vals)
		return vals
	}, testCases)
}

func TestValues_WithIntKeys(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[map[int]string], []string]{
		{
			Name: "Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{A: map[int]string{}},
			Want: []string{},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{A: map[int]string{1: "a", 2: "b"}},
			Want: []string{"a", "b"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Values()", func(m map[int]string) []string {
		vals := maps.Values[int, string](m)
		sort.Strings(vals)
		return vals
	}, testCases)
}
