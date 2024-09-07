package main

import "fmt"

func main() {
	var fruit string
	fruit = "apple" // note keep string in "" not single quotes in goland it used for diff
	//string concatenation
	//Note that + does not add in additional spaces or punctuation.
	var description string
	var nameOfSong string = "Stop Stop"
	var nameOfArtist string = "The Julie Ruin"
	description = nameOfSong + " is by: " + nameOfArtist + "."
	fmt.Println(description) // Prints "Stop Stop is by: The Julie Ruin.

	command := "Hold my "
	beverage := "soda"

	command += beverage
	fmt.Println(command) // Prints: Hold my soda
}
