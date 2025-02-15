package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	// Crate a huge number of people
	people := make([]Person, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		people[i] = Person{
			firstName: "John",
			lastName:  "Doe",
			age:       i,
		}
	}

	// Print the first 10 people
	for _, person := range people {
		fmt.Println(person)
	}
}
