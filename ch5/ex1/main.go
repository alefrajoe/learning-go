package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func plus(a, b int) (int, error) {
	return a + b, nil
}

func minus(a, b int) (int, error) {
	return a - b, nil
}

func multiply(a, b int) (int, error) {
	return a * b, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {

	// Read the user input from cli
	argv := os.Args[1:]
	if len(argv) != 3 {
		fmt.Println("Usage: go run main.go <operation> <number1> <number2>")
		return
	}

	// Strip blank spaces from all arguments
	for i, arg := range argv {
		argv[i] = strings.TrimSpace(arg)
	}

	// Define the operations
	operations := map[string]func(int, int) (int, error){
		"+": plus,
		"-": minus,
		"*": multiply,
		"/": divide,
	}

	// Check whether argv[1] is a valid operation
	if _, ok := operations[argv[1]]; !ok {
		fmt.Println("Invalid operation ", argv[1])
		return
	}

	// Convert argv[0] and argv[2] to integers
	number1, err := strconv.Atoi(argv[0])
	if err != nil {
		fmt.Println("Invalid number ", argv[0])
		return
	}
	number2, err := strconv.Atoi(argv[2])
	if err != nil {
		fmt.Println("Invalid number ", argv[2])
		return
	}

	// Get the operation function
	operation := operations[argv[1]]

	// Perform the operation
	result, err := operation(number1, number2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the result
	fmt.Println(result)
}
