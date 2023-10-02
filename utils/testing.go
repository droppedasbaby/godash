package utils

import (
	"fmt"
	"reflect"
	"testing"
)

// GenericTestCase is a generic test case.
type GenericTestCase[A any, W any] struct {
	Name string
	Args A
	Want W
}

// GenericTestCaseWithNoWant is a generic test case with an option want.
type GenericTestCaseWithNoWant[A any] struct {
	Name string
	Args A
}

// SingleArgumentTestCasesArgsType is the type of the arguments for a single argument test case.
type SingleArgumentTestCasesArgsType[T any] struct {
	A T
}

// RunSingleArgumentTestCases runs a sets of test cases for a function that takes a single argument.
func RunSingleArgumentTestCases[T, R any](
	t *testing.T,
	testName string,
	runFn func(T) R,
	testCases []GenericTestCase[SingleArgumentTestCasesArgsType[T], R],
) {
	t.Helper()

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			got := runFn(testCase.Args.A)
			if !reflect.DeepEqual(got, testCase.Want) {
				t.Errorf("%v = %v, want %v", testName, got, testCase.Want)
			}
		})
	}
}

// TwoArgumentTestCasesArgsType is the type of the arguments for a two argument test case.
type TwoArgumentTestCasesArgsType[T, U any] struct {
	A T
	B U
}

// RunTwoArgumentTestCases runs a sets of test cases for a function that takes two arguments.
func RunTwoArgumentTestCases[T, U, R any](
	t *testing.T,
	testName string,
	runFn func(T, U) R,
	testCases []GenericTestCase[TwoArgumentTestCasesArgsType[T, U], R],
) {
	t.Helper()

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			got := runFn(testCase.Args.A, testCase.Args.B)
			if !reflect.DeepEqual(got, testCase.Want) {
				t.Errorf("%v = %v, want %v", testName, got, testCase.Want)
			}
		})
	}
}

// ThreeArgumentTestCasesArgsType is the type of the arguments for a three argument test case.
type ThreeArgumentTestCasesArgsType[T, U, V any] struct {
	A T
	B U
	C V
}

// RunThreeArgumentTestCases runs a sets of test cases for a function that takes three arguments.
func RunThreeArgumentTestCases[T, U, V, R any](
	t *testing.T,
	testName string,
	runFn func(T, U, V) R,
	testCases []GenericTestCase[ThreeArgumentTestCasesArgsType[T, U, V], R],
) {
	t.Helper()

	for _, tc := range testCases {
		testCase := tc
		t.Run(testCase.Name, func(t *testing.T) {
			t.Parallel()
			got := runFn(testCase.Args.A, testCase.Args.B, testCase.Args.C)
			if !reflect.DeepEqual(got, testCase.Want) {
				t.Errorf("%v = %v, want %v", testName, got, testCase.Want)
			}
		})
	}
}

var ErrTestCasesAndWantsLenDiff = fmt.Errorf("test cases and wants length differ")

// TestCasesArgsType is the type of the arguments for a test case.
type TestCasesArgsType[T any] interface {
	SingleArgumentTestCasesArgsType[T] | TwoArgumentTestCasesArgsType[T, T]
}

// AddWantToTestCases adds a wants to a sets of test cases with option wants.
func AddWantToTestCases[A any, W any, WS []W](
	data []GenericTestCaseWithNoWant[SingleArgumentTestCasesArgsType[A]],
	wants []WS) Result[[]GenericTestCase[SingleArgumentTestCasesArgsType[A], WS]] {
	testCases := make([]GenericTestCase[SingleArgumentTestCasesArgsType[A], WS], len(data))

	if len(wants) != len(data) {
		return *NewResult(testCases, ErrTestCasesAndWantsLenDiff)
	}

	for i, element := range data {
		want := []W{}
		copy(want, wants[i])

		testCases[i] = GenericTestCase[SingleArgumentTestCasesArgsType[A], WS]{
			Name: element.Name,
			Args: element.Args,
			Want: want,
		}
	}

	return *NewResult(testCases, nil)
}
