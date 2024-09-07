package main

import "fmt"

func main() {
	// Declare an array named 'numbers' with an implicit size, based on the number of elements provided.
	numbers := [...]int{1, 23, 32, 42}

	// Access and print the value at index 0 of the 'numbers' array.
	// Array indexing starts at 0, so numbers[0] refers to the first element, which is 1.
	fmt.Println(numbers[0]) // Output: 1

	// Store the value at index 2 of the 'numbers' array in a variable named 'isthirtytwo'.
	// Here, numbers[2] refers to the third element in the array, which is 32.
	isthirtytwo := numbers[2]

	fmt.Println(isthirtytwo) // Output: 32

	// Modify the value at index 2 of the 'numbers' array.
	// This sets the value at index 2 to 113.
	numbers[2] = 113

	// The output will be 113 after modification.
	fmt.Println(numbers[2]) // Output: 113

	// Attempting to use a value of a different type (e.g., a string) would result in a compile-time error.
	// numbers[2] = "stringValue" // This line would raise a compile-time error
}
