package slices_test

import (
	"reflect"
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
)

func TestShuffle(t *testing.T) {
	t.Parallel()

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.SingleArgumentTestCasesArgsType[[]int]]{
		Name: "Shuffle with slice of size 1",
		Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{5}},
	}

	testCaseSize10 := utils.GenericTestCaseWithNoWant[utils.SingleArgumentTestCasesArgsType[[]int]]{
		Name: "Shuffle with slice of size 10",
		Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.Shuffle(testCaseSize1.Args.A)
		if len(testCaseSize1.Args.A) == 1 && reflect.DeepEqual(got, testCaseSize1.Args.A[0]) {
			t.Errorf("expected %v, but got %v", testCaseSize1.Args.A[0], got)
		}
	})
	t.Run(testCaseSize10.Name, func(t *testing.T) {
		t.Parallel()
		const numTrials = 100000
		for i := 0; i < numTrials; i++ {
			got := slices.Shuffle(testCaseSize10.Args.A)
			if !utils.TestingSlicesEqualWithoutOrder(got, testCaseSize10.Args.A) {
				t.Errorf("expected %v, but got %v", testCaseSize10.Args.A, got)
			}
		}
	})
}

func TestShuffleWithSeed(t *testing.T) {
	t.Parallel()
	seed := int64(42)

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int64]]{
		Name: "ShuffleWithSeed with slice of size 1",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int64]{A: []int{5}, B: seed},
	}

	testCaseSize10 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int64]]{
		Name: "ShuffleWithSeed with slice of size 10",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int64]{A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, B: seed},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.ShuffleWithSeed(testCaseSize1.Args.A, testCaseSize1.Args.B)
		if len(testCaseSize1.Args.A) == 1 && reflect.DeepEqual(got, testCaseSize1.Args.A[0]) {
			t.Errorf("expected %v, but got %v", testCaseSize1.Args.A[0], got)
		}
	})
	t.Run(testCaseSize10.Name, func(t *testing.T) {
		t.Parallel()
		const numTrials = 100000
		for i := 0; i < numTrials; i++ {
			got := slices.ShuffleWithSeed(testCaseSize10.Args.A, testCaseSize10.Args.B)
			if !utils.TestingSlicesEqualWithoutOrder(got, testCaseSize10.Args.A) {
				t.Errorf("expected %v, but got %v", testCaseSize10.Args.A, got)
			}
		}
	})
}
