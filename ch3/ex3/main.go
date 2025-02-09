package main

import "fmt"

type Employee struct {
	id        int
	firstName string
	lastName  string
}

func main() {

	// 1. Create a new employee
	employee1 := Employee{
		1,
		"John",
		"Doe",
	}

	// 2. Define the employee2
	employee2 := Employee{
		id:        2,
		firstName: "Jane",
		lastName:  "Smith",
	}

	// 3. Define the employee from a var
	var employee3 Employee
	employee3.id = 3
	employee3.firstName = "Alice"
	employee3.lastName = "Johnson"

	// 4. Print the employees
	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
