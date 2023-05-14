package maps_test

import (
	"testing"

	"github.com/GrewalAS/godash/utils"
)

func TestRemoveAll(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[string]int, []string], map[string]int]{
		{
			Name: "Remove Multiple Existing Keys",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[string]int, []string]{
				A: map[string]int{"a": 1, "b": 2, "c": 3}, B: []string{"a", "b"},
			},
			Want: map[string]int{"c": 3},
		},
		{
			Name: "Remove Non-Existing Keys",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[string]int, []string]{
				A: map[string]int{"a": 1, "b": 2, "c": 3}, B: []string{"d", "e"},
			},
			Want: map[string]int{"a": 1, "b": 2, "c": 3},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()", func(m map[string]int, ks []string) map[string]int {
		maps.RemoveAll[string, int](&m, ks...)
		return m
	}, testCases)
}

func TestRemoveAll_WithIntKeys(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[map[int]string, []int], map[int]string]{
		{
			Name: "Remove Multiple Existing Keys",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[int]string, []int]{
				A: map[int]string{1: "a", 2: "b", 3: "c"}, B: []int{1, 2},
			},
			Want: map[int]string{3: "c"},
		},
		{
			Name: "Remove Non-Existing Keys",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[int]string, []int]{
				A: map[int]string{1: "a", 2: "b", 3: "c"}, B: []int{4, 5},
			},
			Want: map[int]string{1: "a", 2: "b", 3: "c"},
		},
	}

	utils.RunTwoArgumentTestCases(t, "RemoveAll()", func(m map[int]string, ks []int) map[int]string {
		maps.RemoveAll[int, string](&m, ks...)
		return m
	}, testCases)
}
