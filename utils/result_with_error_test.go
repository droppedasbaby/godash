package utils_test

import (
	"errors"
	"testing"

	"godash/utils"
)

var (
	errRand = errors.New("error")
	errOops = errors.New("oops")
)

func TestNewResultWithErrorForError(t *testing.T) {
	t.Parallel()

	got := utils.NewResultWithError[any](nil, errOops)
	if got.Error == nil {
		t.Errorf("ResultWithError.Error = %v, want %v", got.Error, errOops)
	}
}

func TestResultWithError_IsOk(t *testing.T) {
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

	utils.RunTwoArgumentTestCases(t, "ResultWithError.IsOk()", func(a any, e error) any {
		r := utils.NewResultWithError[any](a, e)
		return r.IsOk()
	}, testCases)
}

func TestResultWithError_IsErr(t *testing.T) {
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

	utils.RunTwoArgumentTestCases(t, "ResultWithError.IsErr()", func(a any, e error) any {
		r := utils.NewResultWithError[any](a, e)
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

	utils.RunThreeArgumentTestCases(t, "ResultWithError.IsErr()", func(a any, e error, i []error) any {
		r := utils.NewResultWithErrorWithIgnoredErrs[any](a, e, i...)
		return r.IsErr()
	}, testCasesIgnored)
}

func TestResultWithError_Unwrap(t *testing.T) {
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

	utils.RunTwoArgumentTestCases(t, "ResultWithError.Unwrap()", func(a any, e error) any {
		r := utils.NewResultWithError[any](a, e)
		return r.Unwrap()
	}, testCases)
}

func TestResultWithError_UnwrapPanics(t *testing.T) {
	t.Parallel()

	got := utils.NewResultWithError[any](nil, errOops)
	panics := utils.Panics(func() {
		got.Unwrap()
	})

	if !panics {
		t.Errorf("ResultWithError.Unwrap() should panic")
	}
}

func TestResultWithError_UnwrapOr(t *testing.T) {
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

	utils.RunTwoArgumentTestCases(t, "ResultWithError.UnwrapOr()", func(a any, e error) any {
		r := utils.NewResultWithError[any](a, e)
		return r.UnwrapOr("default")
	}, testCases)
}

func TestResultWithError_AddInfo(t *testing.T) {
	t.Parallel()

	testResultWithNoError := utils.NewResultWithError[int](1, nil)
	testResultWithNoError.FailTestIfErr(t)
	isErr := testResultWithNoError.AddInfoIfErr("test", "other")
	testResultWithNoError.FailTestIfErr(t)
	if isErr {
		t.Errorf("ResultWithError.AddInfo should return false if there is no error")
	}

	testResultWithError := utils.NewResultWithError[any](nil, errOops)
	testResultWithError.FailTestIfOk(t)
	isErr = testResultWithError.AddInfoIfErr("TestResultWithError_AddInfo", "second")
	if !isErr {
		t.Errorf("ResultWithError.AddInfo should return true if there is an error")
	}
	testResultWithError.FailTestIfOk(t)

	want := "TestResultWithError_AddInfo: second: oops"
	if testResultWithError.Error.Error() != want {
		t.Errorf("ResultWithError should have correct error message, wanted: %s, got: %s",
			want,
			testResultWithError.Error.Error(),
		)
	}
}
