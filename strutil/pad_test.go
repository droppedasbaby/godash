package strutil_test

import (
	"testing"

	"github.com/GrewalAS/godash/strutil"
	"github.com/GrewalAS/godash/utils"
)

func TestPadLeft(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[string, int, string], string]{
		{
			Name: "Pad single character left",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: "0"},
			Want: "00abc",
		},
		{
			Name: "Pad multiple characters left",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 6, C: "xy"},
			Want: "xyxabc",
		},
		{
			Name: "Original string longer than padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 2, C: "0"},
			Want: "abc",
		},
		{
			Name: "Empty original string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 2, C: "0"},
			Want: "00",
		},
		{
			Name: "Empty padding string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: ""},
			Want: "abc",
		},
		{
			Name: "Empty input and padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 0, C: ""},
			Want: "",
		},
		{
			Name: "Negative length",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: -1, C: "0"},
			Want: "abc",
		},
	}

	utils.RunThreeArgumentTestCases(t, "PadLeft", strutil.PadLeft, testCases)
}

func TestPadRight(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[string, int, string], string]{
		{
			Name: "Pad single character right",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: "0"},
			Want: "abc00",
		},
		{
			Name: "Pad multiple characters right",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 6, C: "xy"},
			Want: "abcxyx",
		},
		{
			Name: "Original string longer than padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 2, C: "0"},
			Want: "abc",
		},
		{
			Name: "Empty original string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 2, C: "0"},
			Want: "00",
		},
		{
			Name: "Empty padding string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: ""},
			Want: "abc",
		},
		{
			Name: "Empty input and padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 0, C: ""},
			Want: "",
		},
		{
			Name: "Negative length",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: -1, C: "0"},
			Want: "abc",
		},
	}

	utils.RunThreeArgumentTestCases(t, "PadRight", strutil.PadRight, testCases)
}

func TestPad(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[string, int, string], string]{
		{
			Name: "Pad single character on both sides",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: "0"},
			Want: "0abc0",
		},
		{
			Name: "Pad multiple characters on both sides",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 6, C: "xy"},
			Want: "xyabcx",
		},
		{
			Name: "Original string longer than padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 2, C: "0"},
			Want: "abc",
		},
		{
			Name: "Empty original string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 2, C: "0"},
			Want: "00",
		},
		{
			Name: "Empty padding string",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: 5, C: ""},
			Want: "abc",
		},
		{
			Name: "Empty input and padding",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "", B: 0, C: ""},
			Want: "",
		},
		{
			Name: "Negative length",
			Args: utils.ThreeArgumentTestCasesArgsType[string, int, string]{A: "abc", B: -1, C: "0"},
			Want: "abc",
		},
	}

	utils.RunThreeArgumentTestCases(t, "Pad", strutil.Pad, testCases)
}
