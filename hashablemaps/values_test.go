package hashablemaps_test

import (
	"sort"
	"testing"

	"godash/hashablemaps"
	"godash/utils"
)

func TestValues(t *testing.T) {
	t.Parallel()
	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]], []TestHashable]{
		{
			Name: "Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{},
			},
			Want: []TestHashable{},
		},
		{
			Name: "Non-Empty Map",
			Args: utils.SingleArgumentTestCasesArgsType[hashablemaps.HashableMap[TestHashable]]{
				A: hashablemaps.HashableMap[TestHashable]{"a": TestHashable{"a"}, "b": TestHashable{"b"}},
			},
			Want: []TestHashable{TestHashable{"a"}, TestHashable{"b"}},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Values()",
		func(m hashablemaps.HashableMap[TestHashable]) []TestHashable {
			vals := hashablemaps.Values[TestHashable](m)
			sort.Slice(vals, func(i, j int) bool { return vals[i].Value < vals[j].Value })
			return vals
		}, testCases)
}
