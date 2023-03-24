package maps_test

import (
	"testing"

	"godash/maps"
	"godash/utils"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[map[string]int, string, int], map[string]int]{
		{
			Name: "Empty Map",
			Args: utils.ThreeArgumentTestCasesArgsType[map[string]int, string, int]{
				A: map[string]int{}, B: "a", C: 1,
			},
			Want: map[string]int{"a": 1},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.ThreeArgumentTestCasesArgsType[map[string]int, string, int]{
				A: map[string]int{"a": 1}, B: "b", C: 2,
			},
			Want: map[string]int{"a": 1, "b": 2},
		},
	}

	utils.RunThreeArgumentTestCases(t, "Add", func(m map[string]int, k string, v int) map[string]int {
		maps.Add(&m, k, v)
		return m
	}, testCases)
}

func TestAdd_WithString(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[map[int]string, int, string], map[int]string]{
		{
			Name: "Empty Map",
			Args: utils.ThreeArgumentTestCasesArgsType[map[int]string, int, string]{
				A: map[int]string{}, B: 1, C: "a",
			},
			Want: map[int]string{1: "a"},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.ThreeArgumentTestCasesArgsType[map[int]string, int, string]{
				A: map[int]string{1: "a"}, B: 2, C: "b",
			},
			Want: map[int]string{1: "a", 2: "b"},
		},
	}

	utils.RunThreeArgumentTestCases(t, "Add()", func(m map[int]string, k int, v string) map[int]string {
		maps.Add(&m, k, v)
		return m
	}, testCases)
}
