package main

import "fmt"

func main() {
	//A variable is a named value (like a constant) with the added feature that it can change during the running of a program
	var lengthOfSong uint16
	var isMusicOver bool
	var songRating float32
	var emptyString string
	/*
		An unsigned 16-bit integer called lengthOfSong.
		A boolean value called isMusicOver.
		A 32-bit float called songRating.
	*/
	fmt.Println(lengthOfSong) // returns 0
	fmt.Println(isMusicOver)  // returns false if nothing is defined
	fmt.Println(songRating)   // returns 0
	fmt.Println(emptyString)  // Doesn't print anything
	//to redefine
	lengthOfSong = 9
	fmt.Println(lengthOfSong) // returns 9

	//also valid
	var kilometersToMars int32 = 62100000 // good practice to write datatype included
	fmt.Println(kilometersToMars)
}
