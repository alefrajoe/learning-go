package main

import "fmt"

func main() {

	// Define greetings variables
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	// Define the 3 subslices
	slice1 := greetings[:2]
	slice2 := greetings[1:4]
	slice3 := greetings[3:]

	// Print the greetings
	fmt.Println(greetings)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
