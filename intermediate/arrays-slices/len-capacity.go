package main

import "fmt"

func main() {
	// Define an array with a fixed size of 6 elements.
	numbers := [...]int{1, 2, 3, 4, 4, 23}
	// Define a slice with initial values. The slice is created with the default capacity of 4.
	numbers_1 := []int{1, 23, 2342, 24}

	// For arrays, length is fixed and determined at compile time.
	// For slices, length is the number of elements currently in the slice.
	fmt.Println("Length of numbers array:", len(numbers))     // Output: 6
	fmt.Println("Length of numbers_1 slice:", len(numbers_1)) // Output: 4

	// Capacity is the number of elements the slice can hold before needing to resize itself.
	fmt.Println("numbers_1 slice:", numbers_1, "Length:", len(numbers_1), "Capacity:", cap(numbers_1))

	// Append a new element to the slice. This may cause the slice to be resized if necessary.
	numbers_1 = append(numbers_1, 12)

	// When appending, if the slice's capacity is exceeded, it is increased, usually doubling in size.
	fmt.Println("Updated numbers_1 slice:", numbers_1, "Length:", len(numbers_1), "Capacity:", cap(numbers_1))
	/*
	   When the slice runs out of capacity and needs more space,
	   Go typically doubles its capacity. This means if the slice
	   can currently hold 4 elements, it will grow to hold 8, then 16,
	   and so on. Doubling the capacity reduces the frequency of
	   resizing operations, making the program more efficient by minimizing
	   the amount of time spent reallocating memory.
	*/

}
