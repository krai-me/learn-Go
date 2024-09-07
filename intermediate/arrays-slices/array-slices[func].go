package main

import "fmt"

func main() {
	// Define an array and a slice
	myArray := [4]int{10, 20, 30, 40}
	mySlice := []int{10, 20, 30, 40, 50}

	// Call functions to print the first and last elements
	printFirstLastArray(myArray)
	printFirstLastSlice(mySlice)

	// Modify the array and slice
	changeFirst(myArray, 99)
	changeFirstSlice(mySlice, 99)

	// Print the modified array and slice
	fmt.Println("Array after changeFirst:", myArray)
	fmt.Println("Slice after changeFirstSlice:", mySlice)
}

// Function to print the first and last elements of an array
func printFirstLastArray(array [4]int) {
	fmt.Println("First (Array):", array[0])
	fmt.Println("Last (Array):", array[3])
}

// Function to print the first and last elements of a slice
func printFirstLastSlice(slice []int) {
	length := len(slice)
	if length > 0 {
		fmt.Println("First (Slice):", slice[0])
		fmt.Println("Last (Slice):", slice[length-1])
	}
}

// Function to change the first element of an array
// Note: This change will only be local to the function
func changeFirst(array [4]int, value int) {
	array[0] = value
}

// Function to change the first element of a slice
// Note: This change will be reflected outside the function
func changeFirstSlice(slice []int, value int) {
	if len(slice) > 0 {
		slice[0] = value
	}
}
