package main

import "fmt"

// Define a function that multiplies two integers and returns the result.
func multiplier(x int32, y int32) int32 {
	// Return the product of x and y.
	return x * y
}

func main() {
	// Example 1: Calling the function with literal values.
	var product int32
	product = multiplier(25, 4)
	fmt.Println(product) // Prints: 100

	// Example 2: Calling the function with variables.
	var mainX, mainY, newProduct int32
	mainX = 6
	mainY = 7
	newProduct = multiplier(mainX, mainY)
	fmt.Println(newProduct) // Prints: 42

	/*
	   It is crucial to provide the correct number and type of arguments when calling a function.
	   If the number of arguments does not match the number of parameters, or if the types of arguments do not match the parameter types,
	   Go will produce a compilation error.

	   For example:
	   - If you call multiplier with incorrect arguments, such as:
	   - multiplier(10, "text") // Error: cannot use "text" (value of type string) as int32 value in argument
	   - This error occurs because the second argument "text" is of type string, while the function expects an int32.

	   Additionally, if you provide too few arguments:
	   - multiplier(5) // Error: not enough arguments in call to multiplier
	   - The function expects two arguments, but only one is provided.

	   **Parameter Type Definition**:
	   - Function parameters must be defined with types. If you omit the type when defining a parameter, Go will throw an error.
	   - For example:
	   - func invalidMultiplier(x, y) int32 // Error: parameter x and y must have a type
	   - All parameters need explicit types in the function definition to avoid compilation errors.
	*/
}
