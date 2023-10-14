package search_test

import (
	"github.com/GrewalAS/godash/search"
	"github.com/GrewalAS/godash/utils"
	"testing"
)

func TestBisectLeftWith_Int(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool], int]{
		{
			Name: "Empty Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{3}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Not Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{3}, B: 4, C: func(a, b int) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is First Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 1, C: func(a, b int) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Target Is Last Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 5, C: func(a, b int) bool { return a < b },
			},
			Want: 4,
		},
		{
			Name: "Target In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 2,
		},
		{
			Name: "Target Not In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 6, C: func(a, b int) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target With Duplicates",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 1, 2, 2, 3, 3}, B: 2, C: func(a, b int) bool { return a < b },
			},
			Want: 2,
		},
	}
	utils.RunThreeArgumentTestCases(t, "BisectLeftWith", search.BisectLeftWith[int], testCases)
}

func TestBisectLeftWith_Float(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool], int]{
		{
			Name: "Empty Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{3.0}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Not Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{3.0}, B: 4.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is First Element In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 1.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Target Is Last Element In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 5.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 4,
		},
		{
			Name: "Target In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 2,
		},
		{
			Name: "Target Not In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 6.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target With Duplicates",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 1.0, 2.0, 2.0, 3.0, 3.0}, B: 2.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 2,
		},
	}
	utils.RunThreeArgumentTestCases(t, "BisectLeftWith", search.BisectLeftWith[float64], testCases)
}

func TestBisectRightWith_Int(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool], int]{
		{
			Name: "Empty Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{3}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Single Element Slice, Target Not Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{3}, B: 4, C: func(a, b int) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is First Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 1, C: func(a, b int) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is Last Element In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 5, C: func(a, b int) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 3, C: func(a, b int) bool { return a < b },
			},
			Want: 3,
		},
		{
			Name: "Target Not In Int Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 2, 3, 4, 5}, B: 6, C: func(a, b int) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target With Duplicates",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 1, 2, 2, 3, 3}, B: 2, C: func(a, b int) bool { return a < b },
			},
			Want: 4,
		},
	}
	utils.RunThreeArgumentTestCases(t, "BisectRightWith", search.BisectRightWith[int], testCases)
}

func TestBisectRightWith_Float(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool], int]{
		{
			Name: "Empty Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 0,
		},
		{
			Name: "Single Element Slice, Target Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{3.0}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Single Element Slice, Target Not Found",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{3.0}, B: 4.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is First Element In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 1.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 1,
		},
		{
			Name: "Target Is Last Element In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 5.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 3.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 3,
		},
		{
			Name: "Target Not In Float Slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 2.0, 3.0, 4.0, 5.0}, B: 6.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 5,
		},
		{
			Name: "Target With Duplicates",
			Args: utils.ThreeArgumentTestCasesArgsType[[]float64, float64, func(float64, float64) bool]{
				A: []float64{1.0, 1.0, 2.0, 2.0, 3.0, 3.0}, B: 2.0, C: func(a, b float64) bool { return a < b },
			},
			Want: 4,
		},
	}
	utils.RunThreeArgumentTestCases(t, "BisectRightWith", search.BisectRightWith[float64], testCases)
}

func TestBisectWith(t *testing.T) {
	// BisectWith is essentially calling BisectRightWith
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool], int]{
		{
			Name: "Target With Duplicates",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, func(int, int) bool]{
				A: []int{1, 1, 2, 2, 3, 3}, B: 2, C: func(a, b int) bool { return a < b },
			},
			Want: 4,
		},
	}
	utils.RunThreeArgumentTestCases(t, "BisectWith", search.BisectWith[int], testCases)
}
