// error handling
package main

import (
	"errors"
	"fmt"
)

// divide performs division and handles errors
func divide(a, b float64) (value float64, err error) {
	if b == 0 {
		// Return 0 and an error if the divisor is zero
		// errors.New("message") creates a new error with the specified message.
		// In this case, it indicates that division by zero is not possible.
		return 0, errors.New("division by 0 is not possible")
	} else {
		// Perform division and return the result with a nil error
		// Returning nil for the error indicates that the operation was successful.
		value = a / b
		return value, nil
	}
}

func main() {
	var dividend float64
	var divisor float64

	// Read user input
	fmt.Print("Enter dividend value: ")
	fmt.Scan(&dividend)
	fmt.Print("Enter divisor value: ")
	fmt.Scan(&divisor)

	// Call the divide function and handle the result
	val, err := divide(dividend, divisor)
	if err != nil {
		/*
			In practice, both `err` and `err.Error()` will often produce the same output when printed,
			but `err` provides the error value itself, which is more versatile.
			`err.Error()`: This method returns a string representation of the error.
		*/
		fmt.Println("Error:", err)                 // This will print the error message
		fmt.Println("Error message:", err.Error()) // This also prints the same error message
	} else {
		// Print result if there's no error
		fmt.Println("Value:", val)
	}
}
