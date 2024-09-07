package main

import "fmt"

func main() {
	// fmt.Println() automatically adds a new line after printing.
	fmt.Println("help") // Prints: help (and moves to a new line)

	// fmt.Print() does not add a new line, so the next print starts on the same line.
	fmt.Print("doesnt print in new line") // Prints: doesnt print in new line

	// This will also print on the same line because fmt.Print() is used.
	fmt.Print("this will print after line words") // Prints: this will print after line words

	// fmt.Println() separates each argument with a space and adds a new line after printing.
	fmt.Println("this", "will", "print", "with", "space") // Prints: this will print with space

	// fmt.Print() does not automatically add spaces between arguments or a new line.
	fmt.Print("this", "wont", "print", "with", "space") // Prints: thiswontprintwithspace
}
