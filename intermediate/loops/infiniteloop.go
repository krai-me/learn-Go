package main

import (
	"fmt"
)

var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)
	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {

	count()
	/*syntax for infinite loop its keeps repeating as there logic for stopping*/
	for {
		count()
	}
}
