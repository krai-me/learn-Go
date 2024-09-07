package main

import "fmt"

/*
	func addHundred(num int) {
		num += 100
	}

	func main() {
		x := 1
		addHundred(x)
		fmt.Println(x) // Prints 1

		//	Remember, Go is a pass-by-value language. When we call addHundred(x) we’re providing addHundred() with a value of 1.
		//	We’re not actually providing the address of x for addHundred() to go in and change the value stored there.

}
*/
func addHundred(numPtr *int) {
	*numPtr += 100
}

func main() {
	x := 1
	addHundred(&x)
	fmt.Println(x) // Prints 101
	/*
		  	addHundred(&x) passes the address of x to the function. The function receives a pointer (numPtr) that refers to the original x. By dereferencing numPtr with *numPtr,
			the function directly modifies the value stored at that address, which affects x in main.
	*/
}
