// Arrays: Have a fixed size determined at compile time.
// Once you declare an array with a specific size, it cannot grow or shrink.
package main

import "fmt"

func main() {
	// Declare an array named 'playerScores' with a fixed size of 4 elements, all of which are of type int.
	var playerScores [4]int

	// At this point, the array has not been initialized with any values, so all elements are set to the default zero value for ints (0).
	fmt.Println(playerScores) // Output: [0 0 0 0]

	// Declare an array named 'fruits' with a fixed size of 3 elements, and initialize it with specific string values.
	fruits := [3]string{"apple", "mango", "banana"}

	// This array is explicitly initialized with 3 string values.
	fmt.Println(fruits) // Output: [apple mango banana]

	// Declare an array named 'autoarray' with an unspecified size, and let the compiler determine the size based on the number of elements provided.
	autoarray := [...]string{"it", "can", "be", "of", "any", "length"}

	// The compiler determines the size of this array based on the number of elements in the initialization list.
	fmt.Println(autoarray) // Output: [it can be of any length]
}
