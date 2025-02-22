package main

import (
	"errors"
	"fmt"
	"strings"
)

type ErrInvalidId int

type Employee struct {
	Id     int
	Name   string
	Salary int
}

const (
	InvalidID ErrInvalidId = iota
)

func (e ErrInvalidId) Error() string {
	return fmt.Sprintf("Invalid id: %d", e)
}

// ErrEmptyField represents an error for an empty Employee field
type ErrEmptyField struct {
	FieldName string
}

// Error implements the error interface for ErrEmptyField
func (e *ErrEmptyField) Error() string {
	return fmt.Sprintf("empty field: %s", e.FieldName)
}

// ValidationErrors holds multiple validation errors
type ValidationErrors struct {
	Errors []error
}

// Error implements the error interface for ValidationErrors
func (v *ValidationErrors) Error() string {
	if len(v.Errors) == 0 {
		return "no validation errors"
	}
	var errMsgs []string
	for _, err := range v.Errors {
		errMsgs = append(errMsgs, err.Error())
	}
	return fmt.Sprintf("validation failed: %s", strings.Join(errMsgs, ", "))
}

// validateId validates the id
func validateId(id int) error {
	if id > 1_000 {
		return ErrInvalidId(id)
	}
	return nil
}

// validateEmployee checks for all empty fields in an Employee
func validateEmployee(emp Employee) error {
	var validationErrors ValidationErrors

	if emp.Name == "" {
		validationErrors.Errors = append(validationErrors.Errors, &ErrEmptyField{FieldName: "Name"})
	}
	if emp.Id <= 0 {
		validationErrors.Errors = append(validationErrors.Errors, &ErrEmptyField{FieldName: "Id"})
	}
	if emp.Salary <= 0 {
		validationErrors.Errors = append(validationErrors.Errors, &ErrEmptyField{FieldName: "Salary"})
	}
	if idErr := validateId(emp.Id); idErr != nil {
		validationErrors.Errors = append(validationErrors.Errors, idErr)
	}

	if len(validationErrors.Errors) > 0 {
		return &validationErrors
	}
	return nil
}

func main() {
	emp := Employee{
		Id:     0,  // Invalid ID
		Name:   "", // Empty name
		Salary: 0,  // Invalid salary
	}

	err := validateEmployee(emp)
	if err != nil {
		var validationErrors *ValidationErrors
		if errors.As(err, &validationErrors) {
			fmt.Println("Found multiple validation errors:")
			for _, err := range validationErrors.Errors {
				var emptyFieldErr *ErrEmptyField
				if errors.As(err, &emptyFieldErr) {
					fmt.Printf("- Field %s is empty\n", emptyFieldErr.FieldName)
				} else {
					fmt.Printf("- %v\n", err)
				}
			}
		}
		return
	}

	var id int
	fmt.Println("Enter an id: ")
	_, err = fmt.Scanln(&id)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = validateId(id)
	if errors.Is(err, ErrInvalidId(id)) {
		fmt.Println(err)
		return
	}
}
