// Package slices offers various utilities for handling and manipulating
// slice data structures in Go.
package slices_test

import (
	"testing"

	"github.com/GrewalAS/godash/slices"
)

func testForeachSumInt(t *testing.T, input []int, want int) {
	t.Helper()
	sum := 0
	slices.Foreach(input, func(i int) { sum += i })
	if sum != want {
		t.Errorf("sum = %d, want %d", sum, want)
	}
}

func testForeachMultiplyInt(t *testing.T, input []int, want int) {
	t.Helper()
	product := 1
	slices.Foreach(input, func(v int) { product *= v })
	if product != want {
		t.Errorf("product = %d, want %d", product, want)
	}
}

func TestForeach_WithInt(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name   string
		input  []int
		want   int
		rnFunc func(*testing.T, []int, int)
	}{
		{"sum of elements", []int{1, 2, 3, 4, 5}, 15, testForeachSumInt},
		{"empty slice sum", []int{}, 0, testForeachSumInt},
		{"single element sum", []int{5}, 5, testForeachSumInt},
		{"negative elements sum", []int{-1, -2, -3, -4, -5}, -15, testForeachSumInt},
		{"mixed elements sum", []int{-1, 2, -3, 4, -5}, -3, testForeachSumInt},
		{"product of elements", []int{1, 2, 3, 4, 5}, 120, testForeachMultiplyInt},
		{"empty slice product", []int{}, 1, testForeachMultiplyInt},
		{"single element product", []int{5}, 5, testForeachMultiplyInt},
		{"negative elements product", []int{-1, -2, -3, -4, -5}, -120, testForeachMultiplyInt},
		{"mixed elements product", []int{-1, 2, -3, 4, -5}, -120, testForeachMultiplyInt},
	}

	for _, tc := range testCases {
		tcCopy := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tcCopy.rnFunc(t, tcCopy.input, tcCopy.want)
		})
	}
}

func testForeachSumFloat64(t *testing.T, input []float64, want float64) {
	t.Helper()
	sum := 0.0
	slices.Foreach(input, func(f float64) { sum += f })
	if sum != want {
		t.Errorf("sum = %f, want %f", sum, want)
	}
}

func TestForeach_WithFloat64(t *testing.T) {
	t.Parallel()
	t.Run("sum of elements", func(t *testing.T) {
		t.Parallel()
		input := []float64{1.1, 2.2, 3.3}
		testForeachSumFloat64(t, input, 6.6)
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		input := []float64{}
		testForeachSumFloat64(t, input, 0.0)
	})
}

func testForeachConcatString(t *testing.T, input []string, want string) {
	t.Helper()
	result := ""
	slices.Foreach(input, func(s string) { result += s })
	if result != want {
		t.Errorf("result = %s, want %s", result, want)
	}
}

func TestForeach_WithString(t *testing.T) {
	t.Parallel()
	t.Run("concatenate elements", func(t *testing.T) {
		t.Parallel()
		input := []string{"a", "b", "c"}
		testForeachConcatString(t, input, "abc")
	})

	t.Run("empty slice", func(t *testing.T) {
		t.Parallel()
		input := []string{}
		testForeachConcatString(t, input, "")
	})
}
