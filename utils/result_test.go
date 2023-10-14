package utils_test

import (
	"errors"
	"testing"

	"github.com/GrewalAS/godash/utils"
)

var (
	errRand = errors.New("error")
	errOops = errors.New("oops")
)

func TestNewResultForError(t *testing.T) {
	t.Parallel()

	got := utils.NewResult[any](nil, errOops)
	if got.Error == nil {
		t.Errorf("Result.Error = %v, want %v", got.Error, errOops)
	}
}

func TestResult_IsOk(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, error], any]{
		{
			Name: "result with error with int as value",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: 1,
				B: errRand,
			},
		},
		{
			Name: "result with error with struct as value",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: struct{ Name string }{Name: "John"},
				B: errRand,
			},
		},
		{
			Name: "result with error with nil as value and non nil error",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: nil,
				B: errRand,
			},
		},
		{
			Name: "result with error with nil as value and nil error",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: nil,
				B: nil,
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Result.IsOk()", func(a any, e error) any {
		r := utils.NewResult[any](a, e)
		return r.IsOk()
	}, testCases)
}

func TestResult_IsErr(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, error], any]{
		{
			Name: "result with error with int as value",
			Want: true,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: 1,
				B: errRand,
			},
		},
		{
			Name: "result with error with struct as value and nil error",
			Want: false,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: struct{ Name string }{Name: "John"},
				B: nil,
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Result.IsErr()", func(a any, e error) any {
		r := utils.NewResult[any](a, e)
		return r.IsErr()
	}, testCases)

	testCasesIgnored := []utils.GenericTestCase[utils.ThreeArgumentTestCasesArgsType[any, error, []error], any]{
		{
			Name: "result with error that ignores an error",
			Want: false,
			Args: utils.ThreeArgumentTestCasesArgsType[any, error, []error]{
				A: 1,
				B: errRand,
				C: []error{errRand},
			},
		},
		{
			Name: "result with error that does not ignore an error",
			Want: false,
			Args: utils.ThreeArgumentTestCasesArgsType[any, error, []error]{
				A: struct{ Name string }{Name: "John"},
				B: errRand,
				C: []error{errRand},
			},
		},
	}

	utils.RunThreeArgumentTestCases(t, "Result.IsErr()", func(a any, e error, i []error) any {
		r := utils.NewResultWithIgnoredErrs[any](a, e, i...)
		return r.IsErr()
	}, testCasesIgnored)
}

func TestResult_Unwrap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, error], any]{
		{
			Name: "result with error with int as value",
			Want: 1,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: 1,
				B: nil,
			},
		},
		{
			Name: "result with error with struct as value and nil error",
			Want: struct{ Name string }{Name: "John"},
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: struct{ Name string }{Name: "John"},
				B: nil,
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Result.Unwrap()", func(a any, e error) any {
		r := utils.NewResult[any](a, e)
		return r.Unwrap()
	}, testCases)
}

func TestResult_UnwrapPanics(t *testing.T) {
	t.Parallel()

	got := utils.NewResult[any](nil, errOops)
	panics := utils.Panics(func() {
		got.Unwrap()
	})

	if !panics {
		t.Errorf("Result.Unwrap() should panic")
	}
}

func TestResult_UnwrapOr(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, error], any]{
		{
			Name: "result with error with int as value",
			Want: 1,
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: 1,
				B: nil,
			},
		},
		{
			Name: "result with error with struct as value and nil error",
			Want: struct{ Name string }{Name: "John"},
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: struct{ Name string }{Name: "John"},
				B: nil,
			},
		},
		{
			Name: "result with error with nil as value and non nil error",
			Want: "default",
			Args: utils.TwoArgumentTestCasesArgsType[any, error]{
				A: nil,
				B: errRand,
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Result.UnwrapOr()", func(a any, e error) any {
		r := utils.NewResult[any](a, e)
		return r.UnwrapOr("default")
	}, testCases)
}

func TestResult_AddInfo(t *testing.T) {
	t.Parallel()

	testResultWithNoError := utils.NewResult[int](1, nil)
	testResultWithNoError.FailTestIfErr(t)
	isErr := testResultWithNoError.AddInfoIfErr("test", "other")
	testResultWithNoError.FailTestIfErr(t)
	if isErr {
		t.Errorf("Result.AddInfo should return false if there is no error")
	}

	testResult := utils.NewResult[any](nil, errOops)
	testResult.FailTestIfOk(t)
	isErr = testResult.AddInfoIfErr("TestResult_AddInfo", "second")
	if !isErr {
		t.Errorf("Result.AddInfo should return true if there is an error")
	}
	testResult.FailTestIfOk(t)

	want := "TestResult_AddInfo: second: oops"
	if testResult.Error.Error() != want {
		t.Errorf("Result should have correct error message, wanted: %s, got: %s",
			want,
			testResult.Error.Error(),
		)
	}
}
