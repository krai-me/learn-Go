package main

import "fmt"

// Define the Employee struct
type Employee struct {
	firstName string
	lastName  string
	age       int
	title     string
}

func main() {
	// Create an instance of Employee
	steve := Employee{firstName: "xyz", lastName: "xyz", age: 34, title: "xyz"}

	// Create a pointer to the instance of Employee
	pointerToSteve := &steve

	// Modify the fields using the pointer

	// Option 1: Using the dereferencing operator (*pointerToSteve)
	(*pointerToSteve).firstName = "Steve"

	// Option 2: Simpler method (recommended in Go)
	pointerToSteve.lastName = "Jobs"

	// Print the updated employee details
	fmt.Println("Updated Employee:", steve) // Output: Updated Employee: {Steven Jobs 34 Junior Manager}

	// Using pointers in a function to modify the struct
	rect := Rectangle{length: 10, width: 5}
	fmt.Println("Before modifying rectangle length:", rect.length) // Output: 10

	rect.modify(20)
	fmt.Println("After modifying rectangle length:", rect.length) // Output: 20
}

// Define the Rectangle struct
type Rectangle struct {
	length float32
	width  float32
}

// Function to modify the length of a rectangle (uses a pointer)
func (rectangle *Rectangle) modify(newLength float32) {
	rectangle.length = newLength
}
