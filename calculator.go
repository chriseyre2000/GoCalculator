package calculator

// Takes a + b and returns the sum
func Add(a, b float64) float64 {
	return a + b
}

func Divide(a, b float64) (float64, error) {
	return a / b, nil
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Subtract(a, b float64) float64 {
	return a - b
}
