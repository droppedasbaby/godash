package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestUnion_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[int], sets.Set[int]], sets.Set[int]]{
		{
			Name: "Union of two non-empty sets with no common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{3: true, 4: true},
			},
			Want: sets.Set[int]{1: true, 2: true, 3: true, 4: true},
		},
		{
			Name: "Union of two non-empty sets with common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true, 3: true}, B: sets.Set[int]{2: true, 3: true, 4: true},
			},
			Want: sets.Set[int]{1: true, 2: true, 3: true, 4: true},
		},
		{
			Name: "Union of a non-empty set and an empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{}},
			Want: sets.Set[int]{1: true, 2: true},
		},
		{
			Name: "Union of two empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{A: sets.Set[int]{}, B: sets.Set[int]{}},
			Want: sets.Set[int]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Union()", sets.Union[int], testCases)
}

func TestUnion_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], sets.Set[string]], sets.Set[string]]{
		{
			Name: "Union of two non-empty sets with no common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{"c": true, "d": true},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true, "d": true},
		},
		{
			Name: "Union of two non-empty sets with common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: sets.Set[string]{"b": true, "c": true, "d": true},
			},
			Want: sets.Set[string]{"a": true, "b": true, "c": true, "d": true},
		},
		{
			Name: "Union of a non-empty set and an empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{}},
			Want: sets.Set[string]{"a": true, "b": true},
		},
		{
			Name: "Union of two empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{A: sets.Set[string]{}, B: sets.Set[string]{}},
			Want: sets.Set[string]{},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Union()", sets.Union[string], testCases)
}
