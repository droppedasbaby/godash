package slices_test

import (
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestContains(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]string, string],
		bool]{
		{
			Name: "Contains returns true for existing element",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, string]{
				A: []string{"a", "b", "c"},
				B: "b",
			},
			Want: true,
		},
		{
			Name: "Contains returns false for non-existing element",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, string]{
				A: []string{"a", "b", "c"},
				B: "d",
			},
			Want: false,
		},
		{
			Name: "Contains returns false for empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, string]{
				A: []string{},
				B: "d",
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", slices.Contains[string], testCases)
}

func TestContainsWith(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool],
		bool]{
		{
			Name: "ContainsWith returns true for existing element",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: "b",
				C: func(a, b string) bool { return a == b },
			},
			Want: true,
		},
		{
			Name: "ContainsWith returns false for non-existing element",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: "d",
				C: func(a, b string) bool { return a == b },
			},
			Want: false,
		},
		{
			Name: "ContainsWith returns false for empty slice",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, string, func(string, string) bool]{
				A: []string{},
				B: "d",
				C: func(a, b string) bool { return a == b },
			},
		},
	}

	utils.RunThreeArgumentTestCases(t, "ContainsWith()", slices.ContainsWith[string, string], testCases)
}

func TestContainsHashable(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testHashable, testHashable],
		bool]{
		{
			Name: "ContainsHashable returns true for existing element",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, testHashable]{
				A: []testHashable{{Value: "a"}, {Value: "b"}, {Value: "c"}},
				B: testHashable{Value: "b"},
			},
			Want: true,
		},
		{
			Name: "ContainsHashable returns false for non-existing element",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, testHashable]{
				A: []testHashable{{Value: "a"}, {Value: "b"}, {Value: "c"}},
				B: testHashable{Value: "d"},
			},
			Want: false,
		},
		{
			Name: "ContainsHashable returns false for empty slice",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, testHashable]{
				A: []testHashable{},
				B: testHashable{Value: "d"},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "ContainsHashable()", slices.ContainsHashable[testHashable], testCases)
}
