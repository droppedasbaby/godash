package sets_test

import (
	"testing"

	"godash/sets"
	"godash/utils"
)

func TestIsSupersetOf_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			sets.Set[int], sets.Set[int]], bool]{
		{
			Name: "Non-empty set is a superset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{},
			},
			Want: true,
		},
		{
			Name: "Empty set is not a superset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: false,
		},
		{
			Name: "Empty set is a superset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{}, B: sets.Set[int]{},
			},
			Want: true,
		},
		{
			Name: "Superset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true, 3: true}, B: sets.Set[int]{1: true, 2: true},
			},
			Want: true,
		},
		{
			Name: "Not a superset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[int], sets.Set[int]]{
				A: sets.Set[int]{1: true, 2: true}, B: sets.Set[int]{1: true, 2: true, 3: true},
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSupersetOf()", sets.IsSupersetOf[int], testCases)
}

func TestIsSupersetOf_WithString(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.TwoArgumentTestCasesArgsType[
			sets.Set[string], sets.Set[string]], bool]{
		{
			Name: "Non-empty set is a superset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{},
			},
			Want: true,
		},
		{
			Name: "Empty set is not a superset of non-empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{}, B: sets.Set[string]{"a": true, "b": true},
			},
			Want: false,
		},
		{
			Name: "Empty set is a superset of empty set",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{}, B: sets.Set[string]{},
			},
			Want: true,
		},
		{
			Name: "Superset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true, "c": true}, B: sets.Set[string]{"a": true, "b": true},
			},
			Want: true,
		},
		{
			Name: "Not a superset with non-empty sets",
			Args: utils.TwoArgumentTestCasesArgsType[
				sets.Set[string], sets.Set[string]]{
				A: sets.Set[string]{"a": true, "b": true}, B: sets.Set[string]{"a": true, "b": true, "c": true},
			},
			Want: false,
		},
	}

	utils.RunTwoArgumentTestCases(t, "IsSupersetOf()", sets.IsSupersetOf[string], testCases)
}
