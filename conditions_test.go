// Package godash provides a collection of utility functions for Go,
// inspired by JavaScript's lodash library.
package godash_test

import (
	"testing"

	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/utils"
)

func runTwoArgumentTestCasesForComparableEquality[T comparable](
	t *testing.T, testCases []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[T, T], bool],
) {
	t.Helper()
	utils.RunTwoArgumentTestCases[T](t, "ComparableEquality()", godash.ComparableEquality[T], testCases)
}

func runTestCaseForNumberLessThan[N godash.Number](
	t *testing.T, testCases []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[N, N], bool],
) {
	t.Helper()
	utils.RunTwoArgumentTestCases[N](t, "NumberLessThan()", godash.NumberLessThan[N], testCases)
}

func TestComparableEquality(t *testing.T) {
	t.Parallel()

	testCasesInts := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[int64, int64], bool]{
		{
			Name: "equality of two integers",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[int64, int64]{
				A: 1,
				B: 1,
			},
		},
		{
			Name: "inequality of two integers",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[int64, int64]{
				A: 1,
				B: 2,
			},
		},
	}
	testCasesFloats := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[float64, float64], bool]{
		{
			Name: "equality of two floats",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[float64, float64]{
				A: 1.0,
				B: 1.0,
			},
		},
		{
			Name: "inequality of two floats",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[float64, float64]{
				A: 1.0,
				B: 2.0,
			},
		},
	}
	testCasesStrings := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[string, string], bool]{
		{
			Name: "equality of two strings",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[string, string]{
				A: "hello",
				B: "hello",
			},
		},
		{
			Name: "inequality of two strings",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[string, string]{
				A: "hello",
				B: "world",
			},
		},
	}

	runTwoArgumentTestCasesForComparableEquality(t, testCasesInts)
	runTwoArgumentTestCasesForComparableEquality(t, testCasesFloats)
	runTwoArgumentTestCasesForComparableEquality(t, testCasesStrings)
}

func TestNumberLessThan(t *testing.T) {
	t.Parallel()

	testCasesInts := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[int, int], bool]{
		{
			Name: "less than of two integers",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[int, int]{
				A: 1,
				B: 2,
			},
		},
		{
			Name: "not less than of two integers",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[int, int]{
				A: 2,
				B: 1,
			},
		},
	}

	testCasesFloats := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[float64, float64], bool]{
		{
			Name: "less than of two floats",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[float64, float64]{
				A: 1.0,
				B: 2.0,
			},
		},
		{
			Name: "not less than of two floats",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[float64, float64]{
				A: 2.0,
				B: 1.0,
			},
		},
	}

	runTestCaseForNumberLessThan(t, testCasesInts)
	runTestCaseForNumberLessThan(t, testCasesFloats)
}

type hashableTestType struct {
	value int
}

func (h hashableTestType) Hash() string {
	return string(rune(h.value))
}

func TestHashableEquality(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[hashableTestType, hashableTestType], bool]{
		{
			Name: "equality of two hashable types",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[hashableTestType, hashableTestType]{
				A: hashableTestType{value: 1},
				B: hashableTestType{value: 1},
			},
		},
		{
			Name: "inequality of two hashable types",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[hashableTestType, hashableTestType]{
				A: hashableTestType{value: 1},
				B: hashableTestType{value: 2},
			},
		},
	}

	utils.RunTwoArgumentTestCases[hashableTestType](
		t,
		"HashableEquality()",
		godash.HashableEquality[hashableTestType, hashableTestType],
		testCases,
	)
}

func TestAnyEquality(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, any], bool]{
		{
			Name: "equality of two int collections types",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: []int{1, 2, 3},
				B: []int{1, 2, 3},
			},
		},
		{
			Name: "inequality of two int collections types",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: []int{1, 2, 3},
				B: []int{1, 2, 4},
			},
		},
		{
			Name: "equality of two string collections types",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: []string{"hello", "world"},
				B: []string{"hello", "world"},
			},
		},
		{
			Name: "inequality of two string collections types",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: []string{"hello", "world"},
				B: []string{"hello", "world", "again"},
			},
		},
		{
			Name: "equality of two map types",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: map[string]int{"hello": 1, "world": 2},
				B: map[string]int{"hello": 1, "world": 2},
			},
		},
		{
			Name: "inequality of two map types",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: map[string]int{"hello": 1, "world": 2},
				B: map[string]int{"hello": 1, "world": 3},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AnyEquality()", godash.AnyEquality[any], testCases)
}
