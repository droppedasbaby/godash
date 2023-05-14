package slices_test

import (
	"reflect"
	"testing"

	"github.com/GrewalAS/godash"
	"github.com/GrewalAS/godash/slices"
	"github.com/GrewalAS/godash/utils"
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
		a := tc.Input.A
		wantA := tc.WantA
		wantB := tc.WantB
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			gotA, gotB := slices.Unzip[int, string](a)
			if !reflect.DeepEqual(gotA, wantA) || !reflect.DeepEqual(gotB, wantB) {
				t.Errorf("Unzip() = (%v, %v), want (%v, %v)", gotA, gotB, wantA, wantB)
			}
		})
	}
}
