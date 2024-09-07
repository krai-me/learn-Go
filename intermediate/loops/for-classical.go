// Definite loops can be used to repeat code a fixed number of times.
package main

import "fmt"

func main() {
	//note dont use () here means for(){} is not valid
	for number := 0; number < 5; number++ {
		fmt.Print(number, " ") // Output: 0 1 2 3 4

	}
	/*
		The initial statement, number := 0, creates a new variable to be used within the for loop code block.

		The conditional expression, number < 5, will stop the loop when number reaches the target value of 5.

		The post statement, number++, increases the value of the number variable by 1 each time the code block completes.
	*/
}
