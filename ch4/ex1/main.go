package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Define a slice of int
	numbers := make([]int, 100)

	// Fill the slice with random numbers
	for i := 0; i < 100; i++ {
		numbers[i] = rand.Intn(100)
	}

	// Loop through the slice and print each number
	for i, number := range numbers {

		if i%6 == 0 {
			fmt.Println("Six!")
			continue
		}
		if i%2 == 0 {
			fmt.Println("Two!")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Three!")
			continue
		}
		fmt.Printf("Never mind %d\n", number)
	}
}
