package maps_test

import (
	"testing"

	"godash/maps"
	"godash/utils"
)

func TestContains_WithStringInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[string]int, string], bool]{
		{
			Name: "map contains key",
			Args: utils.TwoArgumentTestCasesArgsType[map[string]int, string]{
				A: map[string]int{"a": 1, "b": 2},
				B: "a",
			},
			Want: true,
		},
		{
			Name: "map does not contain key",
			Args: utils.TwoArgumentTestCasesArgsType[map[string]int, string]{
				A: map[string]int{"a": 1, "b": 2},
				B: "c",
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", maps.Contains[string, int], testCases)
}

func TestContains_WithFloat64Bool(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[float64]bool, float64], bool]{
		{
			Name: "map contains key",
			Args: utils.TwoArgumentTestCasesArgsType[map[float64]bool, float64]{
				A: map[float64]bool{1.0: true, 2.5: false},
				B: 2.5,
			},
			Want: true,
		},
		{
			Name: "map does not contain key",
			Args: utils.TwoArgumentTestCasesArgsType[map[float64]bool, float64]{
				A: map[float64]bool{1.0: true, 2.5: false},
				B: 3.0,
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", maps.Contains[float64, bool], testCases)
}
