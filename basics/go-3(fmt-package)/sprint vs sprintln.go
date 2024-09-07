package main

import "fmt"

func main() {
	grade := "100"
	compliment := "Great job!"

	// fmt.Sprint() concatenates arguments into a single string without printing.
	// It doesn't add spaces or new lines between arguments.
	teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

	// Now we print the concatenated string using fmt.Println().
	fmt.Println(teacherSays) // Prints: You scored a 100 on the test! Great job!

	/*
		fmt.Sprintln() works like fmt.Sprint() but it automatically includes
		spaces between the arguments and adds a new line at the end.
	*/
	quote := fmt.Sprintln("Look ma,", "no spaces!")
	fmt.Println(quote) // Prints: Look ma, no spaces!

	/*
		fmt.Sprintf() is similar to fmt.Printf() but instead of printing,
		it returns the formatted string. This allows us to store the result in a variable.
		It can use format verbs like %v to format variables.
	*/
	correctAns := "A"
	answer := fmt.Sprintf("And the correct answer is… %v!", correctAns)

	// We can print the stored string using fmt.Print().
	fmt.Print(answer) // Prints: And the correct answer is… A!
}
