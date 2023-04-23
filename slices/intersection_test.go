package slices_test

import (
	"godash/slices"
	"godash/utils"
	"sort"
	"testing"
)

func TestIntersection(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[[]string, []string], []string]{
		{
			Name: "Intersection with both slices empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{},
				B: []string{},
			},
			Want: []string{},
		},
		{
			Name: "Intersection with first slice empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{},
				B: []string{"a", "b", "c"},
			},
			Want: []string{},
		},
		{
			Name: "Intersection with second slice empty",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{},
			},
			Want: []string{},
		},
		{
			Name: "Intersection with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"d", "e", "f"},
			},
			Want: []string{},
		},
		{
			Name: "Intersection with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"b", "c", "d"},
			},
			Want: []string{"b", "c"},
		},
		{
			Name: "Intersection with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]string, []string]{
				A: []string{"a", "b", "c"},
				B: []string{"a", "b", "c"},
			},
			Want: []string{"a", "b", "c"},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Intersection()", func(a []string, b []string) []string {
		result := slices.Intersection(a, b)
		sort.Strings(result)
		return result
	}, testCases)
}

func TestIntersectionHashable(t *testing.T) {
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
			Name: "IntersectionHashable with empty slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: []testHashable{},
			},
			Want: []testHashable{},
		},
		{
			Name: "IntersectionHashable with empty A",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: []testHashable{},
				B: testHashableB,
			},
			Want: []testHashable{},
		},
		{
			Name: "IntersectionHashable with empty B",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: []testHashable{},
			},
			Want: []testHashable{},
		},
		{
			Name: "IntersectionHashable with disjoint slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableB,
			},
			Want: []testHashable{},
		},
		{
			Name: "IntersectionHashable with overlapping slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableC,
			},
			Want: []testHashable{
				{Value: "b"},
				{Value: "c"},
			},
		},
		{
			Name: "IntersectionHashable with identical slices",
			Args: utils.TwoArgumentTestCasesArgsType[[]testHashable, []testHashable]{
				A: testHashableA,
				B: testHashableA,
			},
			Want: []testHashable{
				{Value: "a"},
				{Value: "b"},
				{Value: "c"},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "IntersectionHashable()",
		func(a []testHashable, b []testHashable) []testHashable {
			intersection := slices.IntersectionHashable[testHashable](a, b)
			sort.Slice(intersection, func(i, j int) bool { return intersection[i].Value < intersection[j].Value })
			return intersection
		}, testCases)
}

func TestIntersectionWith(t *testing.T) {
	t.Parallel()

	intsA := []int{1, 2, 3}
	intsB := []int{4, 5, 6}
	intsC := []int{2, 3, 7}

	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool], []int]{
		{
			Name: "IntersectionWith with empty slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: []int{},
				B: []int{},
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{},
		},
		{
			Name: "IntersectionWith with empty A",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: []int{},
				B: intsB,
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{},
		},
		{
			Name: "IntersectionWith with empty B",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: []int{},
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{},
		},
		{
			Name: "IntersectionWith with disjoint slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsB,
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{},
		},
		{
			Name: "IntersectionWith with overlapping slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsC,
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{2, 3},
		},
		{
			Name: "IntersectionWith with identical slices",
			Args: utils.ThreeArgumentTestCasesArgsType[[]int, []int, func(int, int) bool]{
				A: intsA,
				B: intsA,
				C: func(x, y int) bool { return x == y },
			},
			Want: []int{1, 2, 3},
		},
	}

	utils.RunThreeArgumentTestCases(t, "IntersectionWith()", slices.IntersectionWith[int], testCases)
}
