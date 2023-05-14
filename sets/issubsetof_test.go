package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestIsSubsetOf_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[int], sets.Set[int]], bool]{
		{
			Name: "Empty set is a subset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: true,
		},
		{
			Name: "Non-empty set is not a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{},
			},
			Want: false,
		},
		{
			Name: "Empty set is a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{}, B: sets.Set[int]{},
			},
			Want: true,
		},
		{
			Name: "Subset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{1: true, 2: true, 3: true},
			},
			Want: true,
		},
		{
			Name: "Not a subset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true, 3: true}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: false,
		},
		{
			Name: "Set is a subset of itself",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: true,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSubsetOf()", sets.IsSubsetOf[int], testCases)
}

func TestIsSubsetOf_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], sets.Set[string]], bool]{
		{
			Name: "Empty set is a subset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{A: sets.Set[string]{}, B: sets.Set[string]{"a": true, "b": true}},
			Want: true,
		},
		{
			Name: "Non-empty set is not a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{},
			},
			Want: false,
		},
		{
			Name: "Empty set is a subset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{}, B: sets.Set[string]{},
			},
			Want: true,
		},
		{
			Name: "Subset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{"a": true, "b": true, "c": true},
			},
			Want: true,
		},
		{
			Name: "Not a subset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: sets.Set[string]{"a": true, "b": true},
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSubsetOf()", sets.IsSubsetOf[string], testCases)
}
