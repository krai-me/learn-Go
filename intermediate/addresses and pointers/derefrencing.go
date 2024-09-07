// By dereferencing pointerForStr, we directly modify the original value of lyrics,
// demonstrating how pointers provide a way to indirectly change values stored in memory.
package main

import "fmt"

func main() {
	// Initialize a string variable and a pointer to the string
	lyrics := "Moments so dear"
	pointerForStr := &lyrics

	// Dereference the pointer to change the value at the address
	*pointerForStr = "Journeys to plan"

	// Print the updated value of the original variable
	fmt.Println(lyrics) // Prints: Journeys to plan
}
