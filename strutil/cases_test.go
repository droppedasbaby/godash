package strutil_test

import (
	"github.com/GrewalAS/godash/strutil"
	"github.com/GrewalAS/godash/utils"
	"testing"
)

func TestCamelCase(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[string],
		string]{
		{
			Name: "Empty string",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: ""},
			Want: "",
		},
		{
			Name: "Single word",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "word"},
			Want: "word",
		},
		{
			Name: "Multiple words",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "hello world"},
			Want: "helloWorld",
		},
		{
			Name: "From snake case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "hello_world"},
			Want: "helloWorld",
		},
		{
			Name: "From kebab case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "hello-world"},
			Want: "helloWorld",
		},
		{
			Name: "With spaces",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "   "},
			Want: "",
		},
		{
			Name: "With multiple underscores",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "___"},
			Want: "",
		},
		{
			Name: "With multiple dashes",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "---"},
			Want: "",
		},
	}

	utils.RunSingleArgumentTestCases(t, "CamelCase", strutil.CamelCase, testCases)
}

func TestSnakeCase(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[string],
		string]{
		{
			Name: "Empty string",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: ""},
			Want: "",
		},
		{
			Name: "Single word",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "word"},
			Want: "word",
		},
		{
			Name: "Multiple words",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "Hello World"},
			Want: "hello_world",
		},
		{
			Name: "From camel case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "helloWorld"},
			Want: "hello_world",
		},
		{
			Name: "From kebab case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "hello-world"},
			Want: "hello_world",
		},
		{
			Name: "With spaces",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "   "},
			Want: "",
		},
		{
			Name: "With multiple underscores",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "___"},
			Want: "",
		},
		{
			Name: "With multiple dashes",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "---"},
			Want: "",
		},
	}

	utils.RunSingleArgumentTestCases(t, "SnakeCase", strutil.SnakeCase, testCases)
}

func TestKebabCase(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[
		utils.SingleArgumentTestCasesArgsType[string],
		string]{
		{
			Name: "Empty string",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: ""},
			Want: "",
		},
		{
			Name: "Single word",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "word"},
			Want: "word",
		},
		{
			Name: "Multiple words",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "Hello World"},
			Want: "hello-world",
		},
		{
			Name: "From camel case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "helloWorld"},
			Want: "hello-world",
		},
		{
			Name: "From snake case",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "hello_world"},
			Want: "hello-world",
		},
		{
			Name: "With spaces",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "   "},
			Want: "",
		},
		{
			Name: "With multiple underscores",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "___"},
			Want: "",
		},
		{
			Name: "With multiple dashes",
			Args: utils.SingleArgumentTestCasesArgsType[string]{A: "---"},
			Want: "",
		},
	}

	utils.RunSingleArgumentTestCases(t, "KebabCase", strutil.KebabCase, testCases)
}
