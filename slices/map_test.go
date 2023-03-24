package slices_test

import (
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestMap_WithInt(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[[]int, func(int) int], []int]{
		{
			Name: "double all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) int]{
				A: []int{1, 2, 3, 4, 5}, B: func(i int) int { return i * 2 },
			},
			Want: []int{2, 4, 6, 8, 10},
		},
		{
			Name: "subtract 3 from all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) int]{
				A: []int{4, 5, 6, 7, 8}, B: func(i int) int { return i - 3 },
			},
			Want: []int{1, 2, 3, 4, 5},
		},
		{
			Name: "square all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) int]{
				A: []int{1, 2, 3, 4}, B: func(i int) int { return i * i },
			},
			Want: []int{1, 4, 9, 16},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]int, func(int) int]{
				A: []int{}, B: func(i int) int { return i * 2 },
			},
			Want: []int{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Map()", slices.Map[int, int], testCases)
}

func TestMap_WithFloat64(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[
		[]float64, func(float64) float64], []float64]{
		{
			Name: "double all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) float64]{
				A: []float64{1.0, 2.5, 3.3, 4.7, 5.0}, B: func(f float64) float64 { return f * 2 },
			},
			Want: []float64{2.0, 5.0, 6.6, 9.4, 10.0},
		},
		{
			Name: "add 1.5 to all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) float64]{
				A: []float64{1.0, 2.5, 3.3, 4.7, 5.0}, B: func(f float64) float64 { return f + 1.5 },
			},
			Want: []float64{2.5, 4.0, 4.8, 6.2, 6.5},
		},
		{
			Name: "square all numbers",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) float64]{
				A: []float64{1.0, 2.0, 3.0, 4.0}, B: func(f float64) float64 { return f * f },
			},
			Want: []float64{1.0, 4.0, 9.0, 16.0},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[
				[]float64, func(float64) float64]{
				A: []float64{}, B: func(f float64) float64 { return f * 2 },
			},
			Want: []float64{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Map()", slices.Map[float64, float64], testCases)
}

func TestMap_WithArbStruct(t *testing.T) {
	t.Parallel()
	type testStruct struct {
		X int
		Y int
	}

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct], []testStruct]{
		{
			Name: "increment X and double Y",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{
					{X: 1, Y: 10},
					{X: 2, Y: 20},
					{X: 3, Y: 30},
				},
				B: func(v testStruct) testStruct {
					return testStruct{X: v.X + 1, Y: v.Y * 2}
				},
			},
			Want: []testStruct{
				{X: 2, Y: 20},
				{X: 3, Y: 40},
				{X: 4, Y: 60},
			},
		},
		{
			Name: "double X and subtract Y by 5",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{
					{X: 5, Y: 15},
					{X: 6, Y: 30},
					{X: 7, Y: 45},
				},
				B: func(v testStruct) testStruct {
					return testStruct{X: v.X * 2, Y: v.Y - 5}
				},
			},
			Want: []testStruct{
				{X: 10, Y: 10},
				{X: 12, Y: 25},
				{X: 14, Y: 40},
			},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{},
				B: func(v testStruct) testStruct {
					return testStruct{X: v.X + 1, Y: v.Y * 2}
				},
			},
			Want: []testStruct{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Map()", slices.Map[testStruct, testStruct], testCases)
}
