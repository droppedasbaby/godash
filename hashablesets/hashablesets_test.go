package hashablesets_test

import (
	"sort"
	"testing"

	"godash/hashablesets"
	"godash/utils"
)

func TestFromSet(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable]],
		[]TestHashable]{
		{
			Name: "Empty set returns empty slice",
			Args: utils.SingleArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{},
			},
			Want: []TestHashable{},
		},
		{
			Name: "Single element set returns single element slice",
			Args: utils.SingleArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{"a": TestHashable{Value: "a"}},
			},
			Want: []TestHashable{{"a"}},
		},
		{
			Name: "Multiple element set returns multiple element slice",
			Args: utils.SingleArgumentTestCasesArgsType[hashablesets.HashableSet[TestHashable]]{
				A: hashablesets.HashableSet[TestHashable]{
					"a": TestHashable{"a"}, "b": TestHashable{"b"}, "c": TestHashable{"c"},
				},
			},
			Want: []TestHashable{{"a"}, {"b"}, {"c"}},
		},
	}

	utils.RunSingleArgumentTestCases(t, "FromSet()",
		func(hs hashablesets.HashableSet[TestHashable]) []TestHashable {
			arr := hashablesets.FromSet[TestHashable](hs)
			sort.Slice(arr, func(i, j int) bool { return arr[i].Value < arr[j].Value })
			return arr
		}, testCases)
}

func TestToSet(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[[]TestHashable],
		hashablesets.HashableSet[TestHashable]]{
		{
			Name: "Empty slice returns empty set",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{},
			},
			Want: hashablesets.HashableSet[TestHashable]{},
		},
		{
			Name: "Single element slice returns single element set",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{{"a"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{"a": TestHashable{"a"}},
		},
		{
			Name: "Multiple element slice returns multiple element set",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{{"a"}, {"b"}, {"c"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
				"c": TestHashable{"c"},
			},
		},
		{
			Name: "Duplicate elements in slice are deduplicated in set",
			Args: utils.SingleArgumentTestCasesArgsType[[]TestHashable]{
				A: []TestHashable{{"a"}, {"b"}, {"a"}},
			},
			Want: hashablesets.HashableSet[TestHashable]{
				"a": TestHashable{"a"},
				"b": TestHashable{"b"},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "ToSet()", hashablesets.ToSet[TestHashable], testCases)
}
