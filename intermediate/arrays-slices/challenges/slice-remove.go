// Slices in Go do not have a built-in method to remove an element.
// We need to implement this functionality using custom logic.

package main

import (
	"fmt"
)

func main() {
	num := []int{23, 1, 10, 12}
	fmt.Println("Initial slice: ", num)
	var uinput int
	fmt.Println("Select a number to remove:")
	fmt.Scan(&uinput)
	if !check(num, uinput) {
		fmt.Println("Number not found in the slice.")
	} else {
		defer fmt.Println("Done...")
		fmt.Println("Starting...")
		// reassigning
		num = remove(num, uinput)
		fmt.Println("Updated slice:", num)
	}
}

// You can try this function by using an int variable instead of returning a bool.
// The underscore (_) is used as a placeholder to ignore values that you don't need.
// This allows you to avoid syntax errors while omitting unnecessary variables.
func check(array []int, value int) bool {

	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func remove(array []int, value int) []int {
	var result []int
	for _, v := range array {
		if v != value {
			// fmt.Println(v)
			result = append(result, v)
		}
	}
	return result
}
