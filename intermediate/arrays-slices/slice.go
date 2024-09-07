//Slices: Are more flexible. They can grow or shrink dynamically, making them ideal for situations where the
//number of elements is not known in advance or changes over time.

package main

import "fmt"

func main() {
	// Creating an empty slice of integers
	var numberSlice []int
	fmt.Println("Empty numberSlice:", numberSlice) // Output: Empty numberSlice: []

	// Creating an empty slice of strings
	stringSlice := []string{}
	fmt.Println("Empty stringSlice:", stringSlice) // Output: Empty stringSlice: []

	// Creating a slice with initial values
	names := []string{"ash", "brock", "misty", "oak"}
	fmt.Println("Names slice:", names) // Output: Names slice: [ash brock misty]

	// Creating an array and then creating slices from that array
	array := [5]int{2, 5, 7, 1, 3}

	// Slice of the whole array
	sliceVersion := array[:]
	fmt.Println("Slice of the entire array:", sliceVersion) // Output: Slice of the entire array: [2 5 7 1 3]

	// Slice of a subset of the array
	partialSlice := array[2:5]
	fmt.Println("Partial slice (indices 2 to 4):", partialSlice) // Output: Partial slice (indices 2 to 4): [7 1 3]

	// Access and modify elements in a slice
	fmt.Println("Element at index 1 in names:", names[1]) // Output: Element at index 1 in names: brock
	names[3] = "serena"
	fmt.Println("Modified element at index 3 in names:", names[3]) // Output: Modified element at index 3 in names: serena

	// Append a new element to the slice
	names = append(names, "james")
	fmt.Println("Names slice after appending an element:", names) // Output: Names slice after appending an element: [ash brock misty serena james]
}
