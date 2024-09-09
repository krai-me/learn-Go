//Custom error handling allows you to create your own error types.
//This is useful when you want to provide more information about an error or handle different errors in different ways.

package main

import (
	"fmt"
)

// Step 1: Define a custom error type
type AgeError struct {
	Age int
}

// Step 2: Implement the Error() method for the custom error type
func (e *AgeError) Error() string {
	return fmt.Sprintf("Invalid age: %d. Age must be at least 18.", e.Age)
}

// Function that returns an error if age is less than 18
func checkAge(age int) error {
	if age < 18 {
		// Step 3: Return a pointer to a new AgeError
		return &AgeError{Age: age}
	}
	// No error if age is valid
	return nil
}

/*
Key Concepts Simplified:
error is just an interface(hover to check):
It's like a contract that says, "Any type that has an Error() method that returns a string can be treated as an error."

Your custom error type (AgeError) implements that contract:
By defining the Error() method for AgeError, you're saying, "Hey Go, my AgeError can be treated like an error!"

Go automatically calls Error():
When you pass your AgeError (or any error) to a function like fmt.Println(), Go knows to call the Error() method to get the error message.

Why This Happens:
When you return a pointer like &AgeError{Age: 16}, you're giving Go a value that implements the error interface.
When you later try to print or use this error, Go automatically calls the Error() method to get the message.
*/
func main() {
	age := 16 // Let's test with an invalid age

	// Check if the age is valid, and handle the error
	err := checkAge(age)
	if err != nil {
		// Step 4: Handle the custom error
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Age is valid.")
	}
}
