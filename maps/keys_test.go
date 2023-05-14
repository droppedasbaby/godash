package maps_test

import (
	"sort"
	"testing"

	"github.com/GrewalAS/godash/utils"
)

func TestKeys_WithIntString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[map[int]string], []int]{
		{
			Name: "keys of empty map",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{
				A: map[int]string{},
			},
			Want: []int{},
		},
		{
			Name: "keys of map with one element",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{
				A: map[int]string{1: "one"},
			},
			Want: []int{1},
		},
		{
			Name: "keys of map with multiple elements",
			Args: utils.SingleArgumentTestCasesArgsType[map[int]string]{
				A: map[int]string{1: "one", 2: "two", 3: "three"},
			},
			Want: []int{1, 2, 3},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Keys()", func(m map[int]string) []int {
		ks := maps.Keys(m)
		sort.Ints(ks)
		return ks
	}, testCases)
}

func TestKeys_WithStringFloat64(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[map[string]float64], []string]{
		{
			Name: "keys of empty map",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]float64]{
				A: map[string]float64{},
			},
			Want: []string{},
		},
		{
			Name: "keys of map with one element",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]float64]{
				A: map[string]float64{"one": 1.0},
			},
			Want: []string{"one"},
		},
		{
			Name: "keys of map with multiple elements",
			Args: utils.SingleArgumentTestCasesArgsType[map[string]float64]{
				A: map[string]float64{"one": 1.0, "two": 2.0, "three": 3.0},
			},
			Want: []string{"one", "three", "two"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Keys()", func(m map[string]float64) []string {
		ks := maps.Keys(m)
		sort.Strings(ks)
		return ks
	}, testCases)
}
