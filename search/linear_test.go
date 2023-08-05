package search_test

import (
	"testing"

	"github.com/GrewalAS/godash/search"
	"github.com/GrewalAS/godash/utils"
)

func TestLinearSearch_Int(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]int, int], int]{
		{
			Name: "Target In Int Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{
				A: []int{1, 2, 3, 4, 5}, B: 3,
			},
			Want: 2,
		},
		{
			Name: "Target Not In Int Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{
				A: []int{1, 2, 3, 4, 5}, B: 6,
			},
			Want: -1,
		},
	}
	utils.RunTwoArgumentTestCases(t, "LinearSearch", search.LinearSearch[int], testCases)
}

func TestLinearSearch_Float(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]float64, float64], int]{
		{
			Name: "Target In Float Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]float64, float64]{
				A: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, B: 3.3,
			},
			Want: 2,
		},
		{
			Name: "Target Not In Float Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]float64, float64]{
				A: []float64{1.1, 2.2, 3.3, 4.4, 5.5}, B: 6.6,
			},
			Want: -1,
		},
	}
	utils.RunTwoArgumentTestCases(t, "LinearSearch", search.LinearSearch[float64], testCases)
}

func TestLinearSearch_String(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]string, string], int]{
		{
			Name: "Target In String Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, string]{
				A: []string{"one", "two", "three", "four", "five"}, B: "three",
			},
			Want: 2,
		},
		{
			Name: "Target Not In String Slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, string]{
				A: []string{"one", "two", "three", "four", "five"}, B: "six",
			},
			Want: -1,
		},
	}
	utils.RunTwoArgumentTestCases(t, "LinearSearch", search.LinearSearch[string], testCases)
}

func TestLinearSearchWith_Int(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool], int]{
		{
			Name: "Empty Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{}, B: 1, C: func(a, b int) bool { return a == b },
			},
			Want: -1,
		},
		{
			Name: "Single Element Slice, Target Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1}, B: 1, C: func(a, b int) bool { return a == b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Not Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1}, B: 2, C: func(a, b int) bool { return a == b },
			},
			Want: -1,
		},
		{
			Name: "Target Is First Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 1, C: func(a, b int) bool { return a == b },
			},
			Want: 0,
		},
		{
			Name: "Target Is Last Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 5, C: func(a, b int) bool { return a == b },
			},
			Want: 4,
		},
		{
			Name: "Target In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 3, C: func(a, b int) bool { return a == b },
			},
			Want: 2,
		},
		{
			Name: "Target Not In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 6, C: func(a, b int) bool { return a == b },
			},
			Want: -1,
		},
	}
	utils.RunThreeArgumentTestCases(t, "LinearSearchWith", search.LinearSearchWith[int], testCases)
}

func TestLinearSearchWith_Float(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool], int]{
		{
			Name: "Target In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				B: 3.3,
				C: func(a, b float64) bool { return a == b },
			},
			Want: 2,
		},
		{
			Name: "Target Not In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				B: 6.6,
				C: func(a, b float64) bool { return a == b },
			},
			Want: -1,
		},
	}
	utils.RunThreeArgumentTestCases(t, "LinearSearchWith", search.LinearSearchWith[float64], testCases)
}

func TestLinearSearchWith_StringLen(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool], int]{
		{
			Name: "Target In String Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool]{
				A: []string{"a", "aa", "aaa", "aaaa"},
				B: "aaa",
				C: func(a, b string) bool { return len(a) == len(b) },
			},
			Want: 2,
		},
		{
			Name: "Target Not In String Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool]{
				A: []string{"a", "aa", "aaa"},
				B: "aaaa",
				C: func(a, b string) bool { return len(a) == len(b) },
			},
			Want: -1,
		},
	}
	utils.RunThreeArgumentTestCases(t, "LinearSearchWith", search.LinearSearchWith[string], testCases)
}
