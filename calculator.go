package calculator

import (
	"errors"
	maths "math"
)

// Takes a + b and returns the sum
func Add(a, b float64) float64 {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative numbers is not allowed")
	}
	return maths.Sqrt(a), nil
}

func Subtract(a, b float64) float64 {
	return a - b
}
