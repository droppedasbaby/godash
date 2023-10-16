package strutil_test

import (
	"github.com/GrewalAS/godash/strutil"
	"github.com/GrewalAS/godash/utils"
	"testing"
)

func TestTruncate(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[string, int], string]{
		{
			Name: "Empty string, no truncation",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "", B: 0},
			Want: "",
		},
		{
			Name: "Empty string, truncation to 3",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "", B: 3},
			Want: "",
		},
		{
			Name: "String shorter than truncation length",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "abc", B: 5},
			Want: "abc",
		},
		{
			Name: "String equal to truncation length",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "abc", B: 3},
			Want: "abc",
		},
		{
			Name: "String longer than truncation length",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "abcdef", B: 3},
			Want: "abc",
		},
		{
			Name: "Truncate to zero",
			Args: utils.TwoArgumentTestCasesArgsType[string, int]{A: "abcdef", B: 0},
			Want: "",
		},
	}

	utils.RunTwoArgumentTestCases(t, "Truncate", strutil.Truncate, testCases)
}
