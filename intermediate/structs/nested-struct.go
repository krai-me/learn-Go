package main

import "fmt"

// Define the Name struct for first and last names
type Name struct {
	firstName string
	lastName  string
}

// Define the Employee struct with a named nested struct
type Employee struct {
	name  Name
	age   int
	title string
}

// Define the Employee struct with an anonymous nested struct
type EmployeeAnonymous struct {
	Name  // Anonymous struct field
	age   int
	title string
}

func main() {
	// Create an instance of Employee with a named Name field
	carl := Employee{name: Name{firstName: "Carl", lastName: "Carlson"}, age: 32, title: "Engineer"}

	// Access fields of the nested struct
	fmt.Println("Named Struct Access:")
	fmt.Println("Last Name:", carl.name.lastName)   // Output: Carlson
	fmt.Println("First Name:", carl.name.firstName) // Output: Carl

	// Create an instance of EmployeeAnonymous with an anonymous Name field
	carlAnon := EmployeeAnonymous{Name{"Carl", "Carlson"}, 32, "Engineer"}

	// Access fields directly as Name is anonymous
	fmt.Println("\nAnonymous Struct Access:")
	fmt.Println("First Name:", carlAnon.firstName) // Output: Carl
	fmt.Println("Last Name:", carlAnon.lastName)   // Output: Carlson
}
