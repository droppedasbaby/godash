package slices_test

import (
	"godash"
	"godash/slices"
	"godash/utils"
	"reflect"
	"testing"
)

func TestUnzip(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		Name  string
		Input utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]]
		WantA []int
		WantB []string
	}{
		{
			Name: "Unzip with non-empty slice of pairs",
			Input: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]]{
				A: []godash.Pair[int, string]{
					{First: 1, Second: "a"},
					{First: 2, Second: "b"},
					{First: 3, Second: "c"},
				},
			},
			WantA: []int{1, 2, 3},
			WantB: []string{"a", "b", "c"},
		},
		{
			Name: "Unzip with empty slice of pairs",
			Input: utils.SingleArgumentTestCasesArgsType[[]godash.Pair[int, string]]{
				A: []godash.Pair[int, string]{},
			},
			WantA: []int{},
			WantB: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			gotA, gotB := slices.Unzip[int, string](tc.Input.A)
			if !reflect.DeepEqual(gotA, tc.WantA) || !reflect.DeepEqual(gotB, tc.WantB) {
				t.Errorf("Unzip() = (%v, %v), want (%v, %v)", gotA, gotB, tc.WantA, tc.WantB)
			}
		})
	}
}
