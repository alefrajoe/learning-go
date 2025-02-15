package main

import "fmt"

func main() {

	// Initialise the total variable
	total := 0

	// Loop from 0 to 10
	for i := 0; i <= 10; i++ {
		total += i
	}

	// Print the total
	fmt.Println(total)
}
