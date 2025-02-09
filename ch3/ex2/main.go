package main

import "fmt"

func main() {

	// Define the message
	message := "Hi Alice and Bob"

	// Print the message
	for i, char := range message {
		if i == 3 {
			fmt.Println(string(char))
			break
		}
	}
}
