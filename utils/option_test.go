package utils_test

import (
	"testing"

	"github.com/GrewalAS/godash/utils"
)

func TestNewOption(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[any], any]{
		{
			Name: "new option of int",
			Want: 1,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: 1,
			},
		},
		{
			Name: "new option of string",
			Want: "hello",
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: "hello",
			},
		},
		{
			Name: "new option of struct",
			Want: struct{ Name string }{Name: "John"},
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: struct{ Name string }{Name: "John"},
			},
		},
		{
			Name: "new option of nil",
			Want: nil,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: nil,
			},
		},
	}

	utils.RunSingleArgumentTestCases[any](t, "NewOption()", func(a any) any {
		o := utils.NewOption[any](a)
		return *o.Value
	}, testCases)
}

func TestOption_IsNil(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[any], bool]{
		{
			Name: "returns true because value is nil",
			Want: true,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: nil,
			},
		},
		{
			Name: "returns false because value is not nil, is struct",
			Want: false,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: struct{ Name string }{Name: "John"},
			},
		},
		{
			Name: "returns false because value is not nil, is string",
			Want: false,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: "hello",
			},
		},
		{
			Name: "returns false because value is not nil, is int",
			Want: false,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: 1,
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Option.IsNil()", func(a any) bool {
		o := utils.NewOption[any](a)
		return o.IsNil()
	}, testCases)
}

func TestOption_IsSome(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[any], bool]{
		{
			Name: "return true because value is not nil, is int",
			Want: true,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: 1,
			},
		},
		{
			Name: "return true because value is not nil, is string",
			Want: true,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: "hello",
			},
		},
		{
			Name: "return true because value is not nil, is struct",
			Want: true,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: struct{ Name string }{Name: "John"},
			},
		},
		{
			Name: "return false because value is nil",
			Want: false,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: nil,
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Option.IsSome()", func(A any) bool {
		o := utils.NewOption[any](A)
		return o.IsSome()
	}, testCases)
}

func TestOption_Unwrap(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.SingleArgumentTestCasesArgsType[any], any]{
		{
			Name: "unwrap of int",
			Want: 1,
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: 1,
			},
		},
		{
			Name: "unwrap of string",
			Want: "hello",
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: "hello",
			},
		},
		{
			Name: "unwrap of struct",
			Want: struct{ Name string }{Name: "John"},
			Args: utils.SingleArgumentTestCasesArgsType[any]{
				A: struct{ Name string }{Name: "John"},
			},
		},
	}

	utils.RunSingleArgumentTestCases(t, "Option.Unwrap()", func(a any) any {
		o := utils.NewOption[any](a)
		return o.Unwrap()
	}, testCases)
}

func TestOption_UnwrapPanic(t *testing.T) {
	t.Parallel()

	option := utils.NewOption[any](nil)
	panics := utils.Panics(func() {
		option.Unwrap()
	})

	if !panics {
		t.Errorf("Option.Unwrap() should panic when Value is nil")
	}
}

func TestOption_UnwrapOr(t *testing.T) {
	t.Parallel()

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, any], any]{
		{
			Name: "unwrap or of int",
			Want: 1,
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: 1,
				B: 10,
			},
		},
		{
			Name: "unwrap or of string",
			Want: "hello",
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: "hello",
				B: "world",
			},
		},
		{
			Name: "unwrap or of struct",
			Want: struct{ Name string }{Name: "John"},
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: struct{ Name string }{Name: "John"},
				B: struct{ Name string }{Name: "Doe"},
			},
		},
		{
			Name: "unwrap or of nil",
			Want: "default",
			Args: utils.TwoArgumentTestCasesArgsType[any, any]{
				A: nil,
				B: "default",
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Option.UnwrapOr()", func(a any, b any) any {
		o := utils.NewOption[any](a)
		return o.UnwrapOr(b)
	}, testCases)
}

func TestOption_UnwrapOrElse(t *testing.T) {
	t.Parallel()

	defaultValue := "default"

	testCases := []utils.GenericTestCase[utils.TwoArgumentTestCasesArgsType[any, func() any], any]{
		{
			Name: "unwrap or else if nil",
			Want: defaultValue,
			Args: utils.TwoArgumentTestCasesArgsType[any, func() any]{
				A: nil,
				B: func() any {
					return defaultValue
				},
			},
		},
		{
			Name: "unwrap or else if not nil",
			Want: "hello",
			Args: utils.TwoArgumentTestCasesArgsType[any, func() any]{
				A: "hello",
				B: func() any {
					return defaultValue
				},
			},
		},
	}

	utils.RunTwoArgumentTestCases(t, "Option.UnwrapOrElse()", func(a any, b func() any) any {
		o := utils.NewOption[any](a)
		return o.UnwrapOrElse(b)
	}, testCases)
}
