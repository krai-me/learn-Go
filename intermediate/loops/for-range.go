//Range-based for loop
//This type of for loop is specifically used to iterate over a collection (like slices, arrays, maps, or strings).

package main

import "fmt"

func main() {
	str := "Hello"

	// Iterate over the string with both index and value
	for index, value := range str {
		// Use %c to print the character representation of a rune
		fmt.Printf("Index: %d, Value: %c\n", index, value)
	}

	// Iterate over the string but ignore the index
	// Using _ to ignore the index prevents syntax errors
	for _, value := range str {
		// Only use value (character representation of a rune)
		fmt.Printf("Value: %c\n", value)
	}

	// Iterate over a string but ignore the value
	// Using _ to ignore the value prevents syntax errors
	for index, _ := range str {
		// Only use index
		fmt.Printf("Index: %d\n", index)
	}
}
