package main

import "fmt"

func main() {
	// Declare a variable and assign a value.
	x := "My very first address"

	// Print the value of x.
	fmt.Println("Value of x:", x)

	// Print the address of x using the & operator.
	fmt.Println("Address of x:", &x)

	/*
		   Imagine the computer's memory as a series of numbered slots (addresses).
			When you declare a variable, the computer assigns it to one of these slots.
			The address you see is the identifier for that slot where your variable's data is stored.

		   Key Points
		   Addresses: Every variable in memory is located at a specific address, which is a numerical value representing a location in memory.
		   Hexadecimal Format: Addresses are often displayed in hexadecimal (base-16) format, which is compact and commonly
	*/
}
