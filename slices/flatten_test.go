package slices_test

import (
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestFlatten(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[][]string],
		[]string]{
		{
			Name: "Flatten with multiple nested slices",
			Args: utils.SingleArgumentTestCasesArgsType[[][]string]{
				A: [][]string{{"a", "b", "c"}, {"d", "e"}, {"f"}},
			},
			Want: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			Name: "Flatten with some empty nested slices",
			Args: utils.SingleArgumentTestCasesArgsType[[][]string]{
				A: [][]string{{"a", "b", "c"}, {}, {"d", "e"}, {}, {"f"}},
			},
			Want: []string{"a", "b", "c", "d", "e", "f"},
		},
		{
			Name: "Flatten with all empty nested slices",
			Args: utils.SingleArgumentTestCasesArgsType[[][]string]{
				A: [][]string{{}, {}, {}},
			},
			Want: []string{},
		},
		{
			Name: "Flatten with an empty input slice",
			Args: utils.SingleArgumentTestCasesArgsType[[][]string]{
				A: [][]string{},
			},
			Want: []string{},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Flatten()", slices.Flatten[string], testCases)
}

func TestFlatten3DArrayWithUtilities(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[[][][]string], []string]{
		{
			Name: "Flatten 3D array into 1D array using Flatten twice",
			Args: utils.SingleArgumentTestCasesArgsType[[][][]string]{
				A: [][][]string{
					{
						{"a", "b"},
						{"c", "d"},
					},
					{
						{"e", "f"},
						{"g", "h"},
					},
				},
			},
			Want: []string{"a", "b", "c", "d", "e", "f", "g", "h"},
		},
	}

	utils.RunSingleArgumentTestCases(t, "flattenTwice()", func(input [][][]string) []string {
		return slices.Flatten(slices.Flatten(input))
	}, testCases)
}
