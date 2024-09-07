// The underscore (_) is used as a placeholder to ignore values that you don't need.
// This allows you to avoid syntax errors while omitting unnecessary variables.
package main

import "fmt"

func main() {
	// Example of using the underscore (_) to ignore the index in a range-based for loop

	str := "Hello"
	for _, value := range str {
		// Print each character in the string
		fmt.Printf("Character: %c\n", value)
	}

	// Example of ignoring an unused return value from a function
	result, _ := someFunction()
	fmt.Println("Result:", result)
}

// someFunction returns two values, but we only need the first one
func someFunction() (int, string) {
	return 42, "example"
}
