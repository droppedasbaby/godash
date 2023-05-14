// Package sets provides functionality for creating and manipulating
// set data structures in Go.
package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestContains_Int(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[sets.Set[int], int], bool]{
		{
			Name: "Element exists in set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{1: true, 2: true, 3: true}, B: 1},
			Want: true,
		},
		{
			Name: "Element does not exist in set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{1: true, 2: true, 3: true}, B: 4},
			Want: false,
		},
		{
			Name: "Element does not exist in empty set",
			Args: utils.TwoArgumentTestCasesArgsType[sets.Set[int], int]{A: sets.Set[int]{}, B: 1},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", sets.Contains[int], testCases)
}

func TestContains_String(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[sets.Set[string], string], bool]{
		{
			Name: "Element exists in set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: "a",
			},
			Want: true,
		},
		{
			Name: "Element does not exist in set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: "d",
			},
			Want: false,
		},
		{
			Name: "Element does not exist in empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], string]{
				A: sets.Set[string]{}, B: "a",
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "Contains()", sets.Contains[string], testCases)
}
