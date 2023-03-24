package maps_test

import (
	"testing"

	"godash/maps"
	"godash/utils"
)

func TestRemove(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[string]int, string], map[string]int]{
		{
			Name: "Remove Existing Key",
			Args: utils.TwoArgumentTestCasesArgsType[map[string]int, string]{A: map[string]int{"a": 1, "b": 2}, B: "a"},
			Want: map[string]int{"b": 2},
		},
		{
			Name: "Remove Non-Existing Key",
			Args: utils.TwoArgumentTestCasesArgsType[map[string]int, string]{A: map[string]int{"a": 1, "b": 2}, B: "c"},
			Want: map[string]int{"a": 1, "b": 2},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Remove()", func(m map[string]int, k string) map[string]int {
		maps.Remove[string, int](&m, k)
		return m
	}, testCases)
}

func TestRemove_WithIntKeys(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[int]string, int], map[int]string]{
		{
			Name: "Remove Existing Key",
			Args: utils.TwoArgumentTestCasesArgsType[map[int]string, int]{A: map[int]string{1: "a", 2: "b"}, B: 1},
			Want: map[int]string{2: "b"},
		},
		{
			Name: "Remove Non-Existing Key",
			Args: utils.TwoArgumentTestCasesArgsType[map[int]string, int]{A: map[int]string{1: "a", 2: "b"}, B: 3},
			Want: map[int]string{1: "a", 2: "b"},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveWithIntKeys()", func(m map[int]string, k int) map[int]string {
		maps.Remove[int, string](&m, k)
		return m
	}, testCases)
}
