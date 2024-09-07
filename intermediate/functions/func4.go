package main

import "fmt"

// Define a function that returns multiple values: a string and a float32.
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
	// Calculate the average grade.
	averageGrade := (midtermGrade + finalGrade) / 2

	// Determine the letter grade based on the average grade.
	var gradeLetter string
	if averageGrade > 90 {
		gradeLetter = "A"
	} else if averageGrade > 80 {
		gradeLetter = "B"
	} else if averageGrade > 70 {
		gradeLetter = "C"
	} else if averageGrade > 60 {
		gradeLetter = "D"
	} else {
		gradeLetter = "F"
	}

	// Return both the letter grade and the average grade.
	return gradeLetter, averageGrade
}

func main() {
	// Declare variables to store midterm and final grades.
	var myMidterm, myFinal float32
	myMidterm = 89.4
	myFinal = 74.9

	// Declare variables to store the results from GPA().
	var myAverage float32
	var myGrade string

	// Call GPA() and assign the returned values to variables.
	myGrade, myAverage = GPA(myMidterm, myFinal)

	// Print the returned values: average grade and letter grade.
	fmt.Println(myAverage, myGrade) // Prints: 82.12 B

	/*
	   **Multiple Return Values**:
	   - In Go, a function can return multiple values. The return types are specified in parentheses after the function's parameter list.
	   - The function `GPA` is defined to return two values: a `string` and a `float32`. This is indicated by the return type `(string, float32)`.
	   - When returning multiple values, they are separated by commas in both the return statement and when assigning the return values.

	   **Example of Returning Multiple Values**:
	   - The function `GPA` calculates an average grade and assigns a letter grade based on the average.
	   - The `return` statement in `GPA` returns both the `gradeLetter` and `averageGrade`.
	   - In `main()`, the variables `myGrade` and `myAverage` receive these returned values. The `fmt.Println` statement then prints both values.

	   **Handling Multiple Return Values**:
	   - Ensure that you capture all return values when calling the function. If you don't, you will get an error. For example:
	   - GPA(89.4, 74.9) // Error: cannot use result of GPA as single value
	*/
}
