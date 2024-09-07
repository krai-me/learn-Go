package main

import "fmt"

// Define a function that returns an integer value.
func getLengthOfCentralPark() int32 {
	// Declare a variable with a value of 51.
	var lengthInBlocks int32
	lengthInBlocks = 51

	// Return the value of lengthInBlocks.
	return lengthInBlocks
}

func main() {
	// Declare a variable to store the return value from getLengthOfCentralPark().
	var centralParkLength int32

	// Call the function and assign its return value to centralParkLength.
	centralParkLength = getLengthOfCentralPark()

	// Print the value stored in centralParkLength.
	fmt.Println(centralParkLength) // Prints: 51

	/*
	   If we try to store the return value of getLengthOfCentralPark() into a variable of a different type,
	   like a float or a string, it will cause a compilation error because the types must match.

	   For example, the following line would cause an error:
	   var wrongType float32
	   wrongType = getLengthOfCentralPark() // Error: cannot use getLengthOfCentralPark() (value of type int32) as float32 value in assignment
	*/
}
