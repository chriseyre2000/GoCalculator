package calculator

import "errors"

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

func Subtract(a, b float64) float64 {
	return a - b
}
