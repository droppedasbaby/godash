package slices_test

import (
	"reflect"
	"testing"

	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
)

func TestSample(t *testing.T) {
	t.Parallel()

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.SingleArgumentTestCasesArgsType[[]int]]{
		Name: "Sample with slice of size 1",
		Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{5}},
	}

	testCaseSize5 := utils.GenericTestCaseWithNoWant[utils.SingleArgumentTestCasesArgsType[[]int]]{
		Name: "Sample with slice of size 5 for distribution",
		Args: utils.SingleArgumentTestCasesArgsType[[]int]{A: []int{1, 2, 3, 4, 5}},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.Sample(testCaseSize1.Args.A)
		if len(testCaseSize1.Args.A) == 1 && got != testCaseSize1.Args.A[0] {
			t.Errorf("expected %v, but got %v", testCaseSize1.Args.A[0], got)
		}
	})
	t.Run(testCaseSize5.Name, func(t *testing.T) {
		t.Parallel()
		counts := make(map[int]int)
		const numTrials = 100000
		for i := 0; i < numTrials; i++ {
			got := slices.Sample(testCaseSize5.Args.A)
			counts[got]++
		}
		avg := numTrials / len(testCaseSize5.Args.A)
		for val, count := range counts {
			if deviation := float64(count-avg) / float64(avg); deviation > 0.05 {
				t.Errorf("value %v has uneven distribution: expected around %v, got %v", val, avg, count)
			}
		}
	})
}

func TestSampleWithSeed(t *testing.T) {
	t.Parallel()
	seed := int64(42)

	testCaseSize1 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int64]]{
		Name: "SampleWithSeed with slice of size 1",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int64]{A: []int{5}, B: seed},
	}

	testCaseSize10 := utils.GenericTestCaseWithNoWant[utils.TwoArgumentTestCasesArgsType[[]int, int64]]{
		Name: "SampleWithSeed with slice of integers and deterministic output",
		Args: utils.TwoArgumentTestCasesArgsType[[]int, int64]{A: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, B: seed},
	}

	t.Run(testCaseSize1.Name, func(t *testing.T) {
		t.Parallel()
		got := slices.SampleWithSeed(testCaseSize1.Args.A, testCaseSize1.Args.B)
		if got != testCaseSize1.Args.A[0] {
			t.Errorf("expected an item from the slice, but got %v", got)
		}
	})
	t.Run(testCaseSize10.Name, func(t *testing.T) {
		t.Parallel()
		const numTrials = 1000
		firstRunResults := slices.SampleWithSeed(testCaseSize10.Args.A, testCaseSize10.Args.B)
		for i := 0; i < numTrials; i++ {
			got := slices.SampleWithSeed(testCaseSize10.Args.A, testCaseSize10.Args.B)
			if !reflect.DeepEqual(got, firstRunResults) {
				t.Errorf("expected %v, but got %v", firstRunResults, got)
			}
		}
	})
}
