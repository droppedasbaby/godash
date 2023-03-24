package slices_test

import (
	"strings"
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestFilter_WithInt(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[
		[]int, func(int) bool], []int]{
		{
			Name: "all even numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(v int) bool { return v%2 == 0 }},
			Want: []int{2, 4, 6},
		},
		{
			Name: "all odd numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(v int) bool { return v%2 != 0 }},
			Want: []int{1, 3, 5},
		},
		{
			Name: "all positive numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{-3, -2, -1, 0, 1, 2, 3}, B: func(v int) bool { return v > 0 }},
			Want: []int{1, 2, 3},
		},
		{
			Name: "all non-positive numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{-3, -2, -1, 0, 1, 2, 3}, B: func(v int) bool { return v <= 0 }},
			Want: []int{-3, -2, -1, 0},
		},
		{
			Name: "all numbers greater than 3",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(v int) bool { return v > 3 }},
			Want: []int{4, 5, 6},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{}, B: func(v int) bool { return v%2 == 0 }},
			Want: []int{},
		},
		{
			Name: "no numbers satisfying the predicate",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(v int) bool { return v > 6 }},
			Want: []int{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Filter()", slices.Filter[int], testCases)
}

func TestFilter_WithFloat64(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[
		[]float64, func(float64) bool], []float64]{
		{
			Name: "all positive numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) bool]{
				A: []float64{-3.5, -2.1, -1.0, 0.0, 1.2, 2.3, 3.7},
				B: func(v float64) bool { return v > 0 },
			},
			Want: []float64{1.2, 2.3, 3.7},
		},
		{
			Name: "all numbers less than 1",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) bool]{
				A: []float64{-0.5, 0.0, 0.5, 1.0, 1.5, 2.0, 3.5},
				B: func(v float64) bool { return v < 1 },
			},
			Want: []float64{-0.5, 0.0, 0.5},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) bool]{A: []float64{}, B: func(v float64) bool { return v > 0 }},
			Want: []float64{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Filter()", slices.Filter[float64], testCases)
}

func TestFilter_WithString(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[
		[]string, func(string) bool], []string]{
		{
			Name: "all strings with length greater than 3",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]string, func(string) bool]{
				A: []string{"apple", "banana", "cherry", "date", "fig", "grape"},
				B: func(v string) bool { return len(v) > 3 },
			},
			Want: []string{"apple", "banana", "cherry", "date", "grape"},
		},
		{
			Name: "all strings starting with 'a' or 'A'",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]string, func(string) bool]{
				A: []string{"Apple", "banana", "avocado", "Date", "Apricot", "grape"},
				B: func(v string) bool { return strings.HasPrefix(strings.ToLower(v), "a") },
			},
			Want: []string{"Apple", "avocado", "Apricot"},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]string, func(string) bool]{A: []string{}, B: func(v string) bool { return len(v) > 0 }},
			Want: []string{},
		},
	}
	utils.RunTwoArgumentTestCases(t, "Filter()", slices.Filter[string], testCases)
}
