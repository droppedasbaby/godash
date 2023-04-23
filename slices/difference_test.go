package slices_test

import (
	"sort"
	"testing"

	"godash/slices"
	"godash/utils"
)

func TestDifferenceWith(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool], []string]{
		{
			Name: "DifferenceWith with empty slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{},
				B: []string{},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{},
		},
		{
			Name: "DifferenceWith with A empty and B not empty",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{},
				B: []string{"d", "e", "f"},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{},
		},
		{
			Name: "DifferenceWith with B empty and A not empty",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: []string{},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{"a", "b", "c"},
		},
		{
			Name: "DifferenceWith with disjoint slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: []string{"d", "e", "f"},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{"a", "b", "c"},
		},
		{
			Name: "DifferenceWith with overlapping slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: []string{"b", "c", "d"},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{"a"},
		},
		{
			Name: "DifferenceWith with identical slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]string, []string, func(string, string) bool]{
				A: []string{"a", "b", "c"},
				B: []string{"a", "b", "c"},
				C: func(a, b string) bool { return a == b },
			},
			Want: []string{},
		},
	}

	utils.RunThreeArgumentTestCases(t, "DifferenceWith()", slices.DifferenceWith[string], testCases)
}

func TestDifferenceHashable(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable], []testHashable]{
		{
			Name: "DifferenceHashable with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: []testHashable{},
			},
			Want: []testHashable{},
		},
		{
			Name: "DifferenceHashable with A empty and B not empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: []testHashable{{"d"}, {"e"}, {"f"}},
			},
			Want: []testHashable{},
		},
		{
			Name: "DifferenceHashable with B empty and A not empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{{"a"}, {"b"}, {"c"}},
				B: []testHashable{},
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}},
		},
		{
			Name: "DifferenceHashable with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{{"a"}, {"b"}, {"c"}},
				B: []testHashable{{"d"}, {"e"}, {"f"}},
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}},
		},
		{
			Name: "DifferenceHashable with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{{"a"}, {"b"}, {"c"}},
				B: []testHashable{{"b"}, {"c"}, {"d"}},
			},
			Want: []testHashable{{"a"}},
		},
		{
			Name: "DifferenceHashable with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{{"a"}, {"b"}, {"c"}},
				B: []testHashable{{"a"}, {"b"}, {"c"}},
			},
			Want: []testHashable{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "DifferenceHashable()",
		func(a []testHashable, b []testHashable) []testHashable {
			got := slices.DifferenceHashable[testHashable](a, b)
			sort.Slice(got, func(i, j int) bool {
				return got[i].Value < got[j].Value
			})
			return got
		}, testCases)
}

func TestDifference(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[[]string, []string], []string]{
		{
			Name: "DifferenceHashable with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{},
				B: []string{},
			},
			Want: []string{},
		},
		{
			Name: "Difference with A empty and B not empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{},
				B: []string{"a", "b", "c"},
			},
			Want: []string{},
		},
		{
			Name: "Difference with B empty and A not empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{},
			},
			Want: []string{"a", "b", "c"},
		},
		{
			Name: "Difference with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"d", "e", "f"},
			},
			Want: []string{"a", "b", "c"},
		},
		{
			Name: "Difference with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"b", "c", "d"},
			},
			Want: []string{"a"},
		},
		{
			Name: "Difference with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"a", "b", "c"},
			},
			Want: []string{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Difference()", func(a []string, b []string) []string {
		got := slices.Difference(a, b)
		sort.Strings(got)
		return got
	}, testCases)
}
