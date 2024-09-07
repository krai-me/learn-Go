package main

import (
	"fmt"
)

func main() {
	// Short variable declaration in an if statement
	x := 8
	y := 9
	if product := x * y; product > 60 {
		fmt.Println(product, "is greater than 60")
	}
	// The following line will produce an error because 'product' is out-of-scope here
	// fmt.Println(product) // Uncommenting this line will cause a compile-time error

	// Short variable declaration in a switch statement
	switch season := "summer"; season {
	case "summer":
		fmt.Println("Go out and enjoy the sun!")
	case "winter":
		fmt.Println("Stay warm and cozy!")
	default:
		fmt.Println("Have a nice day!")
	}
	// 'season' is only accessible within the switch statement block
}
