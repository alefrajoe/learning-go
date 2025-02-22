// Package calculator provides basic arithmetic operations.
// It implements simple mathematical functions for calculations.
package calculator

import "golang.org/x/exp/constraints"

// Number is an interface that combines Integer and Float constraints
type Number interface {
	constraints.Integer | constraints.Float
}

// Add returns the sum of two numbers a and b.
// For more information about addition, see:
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}

// Subtract returns the difference between two integers a and b.
// For more information about subtraction, see:
// https://www.mathsisfun.com/numbers/subtraction.html
func Subtract(a, b int) int {
	return a - b
}
