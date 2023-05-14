// Package maps provides functions for manipulating maps in Go.
package maps_test

import (
	"testing"

	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/maps"
	"github.com/GrewalAS/godash/utils"
)

func TestAddAll(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[map[string]int, []godash.Pair[string, int]], map[string]int]{
		{
			Name: "Empty Map",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[string]int, []godash.Pair[string, int]]{
				A: map[string]int{}, B: []godash.Pair[string, int]{{First: "a", Second: 1}},
			},
			Want: map[string]int{"a": 1},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[string]int, []godash.Pair[string, int]]{
				A: map[string]int{"a": 1}, B: []godash.Pair[string, int]{{First: "b", Second: 2}},
			},
			Want: map[string]int{"a": 1, "b": 2},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()", func(m map[string]int, ps []godash.Pair[string, int]) map[string]int {
		maps.AddAll(&m, ps...)
		return m
	}, testCases)
}

func TestAddAll_WithInts(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[map[int]string, []godash.Pair[int, string]], map[int]string]{
		{
			Name: "Empty Map",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[int]string, []godash.Pair[int, string]]{
				A: map[int]string{}, B: []godash.Pair[int, string]{{First: 1, Second: "a"}},
			},
			Want: map[int]string{1: "a"},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.TwoArgumentTestCasesArgsType[
				map[int]string, []godash.Pair[int, string]]{
				A: map[int]string{1: "a"}, B: []godash.Pair[int, string]{{First: 2, Second: "b"}},
			},
			Want: map[int]string{1: "a", 2: "b"},
		},
	}

	utils.RunTwoArgumentTestCases(t, "AddAll()",
		func(m map[int]string, ps []godash.Pair[int, string]) map[int]string {
			maps.AddAll(&m, ps...)
			return m
		}, testCases)
}
