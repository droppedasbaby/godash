package sets_test

import (
	"testing"

	"github.com/GrewalAS/godash/sets"
	"github.com/GrewalAS/godash/utils"
)

func TestDifference_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[int], sets.Set[int]], sets.Set[int]]{
		{
			Name: "Non-empty sets with no common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{3: true, 4: true},
			},
			Want: sets.Set[int]{1: true, 2: true},
		},
		{
			Name: "Non-empty sets with some common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true, 3: true}, B: sets.Set[int]{2: true, 3: true, 4: true},
			},
			Want: sets.Set[int]{1: true},
		},
		{
			Name: "Non-empty sets with all common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: sets.Set[int]{},
		},
		{
			Name: "First set empty",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: sets.Set[int]{},
		},
		{
			Name: "Second set empty",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{},
			},
			Want: sets.Set[int]{1: true, 2: true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Difference()", sets.Difference[int], testCases)
}

func TestDifference_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[sets.Set[string], sets.Set[string]], sets.Set[string]]{
		{
			Name: "Non-empty sets with no common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{"c": true, "d": true},
			},
			Want: sets.Set[string]{"a": true, "b": true},
		},
		{
			Name: "Non-empty sets with some common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: sets.Set[string]{"b": true, "c": true, "d": true},
			},
			Want: sets.Set[string]{"a": true},
		},
		{
			Name: "Non-empty sets with all common elements",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{"a": true, "b": true},
			},
			Want: sets.Set[string]{},
		},
		{
			Name: "First set empty",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{}, B: sets.Set[string]{"a": true, "b": true},
			},
			Want: sets.Set[string]{},
		},
		{
			Name: "Second set empty",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{},
			},
			Want: sets.Set[string]{"a": true, "b": true},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Difference()", sets.Difference[string], testCases)
}
