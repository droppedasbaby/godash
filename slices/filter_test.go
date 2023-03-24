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
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(i int) bool { return i%2 == 0 }},
			Want: []int{2, 4, 6},
		},
		{
			Name: "all odd numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(i int) bool { return i%2 != 0 }},
			Want: []int{1, 3, 5},
		},
		{
			Name: "all positive numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{-3, -2, -1, 0, 1, 2, 3}, B: func(i int) bool { return i > 0 }},
			Want: []int{1, 2, 3},
		},
		{
			Name: "all non-positive numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{-3, -2, -1, 0, 1, 2, 3}, B: func(i int) bool { return i <= 0 }},
			Want: []int{-3, -2, -1, 0},
		},
		{
			Name: "all numbers greater than 3",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(i int) bool { return i > 3 }},
			Want: []int{4, 5, 6},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{}, B: func(i int) bool { return i%2 == 0 }},
			Want: []int{},
		},
		{
			Name: "no numbers satisfying the predicate",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) bool]{A: []int{1, 2, 3, 4, 5, 6}, B: func(i int) bool { return i > 6 }},
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
				B: func(f float64) bool { return f > 0 },
			},
			Want: []float64{1.2, 2.3, 3.7},
		},
		{
			Name: "all numbers less than 1",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) bool]{
				A: []float64{-0.5, 0.0, 0.5, 1.0, 1.5, 2.0, 3.5},
				B: func(f float64) bool { return f < 1 },
			},
			Want: []float64{-0.5, 0.0, 0.5},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) bool]{A: []float64{}, B: func(f float64) bool { return f > 0 }},
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
				B: func(s string) bool { return len(s) > 3 },
			},
			Want: []string{"apple", "banana", "cherry", "date", "grape"},
		},
		{
			Name: "all strings starting with 'a' or 'A'",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]string, func(string) bool]{
				A: []string{"Apple", "banana", "avocado", "Date", "Apricot", "grape"},
				B: func(s string) bool { return strings.HasPrefix(strings.ToLower(s), "a") },
			},
			Want: []string{"Apple", "avocado", "Apricot"},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]string, func(string) bool]{A: []string{}, B: func(s string) bool { return len(s) > 0 }},
			Want: []string{},
		},
	}
	utils.RunTwoArgumentTestCases(t, "Filter()", slices.Filter[string], testCases)
}

func TestFilter_WithArbStruct(t *testing.T) {
	t.Parallel()
	type testStruct struct {
		X int
		Y int
	}

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) bool], []testStruct]{
		{
			Name: "filter elements where X is even",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) bool]{
				A: []testStruct{
					{X: 1, Y: 10},
					{X: 2, Y: 20},
					{X: 3, Y: 30},
					{X: 4, Y: 40},
				},
				B: func(v testStruct) bool {
					return v.X%2 == 0
				},
			},
			Want: []testStruct{
				{X: 2, Y: 20},
				{X: 4, Y: 40},
			},
		},
		{
			Name: "filter elements where Y is greater than X",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) bool]{
				A: []testStruct{
					{X: 5, Y: 15},
					{X: 6, Y: 4},
					{X: 7, Y: 10},
					{X: 8, Y: 8},
				},
				B: func(v testStruct) bool {
					return v.Y > v.X
				},
			},
			Want: []testStruct{
				{X: 5, Y: 15},
				{X: 7, Y: 10},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Filter()", slices.Filter[testStruct], testCases)
}
