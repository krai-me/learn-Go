package main

import "fmt"

func main() {

	/*
		Using fmt.Println() and fmt.Print(), we can concatenate strings,
		i.e., combine different strings into a single string:
	*/
	guess := "C"
	fmt.Println("Is", guess, "your final answer?") // Prints: Is C your final answer?

	// %v is a placeholder for values (variables).
	// fmt.Printf doesn't automatically add a new line, so it will continue on the same line unless we specify \n.
	selection1 := "soup"
	selection2 := "salad"
	// \n creates a new line after printing.
	fmt.Printf("Do I want %v or %v?\n", selection1, selection2) // Prints: Do I want soup or salad?

	// %T is used to check and print the data type of a value.
	specialNum := 42
	fmt.Printf("This value's type is %T.\n", specialNum) // Prints: This value's type is int.

	// %d is used to interpolate an integer into a string.
	votingAge := 18
	fmt.Printf("You must be %d years old to vote.\n", votingAge) // Prints: You must be 18 years old to vote.

	// %f is used for floating-point numbers.
	// With %.2f, we limit the precision to 2 decimal places.
	gpa := 3.8
	fmt.Printf("You're averaging: %.2f.\n", gpa) // Prints: You're averaging 3.80.

}
