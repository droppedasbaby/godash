package slices_test

import (
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
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

func TestMap_WithTestStruct(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct], []testStruct]{
		{
			Name: "increment x and double y",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{
					{x: 1, y: 10},
					{x: 2, y: 20},
					{x: 3, y: 30},
				},
				B: func(v testStruct) testStruct {
					return testStruct{x: v.x + 1, y: v.y * 2}
				},
			},
			Want: []testStruct{
				{x: 2, y: 20},
				{x: 3, y: 40},
				{x: 4, y: 60},
			},
		},
		{
			Name: "double x and subtract y by 5",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{
					{x: 5, y: 15},
					{x: 6, y: 30},
					{x: 7, y: 45},
				},
				B: func(v testStruct) testStruct {
					return testStruct{x: v.x * 2, y: v.y - 5}
				},
			},
			Want: []testStruct{
				{x: 10, y: 10},
				{x: 12, y: 25},
				{x: 14, y: 40},
			},
		},
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) testStruct]{
				A: []testStruct{},
				B: func(v testStruct) testStruct {
					return testStruct{x: v.x + 1, y: v.y * 2}
				},
			},
			Want: []testStruct{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Map()", slices.Map[testStruct, testStruct], testCases)
}

func TestMap_WithTestStructToInt(t *testing.T) {
	t.Parallel()
	testStructTestCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) int], []int]{
		{
			Name: "empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) int]{
				A: []testStruct{},
				B: func(v testStruct) int { return v.x },
			},
			Want: []int{},
		},
		{
			Name: "multiply x and y fields",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) int]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(v testStruct) int { return v.x * v.y },
			},
			Want: []int{2, 12, 30},
		},
		{
			Name: "subtract x from y fields",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) int]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(v testStruct) int { return v.y - v.x },
			},
			Want: []int{1, 1, 1},
		},
		{
			Name: "sum of x and y fields",
			Args: utils.TwoArgumentTestCasesArgsType[[]testStruct, func(testStruct) int]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(v testStruct) int { return v.x + v.y },
			},
			Want: []int{3, 7, 11},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Map()", slices.Map[testStruct, int], testStructTestCases)
}
