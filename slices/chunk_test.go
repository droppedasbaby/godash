package slices_test

import (
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
)

func TestChunk(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]int, int],
		[][]int]{
		{
			Name: "Empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{}, B: 3},
			Want: [][]int{},
		},
		{
			Name: "Single element slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1}, B: 3},
			Want: [][]int{{1}},
		},
		{
			Name: "Evenly divisible slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2, 3, 4, 5, 6}, B: 3},
			Want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			Name: "Not evenly divisible slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2, 3, 4, 5, 6, 7}, B: 3},
			Want: [][]int{{1, 2, 3}, {4, 5, 6}, {7}},
		},
		{
			Name: "Slice smaller than chunk size",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2}, B: 3},
			Want: [][]int{{1, 2}},
		},
		{
			Name: "Chunk size 1",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2, 3, 4}, B: 1},
			Want: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			Name: "Chunk size equal to slice length",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2, 3, 4}, B: 4},
			Want: [][]int{{1, 2, 3, 4}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Chunk()", slices.Chunk[int], testCases)
}
