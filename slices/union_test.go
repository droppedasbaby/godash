package slices_test

import (
	"godash/slices"
	"godash/utils"
	"sort"
	"testing"
)

func TestUnion(t *testing.T) {
	t.Parallel()

	intsA := []int{1, 2, 3}
	intsB := []int{4, 5, 6}
	intsC := []int{2, 3, 7}

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[[]int, []int], []int]{
		{
			Name: "Union with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: []int{},
				B: []int{},
			},
			Want: []int{},
		},
		{
			Name: "Union with empty A",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: []int{},
				B: intsB,
			},
			Want: []int{4, 5, 6},
		},
		{
			Name: "Union with empty B",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: intsA,
				B: []int{},
			},
			Want: []int{1, 2, 3},
		},
		{
			Name: "Union with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: intsA,
				B: intsB,
			},
			Want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Name: "Union with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: intsA,
				B: intsC,
			},
			Want: []int{1, 2, 3, 7},
		},
		{
			Name: "Union with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]int, []int]{
				A: intsA,
				B: intsA,
			},
			Want: []int{1, 2, 3},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Union()", func(a []int, b []int) []int {
		union := slices.Union[int](a, b)
		sort.Ints(union)
		return union
	}, testCases)
}

func TestUnionHashable(t *testing.T) {
	t.Parallel()

	testHashableA := []testHashable{
		{Value: "a"},
		{Value: "b"},
		{Value: "c"},
	}
	testHashableB := []testHashable{
		{Value: "d"},
		{Value: "e"},
		{Value: "f"},
	}
	testHashableC := []testHashable{
		{Value: "b"},
		{Value: "c"},
		{Value: "g"},
	}

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable],
		[]testHashable]{
		{
			Name: "UnionHashable with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: []testHashable{},
			},
			Want: []testHashable{},
		},
		{
			Name: "UnionHashable with empty A",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: testHashableB,
			},
			Want: []testHashable{{"d"}, {"e"}, {"f"}},
		},
		{
			Name: "UnionHashable with empty B",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: []testHashable{},
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}},
		},
		{
			Name: "UnionHashable with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableB,
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}, {"d"}, {"e"}, {"f"}},
		},
		{
			Name: "UnionHashable with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableC,
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}, {"g"}},
		},
		{
			Name: "UnionHashable with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableA,
			},
			Want: []testHashable{{"a"}, {"b"}, {"c"}},
		},
	}

	utils.RunTwoArgumentTestCases(t, "UnionHashable()", func(a []testHashable, b []testHashable) []testHashable {
		union := slices.UnionHashable[testHashable](a, b)
		sort.Slice(union, func(i, j int) bool {
			return union[i].Value < union[j].Value
		})
		return union
	}, testCases)
}

func TestUnionWith(t *testing.T) {
	t.Parallel()

	intsA := []int{1, 2, 3}
	intsB := []int{4, 5, 6}
	intsC := []int{2, 3, 7}

	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool],
		[]int]{
		{
			Name: "UnionWith with empty slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: []int{},
				B: []int{},
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{},
		},
		{
			Name: "UnionWith with B empty",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: []int{},
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{1, 2, 3},
		},
		{
			Name: "UnionWith with A empty",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: []int{},
				B: intsB,
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{4, 5, 6},
		},
		{
			Name: "UnionWith with disjoint slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsB,
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Name: "UnionWith with overlapping slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsC,
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{1, 2, 3, 7},
		},
		{
			Name: "UnionWith with identical slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsA,
				C: func(a, b int) bool { return a == b },
			},
			Want: []int{1, 2, 3},
		},
	}

	utils.RunThreeArgumentTestCases(t, "UnionWith()", slices.UnionWith[int], testCases)
}
