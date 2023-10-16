package slices_test

import (
	"reflect"
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
)

func TestSampleN(t *testing.T) {
	t.Parallel()
	n := 3

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int]]{
		Name: "Sample with slice of size 1",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{5}, B: 1},
	}

	testCaseSize5 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int]]{
		Name: "Sample with slice of size 5 for distribution",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int]{A: []int{1, 2, 3, 4, 5}, B: n},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.SampleN(testCaseSize1.Args.A, testCaseSize1.Args.B)
		if len(testCaseSize1.Args.A) == 1 && reflect.DeepEqual(got, testCaseSize1.Args.A[0]) {
			t.Errorf("expected %v, but got %v", testCaseSize1.Args.A[0], got)
		}
	})

	t.Run(testCaseSize5.Name, func(t *testing.T) {
		t.Parallel()
		counts := make(map[int]int)
		const numTrials = 100000
		for i := 0; i < numTrials; i++ {
			got := slices.SampleN(testCaseSize5.Args.A, testCaseSize5.Args.B)
			if len(got) != n {
				t.Errorf("expected %v, but got %v", n, len(got))
			}
			for _, val := range got {
				counts[val]++
			}
		}
		avg := (numTrials * 3) / len(testCaseSize5.Args.A)
		for val, count := range counts {
			if deviation := float64(count-avg) / float64(avg); deviation > 0.05 {
				t.Errorf("value %v has uneven distribution: expected around %v, got %v", val, avg*3, count)
			}
		}
	})
}

func TestSampleNWithSeed(t *testing.T) {
	t.Parallel()
	seed := int64(42)

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.ThreeArgumentTestCasesArgsType[[]int, int, int64]]{
		Name: "SampleWithSeed with slice of size 1",
		Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, int64]{A: []int{5}, B: 1, C: seed},
	}

	testCaseSize10 := utils.GenericTestCaseWithNoWant[utils.ThreeArgumentTestCasesArgsType[[]int, int, int64]]{
		Name: "SampleWithSeed with slice of integers and deterministic output",
		Args: utils.ThreeArgumentTestCasesArgsType[[]int, int, int64]{A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, B: 3, C: seed},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.SampleNWithSeed(testCaseSize1.Args.A, testCaseSize1.Args.B, testCaseSize1.Args.C)
		if reflect.DeepEqual(got, testCaseSize1.Args.A[0]) {
			t.Errorf("expected an item from the slice, but got %v", got)
		}
	})

	t.Run(testCaseSize10.Name, func(t *testing.T) {
		t.Parallel()
		const numTrials = 1000
		firstRunResults := slices.SampleNWithSeed(testCaseSize10.Args.A, testCaseSize10.Args.B, testCaseSize10.Args.C)
		for i := 0; i < numTrials; i++ {
			got := slices.SampleNWithSeed(testCaseSize10.Args.A, testCaseSize10.Args.B, testCaseSize10.Args.C)
			if !reflect.DeepEqual(got, firstRunResults) {
				t.Errorf("For seed %v, on run %v, expected %v, but got %v", seed, i+1, firstRunResults[i], got)
			}
		}
	})
}
