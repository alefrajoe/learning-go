package main

import (
	"fmt"
)

func UpdateSlice(slice []string, word string) {
	// Update the last element of the slice
	slice[len(slice)-1] = word
	// Print the slice from the function
	fmt.Println(slice)
}

func GrowSlice(slice []string, word string) {
	// Grow the slice by 1
	slice = append(slice, word)
	// Print the slice from the function
	fmt.Println(slice)
}

func main() {
	slice := []string{"Hello", "World"}
	UpdateSlice(slice, "Golang")
	GrowSlice(slice, "Golang")
	fmt.Println(slice)
}
