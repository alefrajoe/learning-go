package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{firstName: firstName, lastName: lastName, age: age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName: firstName, lastName: lastName, age: age}
}

func main() {
	person := MakePerson("John", "Doe", 30)
	personPointer := MakePersonPointer("John", "Doe", 30)

	fmt.Println(person)
	fmt.Println(personPointer)
}
