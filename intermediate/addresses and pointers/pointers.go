// Pointers are powerful because they let you manipulate data directly through its memory address,
// which can be useful for managing memory efficiently and for creating complex data structures.
// Pointers in Go allow you to reference and manipulate the original value directly, without creating copies.
// This helps reduce memory usage and improves performance by avoiding unnecessary data duplication.
package main

import "fmt"

func main() {
	// Initialize an integer variable and a pointer to an integer
	minutes := 525600
	pointerForInt := &minutes

	// Print value, address, and dereferenced value
	fmt.Println("Value of minutes:", minutes)
	fmt.Println("Address of minutes:", pointerForInt)  // Memory address
	fmt.Println("Dereferenced value:", *pointerForInt) // Value at address

	// Modify value through the pointer
	// * (dereference operator): This accesses the value stored at the memory address a pointer is holding.
	*pointerForInt += 2
	fmt.Println("Updated value of minutes:", minutes) // Updated value

	// Pointer for another integer
	anotherMinutes := 60
	anotherPointer := &anotherMinutes

	// Print value, address, and dereferenced value
	fmt.Println("Value of anotherMinutes:", anotherMinutes)
	fmt.Println("Address of anotherMinutes:", anotherPointer) // Memory address
	fmt.Println("Dereferenced value:", *anotherPointer)       // Value at address

	// Pointer for a string
	message := "Hello, Go!"
	pointerForString := &message

	// Print value, address, and dereferenced value
	fmt.Println("Value of message:", message)
	fmt.Println("Address of message:", pointerForString)  // Memory address
	fmt.Println("Dereferenced value:", *pointerForString) // Value at address

	// Declaring a nil pointer
	var newPointer *int
	//Here, newPointer is declared but not initialized, so its value is nil.
	//This indicates that it doesn't currently point to any valid integer address.
	fmt.Println("Value of newPointer:", newPointer) // Prints nil
}
