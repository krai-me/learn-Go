package main

import "fmt"

func main() {
	// Prints a message to the console asking the user how they are doing
	fmt.Println("How are you doing?")

	// Declares a variable 'response' to store the user input
	var response string

	// fmt.Scan(&response) waits for user input and assigns the value to 'response'
	// The '&' before 'response' is used to pass the address of the variable to the Scan function.
	// It means the user input will be stored in the 'response' variable.
	fmt.Scan(&response)

	// Prints out the user's response using fmt.Printf.
	// %v is a placeholder for the value stored in the 'response' variable.
	fmt.Printf("I'm %v.\n", response)
}
