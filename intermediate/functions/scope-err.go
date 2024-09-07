package main

import "fmt"

// performAddition is a function that performs addition and prints the result.
func performAddition() {
	// Variables x and y are declared in the local scope of performAddition.
	x := 5
	y := 7
	// Prints the sum of x and y.
	fmt.Println("The sum of", x, "and", y, "is", x+y)
}

func main() {
	/*
		There are three different scopes in this example:

		1. **Global Scope**: Contains the function definitions for `main()` and `performAddition()`. In Go, the global scope is where top-level functions and variables are defined.

		2. **Local Scope of `performAddition()`**: Inside this function, variables `x` and `y` are defined. These variables are local to `performAddition()` and cannot be accessed outside of it.

		3. **Local Scope of `main()`**: This scope can access the function `performAddition()` because `performAddition()` is defined at the same level as `main()`. However, `main()` cannot access the local variables `x` and `y` defined within `performAddition()`.
	*/

	// Calls the performAddition function, which prints the sum of x and y.
	performAddition()

	// The following line will cause a compilation error because `x` is not defined in the scope of `main()`.
	fmt.Println("What if", x, "was different?") // ./scope-err.go:12:26: undefined: x

	/*
		The error occurs because `x` is defined within the local scope of `performAddition()`, not in the scope of `main()`.
		Variables defined inside a function (like `x` and `y` in `performAddition()`) are not accessible outside that function.
		In `main()`, trying to use `x` results in an "undefined" error because `x` does not exist in the `main()` scope.
	*/
}
