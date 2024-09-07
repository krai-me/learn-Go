package main

import (
	"fmt"
	"math/rand"
)

/*Default Seed: Without seeding, you might see different numbers by chance, but it’s not reliable for applications requiring true randomness. */
func main() {
	/*
		In our main function, we’re printing out a random number using rand and the Intn() method. With the argument of 100,
		the maximum value that the method will return is 99.
	*/
	fmt.Println(rand.Intn(100))
	fmt.Println(1000 + rand.Intn(9000)) // to gen random 4 digit number

}
