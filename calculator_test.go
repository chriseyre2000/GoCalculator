package calculator_test

import (
	"calculator"
	maths "math"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 1},
		{a: 7, b: 2, want: 3.5},
		{a: 1, b: -1, want: -1},
		{a: 1, b: 3, want: 0.333333},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input got, %v", tc.a, tc.b, err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestDivideInvalidInput(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
	}

	testCases := []testCase{
		{a: 2, b: 0},
	}

	for _, tc := range testCases {
		_, err := calculator.Divide(tc.a, tc.b)
		if err == nil {
			t.Fatalf("Divide(%f, %f): want error for invalid input got nil", tc.a, tc.b)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 3, want: 6},
		{a: 6, b: 0.5, want: 3},
		{a: -5, b: -5, want: 25},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: -1, b: 1, want: -2},
		{a: -5, b: -5, want: 0},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a    float64
		want float64
	}

	testCases := []testCase{
		{a: 9, want: 3},
		{a: 0.01, want: 0.1},
	}

	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("Sqrt(%f): want no error for valid input got, %v", tc.a, err)
		}
		if !closeEnough(tc.want, got, 0.1) {
			t.Errorf("want %f, got %f", tc.want, got)
		}
	}
}

func TestSqrtInvalidInput(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a float64
	}

	testCases := []testCase{
		{a: -1,},
	}

	for _, tc := range testCases {
		_, err := calculator.Sqrt(tc.a)
		if err == nil {
			t.Fatalf("Sqrt(%f): want error for invalid input got nil", tc.a)
		}
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return maths.Abs(a-b) <= tolerance
}
