package slices_test

import (
	"math"
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestReduce_WithInt(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[[]int, func(int, int) int, int], int]{
		{
			Name: "sum of integers",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, func(int, int) int, int]{
				A: []int{1, 2, 3, 4, 5},
				B: func(acc, i int) int { return acc + i },
				C: 0,
			},
			Want: 15,
		},
		{
			Name: "product of integers",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, func(int, int) int, int]{
				A: []int{1, 2, 3, 4, 5},
				B: func(acc, i int) int { return acc * i },
				C: 1,
			},
			Want: 120,
		},
		{
			Name: "difference of integers",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, func(int, int) int, int]{
				A: []int{10, 5, 2},
				B: func(acc, i int) int { return acc - i },
				C: 0,
			},
			Want: -17,
		},
		{
			Name: "empty slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, func(int, int) int, int]{
				A: []int{},
				B: func(acc, i int) int { return acc + i },
				C: 0,
			},
			Want: 0,
		},
	}

	utils.RunThreeArgumentTestCases(t, "Reduce()", slices.Reduce[int, int], testCases)
}

func TestReduce_WithFloat64(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]float64, func(float64, float64) float64, float64], float64]{
		{
			Name: "sum of float64s",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, func(float64, float64) float64, float64]{
				A: []float64{1.1, 2.2, 3.3, 4.4, 5.5},
				B: func(acc, f float64) float64 { return acc + f },
				C: 0,
			},
			Want: 16.5,
		},
		{
			Name: "empty slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, func(float64, float64) float64, float64]{
				A: []float64{},
				B: func(acc, f float64) float64 { return acc + f },
				C: 0,
			},
			Want: 0,
		},
	}

	utils.RunThreeArgumentTestCases(t, "Reduce[float64]", slices.Reduce[float64, float64], testCases)
}

func TestReduce_WithTestStruct(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[
			[]testStruct, func(testStruct, testStruct) testStruct, testStruct], testStruct]{
		{
			Name: "sum of x and y fields",
			Args: utils.ThreeArgumentTestCasesArgsType[
				[]testStruct, func(testStruct, testStruct) testStruct, testStruct]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(acc, t testStruct) testStruct {
					return testStruct{x: acc.x + t.x, y: acc.y + t.y}
				},
				C: testStruct{x: 0, y: 0},
			},
			Want: testStruct{x: 9, y: 12},
		},
		{
			Name: "min x and max y fields",
			Args: utils.ThreeArgumentTestCasesArgsType[
				[]testStruct, func(testStruct, testStruct) testStruct, testStruct]{
				A: []testStruct{
					{x: 1, y: 5},
					{x: 3, y: 2},
					{x: 5, y: 9},
				},
				B: func(acc, t testStruct) testStruct {
					return testStruct{
						x: int(math.Min(float64(acc.x), float64(t.x))),
						y: int(math.Max(float64(acc.y), float64(t.y))),
					}
				},
				C: testStruct{x: math.MaxInt32, y: math.MinInt32},
			},
			Want: testStruct{x: 1, y: 9},
		},
		{
			Name: "empty slice",
			Args: utils.ThreeArgumentTestCasesArgsType[
				[]testStruct, func(testStruct, testStruct) testStruct, testStruct]{
				A: []testStruct{},
				B: func(acc, t testStruct) testStruct {
					return testStruct{x: acc.x + t.x, y: acc.y + t.y}
				},
				C: testStruct{x: 0, y: 0},
			},
			Want: testStruct{x: 0, y: 0},
		},
	}

	utils.RunThreeArgumentTestCases(t, "Reduce()", slices.Reduce[testStruct, testStruct], testCases)
}

func TestReduce_WithTestStructToInt(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]testStruct, func(int, testStruct) int, int], int]{
		{
			Name: "empty slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]testStruct, func(int, testStruct) int, int]{
				A: []testStruct{},
				B: func(acc int, v testStruct) int {
					return acc + v.x
				},
				C: 0,
			},
			Want: 0,
		},
		{
			Name: "sum of x fields",
			Args: utils.ThreeArgumentTestCasesArgsType[[]testStruct, func(int, testStruct) int, int]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(acc int, t testStruct) int {
					return acc + t.x
				},
				C: 0,
			},
			Want: 9,
		},
		{
			Name: "sum of y fields",
			Args: utils.ThreeArgumentTestCasesArgsType[[]testStruct, func(int, testStruct) int, int]{
				A: []testStruct{
					{x: 1, y: 2},
					{x: 3, y: 4},
					{x: 5, y: 6},
				},
				B: func(acc int, t testStruct) int {
					return acc + t.y
				},
				C: 0,
			},
			Want: 12,
		},
	}

	utils.RunThreeArgumentTestCases(t, "Reduce()", slices.Reduce[testStruct, int], testCases)
}
